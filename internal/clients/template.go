package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlLog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/prolixalias/provider-cloudflare/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/prolixalias/provider-cloudflare/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal template credentials as JSON"
)

// TerraformSetupBuilder builds a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder() terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}
		logger := ctrlLog.FromContext(ctx).WithValues(
			"managedType", fmt.Sprintf("%T", mg),
			"managedName", mg.GetName(),
			"managedNamespace", mg.GetNamespace(),
			"providerConfigRef", providerConfigRefSummary(mg),
		)

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			logger.Error(err, "Terraform setup failed while resolving ProviderConfig")
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			logger.Error(err, "Terraform setup failed while extracting credentials", "credentialSource", pcSpec.Credentials.Source)
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			logger.Error(err, "Terraform setup failed while unmarshalling credentials JSON", "credentialBytes", len(data))
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		credKeys := sortedKeys(creds)

		// Set Cloudflare credentials in provider configuration
		ps.Configuration = map[string]any{}
		ps.FrameworkProvider = getFrameworkProvider()
		if ps.FrameworkProvider == nil {
			err := errors.New("terraform framework provider factory returned nil")
			logger.Error(err, "Terraform setup cannot configure framework provider", "hint", "ensure non-ci build includes terraform provider package")
			return ps, err
		}
		hasToken := false
		hasKey := false
		hasEmail := false
		if v, ok := creds["api_token"]; ok && v != "" {
			ps.Configuration["api_token"] = v
			hasToken = true
		}
		if v, ok := creds["api_key"]; ok && v != "" {
			ps.Configuration["api_key"] = v
			hasKey = true
		}
		if v, ok := creds["email"]; ok && v != "" {
			ps.Configuration["email"] = v
			hasEmail = true
		}

		// Cloudflare auth requires api_token OR api_key+email.
		if !(hasToken || (hasKey && hasEmail)) {
			err := errors.New("credentials must include api_token or both api_key and email")
			logger.Error(err, "Terraform setup extracted credentials with unsupported shape", "credentialKeys", credKeys)
			return ps, err
		}

		// Emit extra runtime context for tunnel resources, where failures are currently opaque.
		if isTunnelManaged(mg) {
			tfPath := os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH")
			st, statErr := os.Stat(tfPath)
			if statErr != nil {
				logger.Error(statErr, "Terraform setup (tunnel): provider binary stat failed", "terraformNativeProviderPath", tfPath, "credentialKeys", credKeys)
			} else {
				logger.Info("Terraform setup (tunnel): provider config resolved", "terraformNativeProviderPath", tfPath, "providerBinarySizeBytes", st.Size(), "credentialKeys", credKeys)
			}
		}

		return ps, nil
	}
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	pc, ok := pcObj.(*namespacedv1beta1.ProviderConfig)
	if !ok {
		return nil, errors.New("unknown provider config type")
	}
	pcSpec := pc.Spec
	if pcSpec.Credentials.SecretRef != nil {
		pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
	}
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}

func providerConfigRefSummary(mg resource.Managed) string {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		ref := managed.GetProviderConfigReference()
		if ref == nil {
			return "<nil>"
		}
		return ref.Name
	case resource.ModernManaged:
		ref := managed.GetProviderConfigReference()
		if ref == nil {
			return "<nil>"
		}
		if ref.Kind != "" {
			return fmt.Sprintf("%s/%s", ref.Kind, ref.Name)
		}
		return ref.Name
	default:
		return "<unknown-managed-type>"
	}
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func isTunnelManaged(mg resource.Managed) bool {
	// Keep this generic so generated type package moves don't break diagnostics.
	return strings.Contains(fmt.Sprintf("%T", mg), "TrustTunnelCloudflared")
}
