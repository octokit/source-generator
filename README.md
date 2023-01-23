# source-generator

This repository is a prototype of code generation from GitHub's OpenAPI specification.

## Usage

### Prerequisites
1. Install openapi-generator-cli: `npm i -g @openapitools/openapi-generator-cli`
1. Install the Java runtime (on Ubuntu derivatives, you may use `sudo apt install default-jre`)

## Execution

1. To generate the .NET package, run from the root of the repo: `openapi-generator-cli generate -i schemas/api.github.com/api.github.com.2022-11-28.json -t templates/dotnet -g csharp -o generated/dotnet -p packageName=Octokit`. Alternately, you may run the provided script `generate.sh`.