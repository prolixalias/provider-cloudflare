# Native provider image: no Terraform binary. The Crossplane provider binary
# has the Cloudflare Terraform Plugin Framework provider linked in-process
# (config/provider_tfprovider.go getTerraformProvider()). Do not set
# TERRAFORM_NATIVE_PROVIDER_PATH so upjet uses the in-process provider.
FROM alpine:3.23.2
RUN apk --no-cache add ca-certificates bash

# OCI labels for registry UI (description, vendor, title) and ghcr.io source link
ARG IMAGE_SOURCE
ARG IMAGE_TITLE="Crossplane provider for Cloudflare"
ARG IMAGE_DESCRIPTION="Crossplane provider for Cloudflare. Manages DNS, tunnels, Zero Trust, and other Cloudflare resources."
ARG IMAGE_VENDOR="prolixalias"
LABEL org.opencontainers.image.source="${IMAGE_SOURCE}"
LABEL org.opencontainers.image.title="${IMAGE_TITLE}"
LABEL org.opencontainers.image.description="${IMAGE_DESCRIPTION}"
LABEL org.opencontainers.image.vendor="${IMAGE_VENDOR}"

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532
USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
