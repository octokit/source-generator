#!/bin/sh

./scripts/install-tools.sh

go run schemas/main.go --schema-next=false

# Execute generation
kiota generate -l java --ll trace -o generated/Java/src/main/java/com/github/api -c GitHubClient -n com.github.api -d schemas/downloaded.json --serializer none --deserializer none --co --ebc
