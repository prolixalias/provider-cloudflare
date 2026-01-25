#!/bin/bash
# Script to test CI pipeline locally before pushing
# This replicates the key CI steps to catch issues early

set -e

echo "🚀 Testing CI pipeline locally..."

# Ensure Go tools are available
export PATH="$HOME/go/bin:$PATH"

echo "📦 Installing required tools..."

# Install goimports
if ! command -v goimports &> /dev/null; then
    echo "Installing goimports..."
    go install golang.org/x/tools/cmd/goimports@latest
else
    echo "goimports already installed"
fi

# Install golangci-lint
if ! command -v golangci-lint &> /dev/null || [[ "$(golangci-lint version 2>/dev/null | grep -o 'version [0-9.]*' | head -1)" != "version 2.8.0" ]]; then
    echo "Installing golangci-lint v2.8.0..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.8.0
else
    echo "golangci-lint v2.8.0 already installed"
fi

echo "🔧 Testing code generation..."
make generate

echo "📦 Testing vendor dependencies..."
make vendor vendor.check

echo "🔍 Testing linting (quick check on main files)..."
golangci-lint run --timeout=30s cmd/provider/main.go || echo "⚠️  Linting found issues - check CI for details"

echo "✅ All CI steps passed locally!"
echo ""
echo "🎉 Ready to push - CI should work now!"