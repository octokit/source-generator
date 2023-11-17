#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l csharp --ll trace -o generated/csharp/GitHub/Octokit -c OctokitClient -n GitHub.Octokit -d schemas/downloaded.json --co

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
$(pwd)/post-processors/csharp/post-processor $(pwd)/generated/csharp/GitHub/Octokit

# Build and run the generated code and call the API via generated SDK
cd generated/csharp/GitHub
dotnet restore
dotnet run --project GitHub.Octokit.csproj