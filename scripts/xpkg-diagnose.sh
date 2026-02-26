#!/usr/bin/env bash
# Deep-dive diagnostic for xpkg build: what the CLI sees (package root) and what
# ends up in the built package (post-build). Use to verify TrustTunnelCloudflared
# and other CRDs are included.
#
# Usage:
#   Pre-build (package root):  XPKG_DIR=/path/to/package scripts/xpkg-diagnose.sh pre
#   Post-build (built xpkg):   scripts/xpkg-diagnose.sh post [path/to/package-version.xpkg]
#   Or set XPKG_FILE for post: XPKG_FILE=path/to/package.xpkg scripts/xpkg-diagnose.sh post
#
# Run from repo root. For pre, XPKG_DIR can be the flattened dir (e.g. _output/xpkg-package-flat).

set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

# CRDs we care about for tunnel (must appear in package root and in built xpkg)
TUNNEL_CRD_UPBOUND="zero.cloudflare.upbound.io_trusttunnelcloudflareds.yaml"
TUNNEL_CRD_M="zero.cloudflare.m.upbound.io_trusttunnelcloudflareds.yaml"

phase="${1:-}"

# --- Pre-build: what the Crossplane CLI will read from the package root ---
pre() {
  local dir="${XPKG_DIR:-}"
  if [[ -z "$dir" ]]; then
    echo "ERROR: XPKG_DIR not set. Usage: XPKG_DIR=/path scripts/xpkg-diagnose.sh pre"
    exit 1
  fi
  if [[ ! -d "$dir" ]]; then
    echo "ERROR: XPKG_DIR is not a directory: $dir"
    exit 1
  fi

  echo "=== xpkg pre-build diagnostic: package root ==="
  echo "XPKG_DIR=$dir"
  echo ""

  # Lexical order (same as filepath.Walk / afero.Walk) so we see the exact stream order
  echo "--- Files in package root (lexical order, as FsBackend will read) ---"
  files=()
  while IFS= read -r -d '' f; do
    files+=("$f")
  done < <(find "$dir" -type f \( -name "*.yaml" -o -name "*.yml" \) -print0 | sort -z)
  count="${#files[@]}"
  echo "Total YAML files: $count"
  echo ""

  for f in "${files[@]}"; do
    echo "  $f"
  done
  echo ""

  echo "--- Critical CRDs (TrustTunnelCloudflared) ---"
  local found_upbound=0 found_m=0
  for f in "${files[@]}"; do
    case "$(basename "$f")" in
      $TUNNEL_CRD_UPBOUND) echo "  PRESENT: $TUNNEL_CRD_UPBOUND"; found_upbound=1 ;;
      $TUNNEL_CRD_M)       echo "  PRESENT: $TUNNEL_CRD_M";       found_m=1 ;;
    esac
  done
  if [[ $found_upbound -eq 0 ]]; then
    echo "  MISSING: $TUNNEL_CRD_UPBOUND (cluster-scoped Tunnel CRD)"
  fi
  if [[ $found_m -eq 0 ]]; then
    echo "  MISSING: $TUNNEL_CRD_M (namespaced Tunnel CRD)"
  fi
  if [[ $found_upbound -eq 0 && $found_m -eq 0 ]]; then
    echo ""
    echo "  => create-tunnel Job will wait in init for this CRD. Run 'make generate' and ensure package root includes these files."
    exit 1
  fi
  echo ""

  echo "--- crossplane.yaml (meta) ---"
  if [[ -f "$dir/crossplane.yaml" ]]; then
    echo "  PRESENT"
  else
    echo "  MISSING"
    exit 1
  fi
  echo ""
  echo "=== pre diagnostic done ==="
}

