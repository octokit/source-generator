#!/bin/sh

# convenience script for local development only
# .github/workflows/build-go.yml is the source of truth for the generated Go workflow
# requires cloning github.com/octokit/go-sdk in a sibling directory

./scripts/install-tools.sh

set -ex

if [ -z "$1" ]; then 
  gen_path="../go-sdk"
  go run schemas/main.go --schema-next=false
elif [ "$1" = "ghes" ]; then
  if [ -z "$2" ]; then
    echo "generating for GHES requires that you specify a version as the second argument. ex $0 ghes 3.12"
    exit 1
  else
    ghes_version=$2
    gen_path="../go-sdk-ghes"
    go run schemas/main.go --schema-next=false --ghes=true --ghes-version=3.12
  fi
fi

kiota generate -l go --ll trace -o $(pwd)/${gen_path}/pkg/github -n github.com/octokit/go-sdk/pkg/github -d schemas/downloaded.json --ebc
go build -o post-processors/go/post-processor post-processors/go/main.go
cd $(pwd)/${gen_path}
$(pwd)/../source-generator/post-processors/go/post-processor $(pwd)
go build ./...
