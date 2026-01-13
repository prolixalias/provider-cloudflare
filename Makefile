# ====================================================================================
# Setup Project

VERSION ?= v0.0.0

PROJECT_NAME ?= provider-cloudflare
PROJECT_REPO ?= github.com/prolixalias/$(PROJECT_NAME)

export TERRAFORM_PROVIDER_SOURCE ?= cloudflare/cloudflare
export TERRAFORM_PROVIDER_REPO ?= https://github.com/cloudflare/terraform-provider-cloudflare
export TERRAFORM_PROVIDER_VERSION ?= 5.16.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-cloudflare
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/cloudflare/terraform-provider-cloudflare/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-cloudflare_v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_DOCS_PATH ?= docs/resources



PLATFORMS ?= linux_amd64 linux_arm64

LOCAL_XPKG_BUILD ?= false
ifeq ($(LOCAL_XPKG_BUILD),true)
BUILD_REGISTRY := localhost/provider-cloudflare-build
endif

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= 1.25
GOLANGCILINT_VERSION ?= 2.8.0
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis
-include build/makelib/golang.mk

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.31.0
UPTEST_VERSION = v2.2.0
CRDDIFF_VERSION = v0.12.1
CROSSPLANE_CLI_VERSION = v2.1.3
# for e2e testing
CROSSPLANE_VERSION = 2.1.3
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

REGISTRY_ORGS ?= ghcr.io/prolixalias
IMAGES = $(PROJECT_NAME)
-include build/makelib/imagelight.mk

# ====================================================================================
# Setup XPKG

XPKG_REG_ORGS ?= ghcr.io/prolixalias
# NOTE(hasheddan): skip promoting on xpkg.crossplane.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= ghcr.io/prolixalias
XPKGS = $(PROJECT_NAME)
-include build/makelib/xpkg.mk

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# NOTE(hasheddan): we force image building to happen prior to xpkg build so that
# we ensure image is present in daemon.
xpkg.build.provider-cloudflare: do.build.images

# NOTE(hasheddan): we ensure up is installed prior to running platform-specific
# build steps in parallel to avoid encountering an installation race condition.
build.init: $(UP) $(CROSSPLANE_CLI)

# Post-generation hook to fix CEL validation errors for dynamic fields and PCU getter
generate.done:
	@$(INFO) Fixing CRD validation rules for dynamic fields
	@./hack/fix-crds.sh
	@$(INFO) Fixing namespaced ProviderConfigUsage GetProviderConfigReference for TypedProviderConfigUsage
	@./hack/fix-pcu-getter.sh
	@$(OK) CRD validation rules fixed

# ====================================================================================
# Setup Provider Schema Generation (Modern Upjet v2+ with forked provider module)
# The forked provider at github.com/prolixalias/terraform-provider-cloudflare/v5
# already exports the provider via provider/provider.go wrapper, solving the v5 issue.

pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" ]; then \
   		mkdir -p "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" && \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)"; \
	fi
	@git -C "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: pull-docs

.PHONY: pull-docs

# ====================================================================================
# Targets

# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

go.mod.cachedir:
	@go env GOMODCACHE

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	$(GO_OUT_DIR)/provider --debug

# ====================================================================================
# End to End Testing
CROSSPLANE_NAMESPACE = crossplane-system
-include build/makelib/local.xpkg.mk
-include build/makelib/controlplane.mk

