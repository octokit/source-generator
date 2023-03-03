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
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go build -o post-processors/go/post-processor post-processors/go/main.go` to build the post-processor.
1. Run `post-processors/go/post-processor $(pwd)/generated/go` to execute the post-processor.
1. Run `go build ./...` from the directory in which you've output the generated code to see errors.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/go.html).
<!-- TODO(kfcampbell): create main.go file and run it -->

### C#

1. Run generation: `kiota generate -l csharp --ll trace -o generated/csharp -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go build -o post-processors/csharp/post-processor post-processors/csharp/main.go` to build the post-processor.
1. Run `post-processors/csharp/post-processor $(pwd)/generated/csharp` to execute the post-processor.
1. Run `dotnet build` from the directory in which you've output the generated code to see errors. Fix build errors until build is working.
1. For more info, [see Kiota documentation](https://microsoft.github.io/kiota/get-started/dotnet.html).

### Ruby

1. Run generation: `kiota generate -l ruby --ll trace -o generated/ruby -n kiota -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. There's a command to print each dependency that's required; run it and then install dependencies in the generated/go: `kiota info -d /home/kfcampbell/github/dev/source-generator/schemas/updated/api.github.com.json -l Ruby`
1. Run `go build -o post-processors/ruby/post-processor post-processors/ruby/main.go` to build the post-processor.
1. Run `post-processors/ruby/post-processor $(pwd)/generated/ruby` to execute the post-processor.
1. You may import and test the library by running `irb` and `load 'api_client.rb'`.

Work-in-progress example Ruby code:

```ruby
require "microsoft_kiota_authentication_oauth"
require "microsoft_graph_core"
require "microsoft_graph"
require_relative 'api_client.rb'

context = MicrosoftKiotaAuthenticationOAuth::ClientCredentialContext.new("<the tenant id from your app registration>", "<the client id from your app registration>", "<the client secret from your app registration>")
authentication_provider = MicrosoftGraphCore::Authentication::OAuthAuthenticationProvider.new(context, nil, ["Files.Read"])
adapter = MicrosoftGraphBeta::GraphRequestAdapter.new(authentication_provider)

client = Kiota::ApiClient.new(adapter)

octocat = client.octocat.get()
p octocat
```
