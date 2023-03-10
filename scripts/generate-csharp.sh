#!/bin/sh

# Execute generation
kiota generate -l csharp --ll trace -o generated/csharp/Octokit -n Octokit -d schemas/updated/api.github.com.json --co

# Build and run post-processor
go build -o $(pwd)/post-processors/csharp/post-processor post-processors/csharp/main.go
$(pwd)/post-processors/csharp/post-processor $(pwd)/generated/csharp/Octokit

# Build and run the generated code and call the API via generated SDK
cd generated/csharp
dotnet run --project csharp.csproj