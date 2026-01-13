#!/usr/bin/env bash
# Run CI-style checks locally (generate, build, lint, test, validate).
# Usage:
#   ./scripts/test-ci-local.sh           # full CI pipeline
#   ./scripts/test-ci-local.sh --quick   # skip check-diff (faster iteration)
#   ./scripts/test-ci-local.sh --check-diff  # run only check-diff after other steps
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$ROOT_DIR"

RUN_CHECK_DIFF=true
for arg in "$@"; do
  case "$arg" in
    --quick)      RUN_CHECK_DIFF=false ;;
    --check-diff) RUN_CHECK_DIFF=true ;;
  esac
done

run() {
  if command -v devbox >/dev/null 2>&1; then
    devbox run -- "$@"
  else
    "$@"
  fi
}

section() {
  echo ""
  echo "=== $* ==="
}

section "Submodules"
run make submodules

section "Generate"
run make generate

section "Build"
run make build

section "Go validate (modules, vet, fmt)"
run make go.modules.check
run make vet
run make fmt

section "Lint"
run make lint

section "Unit tests"
run make test

if [ "$RUN_CHECK_DIFF" = true ]; then
  section "Check diff (generate produces no uncommitted changes)"
  run make check-diff
fi

section "Done"
echo "All CI-style checks passed."
if [ "$RUN_CHECK_DIFF" = false ]; then
  echo "Tip: run without --quick to also verify 'make check-diff' (no generated diff)."
fi
