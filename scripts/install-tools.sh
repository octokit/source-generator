#!/usr/bin/env bash

# note: check the version of golangci-lint used in .github/workflows/lint.yml to ensure accuracy
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

# Install Kiota
# This version can only increase, never decrease. If you want to use a lower version, you need to uninstall the tool first.
# dotnet tool uninstall Microsoft.OpenApi.Kiota --global
dotnet tool install --global Microsoft.OpenApi.Kiota
