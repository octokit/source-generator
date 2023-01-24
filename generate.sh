#!/usr/bin/env sh

export CSHARP_POST_PROCESS_FILE="/home/kfcampbell/github/dev/source-generator/csharp-post-process.sh"

openapi-generator-cli generate -i schemas/api.github.com/api.github.com.2022-11-28.json -t templates/dotnet -g csharp -o generated/dotnet -p packageName=Octokit --enable-post-process-file


