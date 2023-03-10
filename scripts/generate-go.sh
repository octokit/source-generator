#!/bin/sh

# Execute generation
kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/updated/api.github.com.json --co

# Build and run post-processor
go build -o $(pwd)post-processors/go/post-processor post-processors/go/main.go
$(pwd)post-processors/go/post-processor $(pwd)/generated/go

# Build and run the generated code and call the API via generated SDK
cd generated/go
go build ./...