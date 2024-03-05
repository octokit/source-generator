#!/usr/bin/env bash

# GitHub Actions runs linting checks from .github/workflows/lint.yml.
# This script may be used to detect issues before pushing to a PR check.
set -ex
./scripts/install-tools.sh
golangci-lint run post-processors/go/main.go
golangci-lint run post-processors/csharp/main.go
golangci-lint run schemas/main.go
