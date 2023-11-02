package main

import (
	"context"
	"fmt"
	"log"
	"os"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
	"github.com/octokit/kiota/octocat"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalf("GITHUB_TOKEN must be provided")
	}

	tokenProvider, err := auth.NewApiKeyAuthenticationProvider(token, "authorization", auth.HEADER_KEYLOCATION)
	if err != nil {
		log.Fatalf("failed to initialize token provider: %v", err)
	}

	adapter, err := http.NewNetHttpRequestAdapter(tokenProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}

	client := NewApiClient(adapter)

	// unauthenticated request
	s := "Salutations"
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")
	requestConfig := &octocat.OctocatRequestBuilderGetRequestConfiguration{
		QueryParameters: &octocat.OctocatRequestBuilderGetQueryParameters{
			S: &s,
		},
		Headers: headers,
	}
	cat, err := client.Octocat().Get(context.Background(), requestConfig)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))

	// authenticated request for private user emails
	// userEmails, err := client.User().Emails().Get(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }

	// for _, v := range userEmails {
	// 	fmt.Printf("%v\n", *v.GetEmail())
	// }
}
