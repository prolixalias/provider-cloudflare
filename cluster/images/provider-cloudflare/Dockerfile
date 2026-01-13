FROM alpine:3.23.2
RUN apk --no-cache add ca-certificates bash

# OCI label so ghcr.io links this image to the source repository
ARG IMAGE_SOURCE
LABEL org.opencontainers.image.source="${IMAGE_SOURCE}"

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# Native Terraform provider configuration
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

ENV PROVIDER_DIR=/terraform/provider
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PROVIDER_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

RUN mkdir -p ${PROVIDER_DIR}

ADD ${TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/provider.zip

RUN unzip /tmp/provider.zip -d ${PROVIDER_DIR} \
  && chmod +x ${PROVIDER_DIR}/* \
  && rm /tmp/provider.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
