# Comparison with Official Upjet Providers (e.g. provider-upjet-azure)

This document summarizes how provider-cloudflare compares to [provider-upjet-azure](https://github.com/crossplane-contrib/provider-upjet-azure) and what was added or adjusted to align with common practices.

## What We Aligned or Added

| Area | Azure / official style | What we did in provider-cloudflare |
|------|------------------------|-------------------------------------|
| **README** | Provider-specific title, badges (CI, Go Report Card, release, Slack), Getting Started, Marketplace/quickstart links, Contributing, Getting help, License + NOTICE | Rewrote README for Cloudflare: badges (CI, Go Report Card), install instructions, development commands, links to AUTHENTICATION.md and issues. |
| **CI** | GitHub Actions: lint, check-diff, unit tests, optional vendor check, optional local-deploy / uptest | Added `.github/workflows/ci.yml`: lint, check-diff, unit tests (no vendor; we use go mod). |
| **NOTICE** | Third-party attributions (e.g. Terraform provider under MPL 2.0) | Added NOTICE with Crossplane and Terraform Cloudflare provider attribution. |
| **AUTHENTICATION** | Dedicated AUTHENTICATION.md with mechanisms and examples | Added AUTHENTICATION.md for Cloudflare: API token (recommended) and legacy API key, with Secret + ProviderConfig examples. |
| **PULL_REQUEST_TEMPLATE** | Short “description of changes” and contribution checklist | We already had a PR template; consider aligning wording with [contribution process](https://git.io/fj2m9) if desired. |
| **Package / image** | OCI label for source repo on ghcr.io | We added `org.opencontainers.image.source` in the provider image Dockerfile and build args. |

## Optional / Future Improvements

- **Reusable CI from provider-workflows**  
  Azure uses `crossplane-contrib/provider-workflows` (e.g. `ci.yml`, `publish-provider-packages.yaml`, `tag.yaml`). We use a self-contained `ci.yml`; you can later switch to the shared workflows for consistency and less maintenance.

- **Issue templates**  
  Azure has `.github/ISSUE_TEMPLATE/` for bug/feature forms. Adding similar templates can make reporting easier.

- **Stale issue workflow**  
  Azure runs `.github/workflows/stale.yml` to mark/close stale issues/PRs. Optional quality-of-life improvement.

- **Tag / release workflow**  
  Azure has `.github/workflows/tag.yaml` (e.g. with `negz/create-tag`) for creating release tags. Adding this (or using provider-workflows’ tag workflow) would standardize releases.

- **Publish workflow**  
  Azure uses a publish workflow that calls `provider-workflows`’ `publish-provider-family.yml` (or similar) to build and push packages. We currently publish via `make build.all` and `make publish`; a GitHub Action could mirror that for releases.

- **Marketplace / docs**  
  Azure points to Upbound Marketplace and family docs. If you publish to the marketplace, add a “Documentation” / “Quickstart” link in the README.

- **OWNERS.md / CODEOWNERS**  
  We have template OWNERS.md and CODEOWNERS. Filling in real maintainers and reviewers improves clarity and PR routing.

- **Extensions (readme / release notes)**  
  Our `extensions/readme/readme.md` and `extensions/release-notes/release_notes.md` are placeholders; filling them helps marketplace and release pages if you use that pipeline.

## Summary

We’ve closed the main gaps that make a provider look “less than popular”: a Cloudflare-specific README with badges and clear install/contributing/help sections, CI (lint, check-diff, unit tests), NOTICE, AUTHENTICATION.md, and OCI image source labeling. The optional items above can be adopted over time to match official providers more closely.
