# Authentication

provider-cloudflare authenticates to the Cloudflare API using credentials configured via a `ProviderConfig` resource. The supported method is **Secret**: a Kubernetes Secret containing your Cloudflare API token (or API key).

## Recommended: API Token

Create a Secret containing your [Cloudflare API Token](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/). The token should have the scopes required for the resources you manage.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: cloudflare-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    CLOUDFLARE_API_TOKEN=<your-api-token>
```

Reference it in a ProviderConfig:

```yaml
apiVersion: cloudflare.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: cloudflare-creds
      namespace: crossplane-system
      key: credentials
```

## Legacy: API Key (not recommended)

You can use the legacy API key with email (see [Cloudflare API keys](https://developers.cloudflare.com/fundamentals/api/get-started/keys/#limitations)):

```yaml
stringData:
  credentials: |
    CLOUDFLARE_EMAIL=your-account-email@example.com
    CLOUDFLARE_API_KEY=<your-api-key>
```

Prefer API tokens over API keys for better security and scoping.

## Scopes

Ensure the API token has the minimum scopes needed for the resources you create (e.g. Zone, DNS, Workers, etc.). See [Cloudflare API token permissions](https://developers.cloudflare.com/fundamentals/api/reference/create-api-token/).
