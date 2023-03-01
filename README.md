# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Development

- Install
	- [.NET 7](https://dotnet.microsoft.com/en-us/download/dotnet/7.0)
	- [Latest version of Go](https://go.dev/dl/)
	- Kiota: `dotnet tool install --global --prerelease Microsoft.OpenApi.Kiota`
- Validation of SDKs in other platforms such as Ruby require that platform to be installed as well

## Usage

### Go

1. Run generation: `kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.A
1. Run `go mod init github.com/octokit/kiota`
1. There's a command to print each dependency that's required; run it and then install dependencies: `kiota info -d /home/kfcampbell/github/dev/source-generator/schemas/updated/api.github.com.json -l Go`
1. Run `go build ./...` from the directory in which you've output the generated code to see errors. Fix build errors until build is working.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/go.html).
<!-- TODO(kfcampbell): create main.go file and run it -->

### C#

1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `dotnet run` from the directory in which you've output the generated code to see errors. Fix build errors until build is working.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/dotnet.html).

### Ruby

1. Run generation: `kiota generate -l ruby --ll trace -o generated/ruby -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. There's a command to print each dependency that's required; run it and then install dependencies in the generated/go: `kiota info -d /home/kfcampbell/github/dev/source-generator/schemas/updated/api.github.com.json -l Ruby`

