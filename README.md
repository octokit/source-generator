# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Usage

### For Kiota

#### Via post-processing

1. Install Kiota: First install .NET 7, then run `dotnet tool install --global --prerelease Microsoft.OpenApi.Kiota`

##### GO

1. Run generation: `kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Build post-processor: `go build -o post-processors/go/post-processor post-processors/go/main.go`
1. Run post-processor: `post-processors/go/post-processor $(pwd)/generated/go`
1. Run `go build ./...` from the `generated/go` directory to build the project.
<!-- TODO(kfcampbell): create main.go file and run it -->

##### C#

1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp/Octokit -n Octokit -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Build post-processor: `go build -o post-processors/csharp/post-processor post-processors/csharp/main.go`
1. Run post-processor: `post-processors/csharp/post-processor $(pwd)/generated/csharp/Octokit`
1. Run `dotnet run` from the directory in which you've output the generated code.

#### Manually

1. Install Kiota: First install .NET 7, then run `dotnet tool install --global --prerelease Microsoft.OpenApi.Kiota`

##### GO

1. Run generation: `kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go mod init github.com/octokit/kiota`
1. There's a command to print each dependency that's required; run it and then install dependencies: `kiota info -d /home/kfcampbell/github/dev/source-generator/schemas/updated/api.github.com.json -l Go`
1. Run `go build ./...` from the directory in which you've output the generated code to see errors. Fix build errors until build is working.
1. For more info, [see](https://microsoft.github.io/kiota/get-started/go.html)

##### C#

1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp/Octokit -n Octokit -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `dotnet run` from the directory in which you've output the generated code to see errors. Fix build errors until build is working.
1. For more info, [see](https://microsoft.github.io/kiota/get-started/dotnet.html)