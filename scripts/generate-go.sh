#!/bin/sh

# convenience script for local development only
# .github/workflows/build-go.yml is the source of truth for the generated Go workflow
# note: before running this script, ensure the generated/go directory has been deleted
# but the sample app restored:
# rm -rf generated/go/
# git restore generated/

source ./scripts/install-tools.sh

set -ex

go run schemas/main.go --schema-next=false
kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/downloaded.json
go build -o post-processors/go/post-processor post-processors/go/main.go
post-processors/go/post-processor $(pwd)/generated/go
cd generated/go
go build ./...
go run *.go
