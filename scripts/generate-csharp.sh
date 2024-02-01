#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l csharp --ll trace -o $(pwd)/../dotnet-sdk/src/GitHub -c GitHubClient -n GitHub -d schemas/downloaded.json --ebc

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
post-processors/csharp/post-processor $(pwd)/../dotnet-sdk/src/GitHub
cd ../dotnet-sdk
dotnet build