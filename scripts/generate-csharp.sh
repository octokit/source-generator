#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l csharp --ll trace -o generated/csharp/Octokit -c OctokitClient -n Octokit.Client -d schemas/downloaded.json --co

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
$(pwd)/post-processors/csharp/post-processor $(pwd)/generated/csharp/Octokit

# Build and run the generated code and call the API via generated SDK
cd generated/csharp
dotnet restore
dotnet run --project csharp.csproj