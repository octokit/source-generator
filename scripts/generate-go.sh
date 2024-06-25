#!/bin/sh

./scripts/install-tools.sh

set -ex

go run schemas/main.go --schema-next=false
kiota generate -l go --ll trace -o $(pwd)/stage/go/go-sdk/pkg/github -n github.com/octokit/go-sdk/pkg/github -d schemas/api.github.com.json --ebc
go build -o post-processors/go/post-processor post-processors/go/main.go
post-processors/go/post-processor $(pwd)/stage/go/go-sdk
cd $(pwd)/stage/go/go-sdk
go build ./...