# This target requires the following environment variables to be set:
# - UPTEST_EXAMPLE_LIST, a comma-separated list of examples to test
#   To ensure the proper functioning of the end-to-end test resource pre-deletion hook, it is crucial to arrange your resources appropriately.
#   You can check the basic implementation here: https://github.com/crossplane/uptest/blob/main/internal/templates/03-delete.yaml.tmpl.
# - UPTEST_CLOUD_CREDENTIALS (optional), multiple sets of AWS IAM User credentials specified as key=value pairs.
#   The support keys are currently `DEFAULT` and `PEER`. So, an example for the value of this env. variable is:
#   DEFAULT='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   PEER='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   The associated `ProviderConfig`s will be named as `default` and `peer`.
# - UPTEST_DATASOURCE_PATH (optional), please see https://github.com/crossplane/uptest#injecting-dynamic-values-and-datasource
uptest: $(UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@$(INFO) running automated tests
	@UPTEST_EXAMPLE_LIST="$(UPTEST_EXAMPLE_LIST)" DATASOURCE_PATH="$(UPTEST_DATASOURCE_PATH)" CLOUD_CREDENTIALS="$(UPTEST_CLOUD_CREDENTIALS)" $(UPTEST) e2e --setup-script=cluster/test/setup.sh --default-timeout=30m --data-source=cluster/test/datasource.yaml || $(FAIL)
	@$(OK) running automated tests

uptest-render: $(UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@$(INFO) rendering manifests
	@UPTEST_EXAMPLE_LIST="$(UPTEST_EXAMPLE_LIST)" DATASOURCE_PATH="$(UPTEST_DATASOURCE_PATH)" CLOUD_CREDENTIALS="$(UPTEST_CLOUD_CREDENTIALS)" $(UPTEST) manifests render --setup-script=cluster/test/setup.sh --default-timeout=30m --data-source=cluster/test/datasource.yaml || $(FAIL)
	@$(OK) rendering manifests

crddiff: $(UPTEST)
	@$(INFO) Checking breaking CRD schema changes
	@for crd in $${MODIFIED_CRD_LIST}; do \
		if ! git cat-file -e "$${GITHUB_BASE_REF}:$${crd}" 2>/dev/null; then \
			echo "CRD $${crd} does not exist in the $${GITHUB_BASE_REF} branch. Skipping..." ; \
			continue ; \
		fi ; \
		echo "Checking $${crd} for breaking API changes..." ; \
		changes_detected=$$(go run github.com/crossplane/uptest/cmd/crddiff@$(CRDDIFF_VERSION) revision --enable-upjet-extensions <(git cat-file -p "$${GITHUB_BASE_REF}:$${crd}") "$${crd}" 2>&1) ; \
		if [[ $$? != 0 ]] ; then \
			printf "\033[31m"; echo "Breaking change detected!"; printf "\033[0m" ; \
			echo "$${changes_detected}" ; \
			echo ; \
		fi ; \
	done
	@$(OK) Checking breaking CRD schema changes

schema-version-diff:
	@$(INFO) Checking for native state schema version changes
	@$(FAIL) Not implemented yet

e2e: local-deploy uptest

uptest-render-old: $(UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@$(INFO) rendering manifests
	@UPTEST_EXAMPLE_LIST="$(UPTEST_EXAMPLE_LIST)" DATASOURCE_PATH="$(UPTEST_DATASOURCE_PATH)" CLOUD_CREDENTIALS="$(UPTEST_CLOUD_CREDENTIALS)" $(UPTEST) manifests render --setup-script=cluster/test/setup.sh --default-timeout=30m --data-source=cluster/test/datasource.yaml || $(FAIL)
	@$(OK) rendering manifests

local-deploy: local.xpkg.deploy.provider.$(PROJECT_NAME)
	@$(INFO) running locally built provider
	@$(INFO) waiting for provider to become healthy...
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 10m || \
		(echo "Provider failed to become healthy. Checking status:"; \
		$(KUBECTL) get provider.pkg $(PROJECT_NAME) -o yaml; \
		echo "Provider pod logs:"; \
		$(KUBECTL) -n crossplane-system logs -l pkg.crossplane.io/provider=$(PROJECT_NAME) --tail=50 || true; \
		exit 1)
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Installed --timeout 2m
	@$(OK) running locally built provider

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

endef
# The reason that we have local.xpkg.deploy before local-deploy is that the
# test-ci-local.sh script defined in 'local.xpkg.deploy' needs to call the
# 'local-deploy' target by referring to CROSSPLANE_MAKE_HELP.
export CROSSPLANE_MAKE_HELP

help: crossplane.help

crossplane.help:
	@echo "\n$$(tput bold)Crossplane targets:$$(tput sgr0)"
	@echo "$$CROSSPLANE_MAKE_HELP"

.PHONY: submodules fallthrough help crossplane.help check-terraform-version build.init local-deploy uptest uptest-render
