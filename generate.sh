#!/bin/sh

openapi-generator-cli generate -i schemas/api.github.com/api.github.com.2022-11-28.json -t templates/dotnet -g csharp -o generated/dotnet