#!/usr/bin/env bash
# After pushing the provider image to ghcr.io, run this to get the digest and the
# cosign sign command. Signing by digest is preferred over tag (avoids tag reassignment).
#
# Usage: ./scripts/get-digest-and-cosign.sh [version]
#   version defaults to v0.0.0
#
# Requires: crane (ghcr.io/google/go-containerregistry/cmd/crane) or docker.
# Example: crane digest ghcr.io/prolixalias/provider-cloudflare:v0.0.0

set -e

VERSION="${1:-v0.0.0}"
IMAGE="ghcr.io/prolixalias/provider-cloudflare"
REF="${IMAGE}:${VERSION}"

if command -v crane &>/dev/null; then
  DIGEST=$(crane digest "$REF" 2>/dev/null || true)
elif docker buildx imagetools inspect "$REF" &>/dev/null; then
  DIGEST=$(docker buildx imagetools inspect "$REF" --format '{{.Digest}}' 2>/dev/null || true)
else
  echo "Need 'crane' or 'docker' to resolve digest. Install crane: go install github.com/google/go-containerregistry/cmd/crane@latest"
  exit 1
fi

if [[ -z "$DIGEST" ]]; then
  echo "Could not get digest for $REF. Is the image pushed to ghcr.io?"
  exit 1
fi

echo "Image:  $REF"
echo "Digest: $DIGEST"
echo ""
echo "Sign with cosign (use the digest, not the tag):"
echo "  cosign sign --key ~/.cosign/cosign.key ${IMAGE}@${DIGEST}"
echo ""
