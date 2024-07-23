# source-generator

This repository generates code from [GitHub's OpenAPI specification](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota). If you're looking for the repositories generated from this code, you may want:

- Go
	- For the standard GitHub.com product
		- [go-sdk repository](https://github.com/octokit/go-sdk)
		- [pkg.go.dev docs link](https://pkg.go.dev/github.com/octokit/go-sdk)
	- For GitHub Enterprise Cloud
		- [go-sdk-enterprise-cloud repository](https://github.com/octokit/go-sdk-enterprise-cloud)
		- [pkg.go.dev docs link](https://pkg.go.dev/github.com/octokit/go-sdk-enterprise-cloud)
	- For GitHub Enterprise Server
		- [go-sdk-enterprise-server repository](https://github.com/octokit/go-sdk-enterprise-server)
		- [pkg.go.dev docs link](https://pkg.go.dev/github.com/octokit/go-sdk-enterprise-server)
- .NET
	- For the standard GitHub.com product
		- [dotnet-sdk repository](https://github.com/octokit/dotnet-sdk)
		- [NuGet link](https://www.nuget.org/packages/GitHub.Octokit.SDK)
	- For GitHub Enterprise Cloud
		- [dotnet-sdk-enterprise-cloud repository](https://github.com/octokit/dotnet-sdk-enterprise-cloud)
	- For GitHub Enterprise Server
		- [dotnet-sdk-enterprise-server repository](https://github.com/octokit/dotnet-sdk-enterprise-server)

## Development

- Install
	- [.NET 7](https://dotnet.microsoft.com/en-us/download/dotnet/7.0)
	- [Latest version of Go](https://go.dev/dl/)
	- Kiota: `dotnet tool install --global Microsoft.OpenApi.Kiota --prerelease`
- For the latest versions of tooling, the `.github/workflows/build-go.yml` file is the source of truth.

## Usage

### Go

Run `./scripts/generate-go.sh`. You may also run individual steps from the `.github/workflows/build-go.yml` file. If you're using VSCode, there are some pre-setup tasks and debugging configurations in the `.vscode` folder.

### C#

Run `./scripts/generate-csharp.sh`
