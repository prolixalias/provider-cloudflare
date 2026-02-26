# Upjet-based Crossplane provider for Cloudflare

[![Go Report Card](https://goreportcard.com/badge/github.com/prolixalias/provider-cloudflare)](https://goreportcard.com/report/github.com/prolixalias/provider-cloudflare)

Provider Cloudflare is a [Crossplane](https://www.crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) and exposes XRM-conformant managed resources for [Cloudflare](https://www.cloudflare.com/).

## Getting Started

- **Install the provider** (see [Installation](#installation)).
- **Configure authentication** – see [AUTHENTICATION.md](AUTHENTICATION.md) for API token and scopes.
- **Examples** – see the [examples](examples/) directory and [ProviderConfig](examples/cluster/providerconfig/) for usage.

## Installation

Install the provider into your cluster (replace the image tag with a [released version](https://github.com/prolixalias/provider-cloudflare/releases) if needed):

```bash
kubectl crossplane install provider ghcr.io/prolixalias/provider-cloudflare:v0.0.0
```

Or use the [installation manifest](examples/install.yaml) and apply with `kubectl apply -f examples/install.yaml`.

## Developing

- **Code generation** (after changing config):
  ```bash
  go run cmd/generator/main.go "$PWD"
  ```
- **Run locally (out-of-cluster)**:
  ```bash
  make run
  ```
- **Build and test**:
  ```bash
  make build
  make test
  ```
- **Build for both linux/amd64 and linux/arm64** (e.g. for clusters that need amd64; default `make build` on Mac only produces the host’s arch):
  ```bash
  VERSION=v0.0.0 make build.multiarch.linux
  ```
  The Makefile runs `binfmt.install` first so QEMU is available for cross-arch Docker builds (e.g. amd64 on arm64). If that step fails (e.g. no `--privileged` docker), run manually: `docker run --privileged --rm tonistiigi/binfmt --install all`
- **Build and push multiarch package to ghcr.io** (requires `docker login ghcr.io`):
  ```bash
  VERSION=v0.0.0 make push.multiarch
  ```
- **Tunnel (TrustTunnelCloudflared)** – This resource uses the Terraform Plugin Framework async connector and requires the provider image to include the Terraform provider binary. Build and push the **terraform-external** image:
  ```bash
  VERSION=v0.0.2 make build.terraform-external.multiarch.linux
  VERSION=v0.0.2 make push.terraform-external.multiarch
  ```
  Use that tag (e.g. `v0.0.2`) in your cluster; the default `make push.multiarch` image is native (no Terraform binary) and will report "cannot retrieve framework provider" for Tunnel.
- **Reviewable (generate, lint, test)**:
  ```bash
  make reviewable
  ```

## Contributing

- Open an [issue](https://github.com/prolixalias/provider-cloudflare/issues) for bugs or feature requests.
- See [Upjet contribution guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md) for adding resources and general contribution flow.

## Getting help

- [Open an issue](https://github.com/prolixalias/provider-cloudflare/issues/new/choose) for bugs or questions about this provider.
- [Crossplane Slack](https://slack.crossplane.io) for general Crossplane and Upjet help.

## License

This provider is released under the [Apache 2.0 license](LICENSE) with [NOTICE](NOTICE).
