package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
	"github.com/octokit/kiota"
	"github.com/octokit/kiota/gists"
	"github.com/octokit/kiota/octocat"
)

func main() {
	authProvider, err := authentication.NewApiKeyAuthenticationProvider(fmt.Sprintf("Bearer %v", os.Getenv("GITHUB_TOKEN")), "Authorization", authentication.HEADER_KEYLOCATION)
	if err != nil {
		log.Fatalf("error creating auth provider: %v", err)
	}

	ra, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		log.Fatalf("error creating requestAdapter: %v", err)
	}

	client := kiota.NewApiClient(ra)

	words := "test string for octocat to say"
	params := &octocat.OctocatRequestBuilderGetQueryParameters{
		S: &words,
	}
	reqConfig := &octocat.OctocatRequestBuilderGetRequestConfiguration{
		QueryParameters: params,
	}

	res, err := client.Octocat().Get(context.TODO(), reqConfig)
	if err != nil {
		log.Fatalf("error getting request: %v", err)
	}
	log.Printf("octocat result: %v", string(res))

	gistParams := &gists.GistsRequestBuilderGetQueryParameters{}

	gistsConfig := &gists.GistsRequestBuilderGetRequestConfiguration{
		QueryParameters: gistParams,
	}

	gistsBuilder := client.Gists()
	allGists, err := gistsBuilder.Get(context.TODO(), gistsConfig)
	if err != nil {
		log.Fatalf("error getting gists: %v", err)
	}

	for _, g := range allGists {
		log.Printf("gist description: %v, public: %v", *g.GetDescription(), *g.GetPublic())
	}

}
