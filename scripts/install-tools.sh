#!/usr/bin/env bash

# note: check the version of golangci-lint used in .github/workflows/lint.yml to ensure accuracy
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56

# Install Kiota
KIOTA_TOOL_NAME="microsoft.openapi.kiota"
# This version can only increase, never decrease. If you want to use a lower version, you need to uninstall the tool first.
# dotnet tool uninstall Microsoft.OpenApi.Kiota --global
KIOTA_TOOL_VERSION="v1.9.0-preview.202311300001"

# Check if the tool is installed globally
if dotnet tool list -g | grep -q $KIOTA_TOOL_NAME; then
  if dotnet tool list -g | grep -q "$KIOTA_TOOL_NAME.*$KIOTA_TOOL_VERSION"; then
    echo "The correct version ($KIOTA_TOOL_VERSION) of $KIOTA_TOOL_NAME is installed...skipping install step"
  else
      echo "$KIOTA_TOOL_NAME is installed. Updating to sync with the required version"
      dotnet tool update --global Microsoft.OpenApi.Kiota --prerelease
  fi
else
  echo "$KIOTA_TOOL_NAME is not installed...installing"
  dotnet tool install --global Microsoft.OpenApi.Kiota --prerelease
fi

