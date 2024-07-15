#!/bin/sh

set -ex

if [ "$#" -lt 1 ] || [ "$#" -gt 2 ]; then
    echo "Usage: $0 <platform> [version]"
    echo "platform: must be one of 'ghes', 'ghec', or 'dotcom'"
    echo "version: a currently supported version of GitHub Enterprise Server, required only for 'ghes' and omitted otherwise"
    exit 1
fi

PLATFORM=$1
VERSION=${2:-""} # version defaults to an empty string if not provided

if [ "$PLATFORM" != "ghes" ] && [ "$PLATFORM" != "ghec" ] && [ "$PLATFORM" != "dotcom" ]; then
    echo "Invalid platform: must be one of 'ghes', 'ghec', or 'dotcom'."
    exit 1
fi

if [ "$PLATFORM" = "ghes" ]; then
	# Note that this logic only validates that the version is a positive number
    if [ "$VERSION" = "" ] || ! echo "$VERSION" | grep -Eq '^[0-9]+(\.[0-9]+)*$'; then
        echo "Invalid version for 'ghes': must be a currently supported version of GitHub Enterprise Server."
        exit 1
    fi
else
    if [ "$VERSION" != "" ]; then
        echo "Invalid version. Must be an empty string for platforms other than 'ghes'."
        exit 1
    fi
fi

SCHEMA_FILE="schemas/api.github.com.json"
if [ "$PLATFORM" != "dotcom" ]; then
    if [ "$VERSION" != "" ]; then
        SCHEMA_FILE="schemas/$PLATFORM-$VERSION.json"
    else
        SCHEMA_FILE="schemas/$PLATFORM.json"
    fi
fi

./scripts/install-tools.sh

if [ "$PLATFORM" = "ghec" ]; then
	NAMESPACE="go-sdk-enterprise-cloud"
elif [ "$PLATFORM" = "ghes" ]; then
	NAMESPACE="go-sdk-enterprise-server"
else
	NAMESPACE="go-sdk"
fi

go run schemas/main.go --schema-next=false --platform=$PLATFORM --version=$VERSION
kiota generate -l go --ll trace -o $(pwd)/stage/go/$NAMESPACE/pkg/github -n github.com/octokit/$NAMESPACE/pkg/github -d $SCHEMA_FILE --ebc
go build -o post-processors/go/post-processor post-processors/go/main.go
post-processors/go/post-processor $(pwd)/stage/go/$NAMESPACE
cd $(pwd)/stage/go/$NAMESPACE
go build ./...
