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
