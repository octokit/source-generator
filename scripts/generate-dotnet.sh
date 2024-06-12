#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l csharp --ll trace -o $(pwd)/stage/dotnet-sdk/src/GitHub -c GitHubClient -n GitHub -d schemas/api.github.com.json --ebc

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
post-processors/csharp/post-processor $(pwd)/stage/dotnet-sdk/src/GitHub
cd stage/dotnet-sdk
dotnet build
