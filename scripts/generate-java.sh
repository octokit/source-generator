#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
# The excluded path is a workaround for: https://github.com/microsoft/kiota/issues/4151

kiota generate -l java --ll trace \
  -o generated/Java/src/main/java/com/github/api \
  -c GitHubClient \
  -n com.github.api \
  -d schemas/downloaded.json \
  --serializer none \
  --deserializer none \
  --exclude-path "/repos/{template_owner}/{template_repo}/generate" \
  --co --ebc
