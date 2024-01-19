#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l csharp --ll trace -o generated/csharp/src/GitHub -c GitHubClient -n GitHub -d schemas/downloaded.json --co --ebc

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
$(pwd)/post-processors/csharp/post-processor $(pwd)/generated/csharp/src/GitHub
