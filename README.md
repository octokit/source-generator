# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Development

- Install
	- [.NET 7](https://dotnet.microsoft.com/en-us/download/dotnet/7.0)
	- [Latest version of Go](https://go.dev/dl/)
	- Kiota: `dotnet tool install --global Microsoft.OpenApi.Kiota --version 1.1.3`
- Validation of SDKs in other platforms such as Ruby require that platform to be installed as well

## Usage

### Go

1. Download the latest schema: `go run schemas/main.go --schema-next=false`
1. Run generation: `kiota generate -l go --ll trace -o generated/go -n kiota -d schemas/downloaded.json > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go build -o post-processors/go/post-processor post-processors/go/main.go` to build the post-processor.
1. Run `post-processors/go/post-processor $(pwd)/generated/go` to execute the post-processor.
1. Run `go build ./...` from the directory in which you've output the generated code to check compilation.
1. Run `go run main.go` from the directory in which you've output the generated code in order to run the sample application.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/go.html).
<!-- TODO(kfcampbell): create main.go file and run it -->

### C#

1. Download the latest schema: `go run schemas/main.go --schema-next=false`
1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp -n kiota -d schemas/downloaded.json > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go build -o post-processors/csharp/post-processor post-processors/csharp/main.go` to build the post-processor.
1. Run `post-processors/csharp/post-processor $(pwd)/generated/csharp` to execute the post-processor.
1. Run `dotnet build` from the directory in which you've output the generated code to check compilation.
1. Run `dotnet run` from the directory in which you've output the generated code in order to run the sample application.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/dotnet.html).
