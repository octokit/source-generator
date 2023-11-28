# source-generator

This repository generates code from [GitHub's OpenAPI specification](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota).

## Development

- Install
	- [.NET 7](https://dotnet.microsoft.com/en-us/download/dotnet/7.0)
	- [Latest version of Go](https://go.dev/dl/)
	- Kiota: `dotnet tool install --global Microsoft.OpenApi.Kiota --version 1.8.2`
- For the latest versions of tooling, the `.github/workflows/build-go.yml` file is the source of truth.

## Usage

### Go

Run `./scripts/generate-go.sh`. You may also run individual steps from the `.github/workflows/build-go.yml` file. If you're using VSCode, there are some pre-setup tasks and debugging configurations in the `.vscode` folder.

### C#

1. Download the latest schema: `go run schemas/main.go --schema-next=false`
1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp -n Octokit -d schemas/downloaded.json > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go build -o post-processors/csharp/post-processor post-processors/csharp/main.go` to build the post-processor.
1. Run `post-processors/csharp/post-processor $(pwd)/generated/csharp` to execute the post-processor.
1. Run `dotnet build` from the directory in which you've output the generated code to check compilation.
1. Run `dotnet run` from the directory in which you've output the generated code in order to run the sample application.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/dotnet.html).
