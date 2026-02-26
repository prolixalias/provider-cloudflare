#!/bin/bash
# Post-generation hook to fix CEL validation errors for dynamic fields
# This script removes validation rules for fields with x-kubernetes-preserve-unknown-fields: true

set -e

echo "Fixing CRD validation rules for dynamic fields..."

cd package/crds

# List of CRDs and their dynamic fields
declare -A crd_fields=(
  ["hostname.cloudflare.m.upbound.io_tlssettings.yaml"]="value"
  ["hostname.cloudflare.upbound.io_tlssettings.yaml"]="value"
  ["dns.cloudflare.m.upbound.io_records.yaml"]="flags"
  ["dns.cloudflare.upbound.io_records.yaml"]="flags"
  ["spectrum.cloudflare.m.upbound.io_applications.yaml"]="originPort"
  ["spectrum.cloudflare.upbound.io_applications.yaml"]="originPort"
  ["worker.cloudflare.m.upbound.io_versions.yaml"]="runWorkerFirst"
  ["worker.cloudflare.upbound.io_versions.yaml"]="runWorkerFirst"
  ["workers.cloudflare.m.upbound.io_scripts.yaml"]="runWorkerFirst"
  ["workers.cloudflare.upbound.io_scripts.yaml"]="runWorkerFirst"
  ["zone.cloudflare.m.upbound.io_settings.yaml"]="value"
  ["zone.cloudflare.upbound.io_settings.yaml"]="value"
)

for crd in "${!crd_fields[@]}"; do
  if [ -f "$crd" ]; then
    field="${crd_fields[$crd]}"
    echo "  Fixing $crd (field: $field)"
    # yq v4 (mikefarah) uses: yq eval -i '...' file
    # yq v3 (kislyuk) uses different syntax; skip if eval -i is not supported
    if yq eval -i "del(.spec.versions[].schema.openAPIV3Schema.properties.spec.x-kubernetes-validations[] | select(.message | contains(\"$field\")))" "$crd" 2>/dev/null; then
      : "fixed"
    else
      echo "  (yq fix skipped for $crd - need yq v4 for -i; CRD unchanged)"
    fi
  fi
done

echo "CRD validation fixes applied!"