# --- Post-build: does the built xpkg contain the Tunnel CRD? ---
post() {
  local xpkg="${XPKG_FILE:-${2:-}}"
  if [[ -z "$xpkg" ]]; then
    # Default: latest built xpkg for this provider
    local out_dir="${ROOT}/_output/xpkg"
    local latest
    latest=$(find "$out_dir" -name "provider-cloudflare-*.xpkg" -type f 2>/dev/null | head -1)
    if [[ -z "$latest" ]]; then
      echo "ERROR: No built .xpkg found. Build first (e.g. make xpkg.build.provider-cloudflare), or pass path: $0 post /path/to/package.xpkg"
      exit 1
    fi
    xpkg="$latest"
  fi
  if [[ ! -f "$xpkg" ]]; then
    echo "ERROR: Not a file: $xpkg"
    exit 1
  fi

  echo "=== xpkg post-build diagnostic: built package ==="
  echo "XPKG_FILE=$xpkg"
  echo ""

  # Try 1: use "up xpkg xp-extract" if available (output is gzipped package stream)
  if command -v up &>/dev/null; then
    echo "--- Extracting package stream (up xpkg xp-extract) ---"
    tmp_gz="$(mktemp)"
    trap "rm -f '$tmp_gz'" EXIT
    if up xpkg xp-extract --from-xpkg "$xpkg" -o "$tmp_gz" 2>/dev/null; then
      echo "Extract OK. Checking for TrustTunnelCloudflared in package stream..."
      if zcat "$tmp_gz" 2>/dev/null | grep -q "trusttunnelcloudflared"; then
        echo "  FOUND: trusttunnelcloudflared appears in package stream (CRD will be installed)."
      else
        echo "  MISSING: trusttunnelcloudflared does NOT appear in package stream (CRD omitted from xpkg)."
        echo "  This indicates the parser or encoder dropped it; check Crossplane CLI version and package root contents (run pre diagnostic)."
        exit 1
      fi
      echo ""
      echo "--- Sample CRD names in stream ---"
      zcat "$tmp_gz" 2>/dev/null | grep -oE 'name: [a-z0-9.]+' | grep -E 'cloudflare|upbound' | sort -u | head -30
      echo "  ..."
      echo ""
      echo "=== post diagnostic done ==="
      return
    fi
  fi

  # Try 2: OCI tarball inspection (tar + jq to find base layer, then extract and grep)
  echo "--- Inspecting OCI image tarball (crane/tar+jq) ---"
  if command -v crane &>/dev/null; then
    # crane export <tarball> writes the image as a tar to stdout; we need the layer with package.yaml
    # Simpler: crane manifest to get config and layers, then crane blob to get a layer and tar -t
    echo "Using crane to list image contents..."
    tmpdir="$(mktemp -d)"
    trap "rm -rf '$tmpdir'" EXIT
    # Load tarball as image ref (crane can read from file with tar: prefix in some versions)
    if crane manifest "tar:${xpkg}" 2>/dev/null | head -1 > "$tmpdir/manifest.json"; then
      # Find layer with io.crossplane.xpkg annotation = base
      base_layer="$(jq -r '.annotations["io.crossplane.xpkg"] // empty | select(. == "base")' "$tmpdir/manifest.json" 2>/dev/null)" || true
      # Actually annotations are per-layer; get layers and config
      layers=($(jq -r '.layers[].digest' "$tmpdir/manifest.json" 2>/dev/null))
      config_digest="$(jq -r '.config.digest' "$tmpdir/manifest.json" 2>/dev/null)"
      # Config has Labels; one of them is io.crossplane.xpkg: base for the layer index
      # This is getting complex. Fallback: just try to extract all layers and grep.
      for layer in "${layers[@]}"; do
        crane blob "tar:${xpkg}" "$layer" 2>/dev/null | tar -t 2>/dev/null | head -5
      done
    fi
    echo "  (crane inspection attempted)"
  fi

  # Fallback: raw tar list
  echo "--- Raw tarball contents (first 40 lines) ---"
  tar -tf "$xpkg" 2>/dev/null | head -40
  echo ""
  echo "  To fully verify, run: make xpkg.build.provider-cloudflare && devbox run -- up xpkg xp-extract --from-xpkg _output/xpkg/linux_amd64/provider-cloudflare-*.xpkg -o - | zcat | grep -c trusttunnelcloudflared"
  echo "=== post diagnostic done ==="
}

case "$phase" in
  pre)  pre ;;
  post) post "$@" ;;
  *)
    echo "Usage: $0 pre   (requires XPKG_DIR)"
    echo "       $0 post [path/to/package.xpkg]"
    exit 1
    ;;
esac
