package main

import (
	"context"
	"fmt"
	"log"
	"os"

	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
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

	cat, err := client.Octocat().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))

	// uncomment to make authenticated request for private user emails
	// x, err := client.User().Emails().Get(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }

	// for _, v := range x {
	// 	fmt.Printf("%v\n", *v.GetEmail())
	// }
}
