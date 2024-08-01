package pkg_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	"github.com/octokit/go-sdk-enterprise-server/pkg"
	auth "github.com/octokit/go-sdk-enterprise-server/pkg/authentication"
	"github.com/octokit/go-sdk-enterprise-server/pkg/github"
	"github.com/octokit/go-sdk-enterprise-server/pkg/github/octocat"
)

// ExampleApiClient_Octocat constructs an unauthenticated API client
// and makes a simple API request.
func ExampleApiClient_Octocat() {
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("octokit/go-sdk.example-functions"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
	)

	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	s := "Salutations"

	// create headers that accept json back; GitHub's OpenAPI definition says
	// octet-stream but that's not actually what the API returns in this case
	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	octocatRequestConfig := &abs.RequestConfiguration[octocat.OctocatRequestBuilderGetQueryParameters]{
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
	// Output:
	// MMM.           .MMM
	//                MMMMMMMMMMMMMMMMMMM
	//                MMMMMMMMMMMMMMMMMMM      _____________
	//               MMMMMMMMMMMMMMMMMMMMM    |             |
	//              MMMMMMMMMMMMMMMMMMMMMMM   | Salutations |
	//             MMMMMMMMMMMMMMMMMMMMMMMM   |_   _________|
	//             MMMM::- -:::::::- -::MMMM    |/
	//              MM~:~ 00~:::::~ 00~:~MM
	//         .. MMMMM::.00:::+:::.00::MMMMM ..
	//               .MM::::: ._. :::::MM.
	//                  MMMM;:::::;MMMM
	//           -MM        MMMMMMM
	//           ^  M+     MMMMMMMMM
	//               MMMMMMM MM MM MM
	//                    MM MM MM MM
	//                    MM MM MM MM
	//                 .~~MM~MM~MM~MM~~.
	//              ~~~~MM:~MM~~~MM~:MM~~~~
	//             ~~~~~~==~==~~~==~==~~~~~~
	//              ~~~~~~==~==~==~==~~~~~~
	//                  :~==~==~==~==~~
}

// ExampleApiClient_Octocat_withoutConvenienceConstructor shows how to
// initialize an unauthenticated API client without using the helpers
// provided in pkg/client.go, then makes a simple API request.
func ExampleApiClient_Octocat_withoutConvenienceConstructor() {
	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithTokenAuthentication("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)
	adapter, err := kiotaHttp.NewNetHttpRequestAdapter(tokenProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}

	// TODO: Rework this test to fit the platform needs of GHES
	// possibly add APIs the GHES SDK to provide base URL values to the formatter
	baseUrl := adapter.GetBaseUrl()

    // Note: This SDK should be used against a GitHub Enterprise instance, and the below URL is the public GitHub one. It's here only so that tests pass when running `go test ./...`, as the OpenAPI schema for GHES understandably does not include a baseURL.
    // When setting up this package for your own usage, call `adapter.SetBaseUrl` or `pkg.WithBaseUrl` with your own GHES base URL.
	if baseUrl == "" {
		adapter.SetBaseUrl("https://api.github.com")
	}

	client := github.NewApiClient(adapter)

	s := "Salutations"

	// create headers that accept json back; GitHub's OpenAPI definition says
	// octet-stream but that's not actually what the API returns in this case
	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	octocatRequestConfig := &abs.RequestConfiguration[octocat.OctocatRequestBuilderGetQueryParameters]{
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
	// Output:
	// MMM.           .MMM
	//                MMMMMMMMMMMMMMMMMMM
	//                MMMMMMMMMMMMMMMMMMM      _____________
	//               MMMMMMMMMMMMMMMMMMMMM    |             |
	//              MMMMMMMMMMMMMMMMMMMMMMM   | Salutations |
	//             MMMMMMMMMMMMMMMMMMMMMMMM   |_   _________|
	//             MMMM::- -:::::::- -::MMMM    |/
	//              MM~:~ 00~:::::~ 00~:~MM
	//         .. MMMMM::.00:::+:::.00::MMMMM ..
	//               .MM::::: ._. :::::MM.
	//                  MMMM;:::::;MMMM
	//           -MM        MMMMMMM
	//           ^  M+     MMMMMMMMM
	//               MMMMMMM MM MM MM
	//                    MM MM MM MM
	//                    MM MM MM MM
	//                 .~~MM~MM~MM~MM~~.
	//              ~~~~MM:~MM~~~MM~:MM~~~~
	//             ~~~~~~==~==~~~==~==~~~~~~
	//              ~~~~~~==~==~==~==~~~~~~
	//                  :~==~==~==~==~~
}
