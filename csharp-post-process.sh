#!/usr/bin/env sh

filename=$1

echo "Running post-process script on $filename"

go run post-process/cmd/dotnet/main.go $filename
