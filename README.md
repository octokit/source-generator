# source-generator

This repository generates code from [GitHub's OpenAPI specification](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota).

## Development

- Install
	- [.NET 7](https://dotnet.microsoft.com/en-us/download/dotnet/7.0)
	- [Latest version of Go](https://go.dev/dl/)
	- Kiota: `dotnet tool install --global Microsoft.OpenApi.Kiota --version 1.8.1`
- For the latest versions of tooling, the `.github/workflows/build-go.yml` file is the source of truth.

## Usage

### Go

Run `./scripts/generate-go.sh`. You may also run individual steps from the `.github/workflows/build-go.yml` file. If you're using VSCode, there are some pre-setup tasks and debugging configurations in the `.vscode` folder.

### C#

Run `./scripts/generate-csharp.sh`
