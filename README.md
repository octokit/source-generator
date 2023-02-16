# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Usage

### For AutoREST

1. `npm install -g autorest`
1. `time autorest autorest-config.yaml --go --debug --verbose` from the repo root for Go code generation
1. Run post-processing with `cd go-post-processor && go run main.go`
1. Build resulting code with `cd output && go build ./...`

Notes:

- Config options are the same as CLI versions, except for `--debug` and `--verbose` which are only able to be applied through the API.

### For Kiota

1. Install Kiota: First install .NET 7, then run `dotnet tool install --global --prerelease Microsoft.OpenApi.Kiota`
1. Run generation: `kiota generate -l go --ll trace -o generated/go -n octokit -d schemas/updated/api.github.com.json --co > output.txt 2>&1`
	1. Alternately, you may debug using the VSCode launch.json in the microsoft/kiota repo.
1. Run `go mod init github.com/octokit/octokit-kiota`
1. There's a command to print each dependency that's required; run that and then install dependencies
1. Run `go build ./...` from the directory in which you've output the generated code to see errors.

To run pre-fixed/compiling Kiota demo app:
1. Export your PAT as GITHUB_TOKEN
1. `cd` to `compiling-kiota-go/go`
1. Run `go run cmd/main.go`
