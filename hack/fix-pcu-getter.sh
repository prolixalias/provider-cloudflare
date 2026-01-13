#!/bin/bash
# Post-generation fix: angryjet generates Get/SetProviderConfigReference with xpv1.Reference,
# but resource.TypedProviderConfigUsage requires xpv1.ProviderConfigReference.
# Run after make generate so the fix persists.

set -e

PCU_FILE="apis/namespaced/v1beta1/zz_generated.pcu.go"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$ROOT_DIR"

[ -f "$PCU_FILE" ] || { echo "Skip: $PCU_FILE not found"; exit 0; }

if ! grep -q 'GetProviderConfigReference() xpv1.Reference' "$PCU_FILE"; then
  echo "Skip: $PCU_FILE already has correct GetProviderConfigReference return type"
  exit 0
fi

python3 << 'PY'
path = "apis/namespaced/v1beta1/zz_generated.pcu.go"
old = """// GetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) GetProviderConfigReference() xpv1.Reference {
	return p.ProviderConfigReference
}"""
new = """// GetProviderConfigReference of this ProviderConfigUsage.
// Returns xpv1.ProviderConfigReference so that *ProviderConfigUsage implements resource.TypedProviderConfigUsage.
// (xpv1.Reference is a value type with Name only; ProviderConfigReference has Kind and Name.)
func (p *ProviderConfigUsage) GetProviderConfigReference() xpv1.ProviderConfigReference {
	ref := p.ProviderConfigReference
	if ref.Name == "" {
		return xpv1.ProviderConfigReference{}
	}
	return xpv1.ProviderConfigReference{
		Kind: "ProviderConfig",
		Name: ref.Name,
	}
}"""
with open(path) as f:
    content = f.read()
if old not in content:
    raise SystemExit("Pattern not found in " + path)
content = content.replace(old, new)

old_set = """// SetProviderConfigReference of this ProviderConfigUsage.
func (p *ProviderConfigUsage) SetProviderConfigReference(r xpv1.Reference) {
	p.ProviderConfigReference = r
}"""
new_set = """// SetProviderConfigReference of this ProviderConfigUsage.
// Accepts xpv1.ProviderConfigReference so that *ProviderConfigUsage implements resource.TypedProviderConfigUsage.
func (p *ProviderConfigUsage) SetProviderConfigReference(r xpv1.ProviderConfigReference) {
	p.ProviderConfigReference = xpv1.Reference{Name: r.Name}
}"""
if old_set in content:
    content = content.replace(old_set, new_set)

with open(path, "w") as f:
    f.write(content)
print("Fixed GetProviderConfigReference and SetProviderConfigReference in", path)
PY

exit 0
