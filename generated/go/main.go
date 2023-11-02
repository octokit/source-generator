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
	"github.com/octokit/kiota/user"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalf("GITHUB_TOKEN must be provided")
	}

	tokenProvider, err := auth.NewApiKeyAuthenticationProvider(fmt.Sprintf("Bearer %s", token), "Authorization", auth.HEADER_KEYLOCATION)
	if err != nil {
		log.Fatalf("failed to initialize token provider: %v", err)
	}

	adapter, err := http.NewNetHttpRequestAdapter(tokenProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}

	client := NewApiClient(adapter)
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	// unauthenticated request
	s := "Salutations"
	octocatRequestConfig := &octocat.OctocatRequestBuilderGetRequestConfiguration{
		QueryParameters: &octocat.OctocatRequestBuilderGetQueryParameters{
			S: &s,
		},
		Headers: headers,
	}
	cat, err := client.Octocat().Get(context.Background(), octocatRequestConfig)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))

	// authenticated request for private user emails
	emailsRequestConfig := &user.EmailsRequestBuilderGetRequestConfiguration{
		Headers: headers,
	}
	userEmails, err := client.User().Emails().Get(context.Background(), emailsRequestConfig)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for _, v := range userEmails {
		fmt.Printf("%v\n", *v.GetEmail())
	}
}
