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
