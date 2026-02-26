# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

PROJECT_REPO ?= github.com/prolixalias/provider-cloudflare
IMAGE_SOURCE ?= https://$(PROJECT_REPO)
# OCI labels for registry UI (optional overrides; Dockerfile has defaults)
IMAGE_TITLE ?= Crossplane provider for Cloudflare
IMAGE_DESCRIPTION ?= Crossplane provider for Cloudflare. Manages DNS, tunnels, Zero Trust, and other Cloudflare resources.
IMAGE_VENDOR ?= prolixalias

# Terraform external-binary build (for TrustTunnelCloudflared etc.). Match root Makefile defaults.
TERRAFORM_PROVIDER_SOURCE ?= cloudflare/cloudflare
TERRAFORM_PROVIDER_VERSION ?= 5.16.0
TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-cloudflare
TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= https://github.com/cloudflare/terraform-provider-cloudflare/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-cloudflare_v$(TERRAFORM_PROVIDER_VERSION)

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) docker build $(IMAGE)
	@if [ "$(BUILD_TERRAFORM_EXTERNAL)" = "1" ]; then \
		$(MAKE) BUILD_ARGS="--load" img.build.terraform-external; \
	else \
		$(MAKE) BUILD_ARGS="--load" img.build.shared; \
	fi
	@$(OK) docker build $(IMAGE)

img.publish:
	@$(INFO) Skipping image publish for $(IMAGE)
	@echo Publish is deferred to xpkg machinery
	@$(OK) Image publish skipped for $(IMAGE)

# Native image: no Terraform binary (provider linked in-process). Use
# img.build.terraform-external for Tunnel/resources that need the async external binary.
img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@docker build $(BUILD_ARGS) \
		--platform $(OS)/$(ARCH) \
		--build-arg TARGETOS=$(OS) \
		--build-arg TARGETARCH=$(ARCH) \
		--build-arg IMAGE_SOURCE=$(IMAGE_SOURCE) \
		--build-arg IMAGE_TITLE="$(IMAGE_TITLE)" \
		--build-arg IMAGE_DESCRIPTION="$(IMAGE_DESCRIPTION)" \
		--build-arg IMAGE_VENDOR="$(IMAGE_VENDOR)" \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

# Terraform external-binary image: required for TrustTunnelCloudflared (async connector).
# Build with: from repo root, BUILD_TERRAFORM_EXTERNAL=1 make build.multiarch.linux
# Or: make -C cluster/images/provider-cloudflare BUILD_TERRAFORM_EXTERNAL=1 IMAGE=ghcr.io/prolixalias/provider-cloudflare:$(VERSION) img.build.terraform-external (per platform).
img.build.terraform-external:
	@cp Dockerfile.terraform-external $(IMAGE_TEMP_DIR)/Dockerfile || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@docker build $(BUILD_ARGS) \
		--platform $(OS)/$(ARCH) \
		--build-arg TARGETOS=$(OS) \
		--build-arg TARGETARCH=$(ARCH) \
		--build-arg IMAGE_SOURCE=$(IMAGE_SOURCE) \
		--build-arg IMAGE_TITLE="$(IMAGE_TITLE)" \
		--build-arg IMAGE_DESCRIPTION="$(IMAGE_DESCRIPTION)" \
		--build-arg IMAGE_VENDOR="$(IMAGE_VENDOR)" \
		--build-arg TERRAFORM_PROVIDER_SOURCE="$(TERRAFORM_PROVIDER_SOURCE)" \
		--build-arg TERRAFORM_PROVIDER_VERSION="$(TERRAFORM_PROVIDER_VERSION)" \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME="$(TERRAFORM_PROVIDER_DOWNLOAD_NAME)" \
		--build-arg TERRAFORM_NATIVE_PROVIDER_BINARY="$(TERRAFORM_NATIVE_PROVIDER_BINARY)" \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX="$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX)" \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)
