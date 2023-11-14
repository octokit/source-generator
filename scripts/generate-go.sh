#!/bin/sh

# convenience script for local development only
# .github/workflows/build-go.yml is the source of truth for the generated Go workflow

./scripts/install-tools.sh

set -ex

go run schemas/main.go --schema-next=false
kiota generate -l go --ll trace -o $(pwd)/../go-sdk -n go-sdk -d schemas/downloaded.json
go build -o post-processors/go/post-processor post-processors/go/main.go
cd $(pwd)/../go-sdk
$(pwd)/../source-generator/post-processors/go/post-processor $(pwd)
go build ./...
