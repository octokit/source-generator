package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/nickfloyd/source-generator/generated/go/gh"
)

func main() {
	fmt.Printf("Hello world!\n")
	fmt.Printf("a random constant: %v\n", gh.WebhookPullRequestUnassignedActionUnassigned)
	fmt.Printf("another random constant: %v\n", gh.ReactionContentRocket)

	// create a client
	/*
		due to an undiagnosed-as-of-yet error, various module arguments are required at client creation time
			(*string, int32, int32, string, string, string, string, int32, int32, int32, string, string, int32, string, int32, int32, int32, string, string, string, int32, int32, string, int32, string, gh.Enum121, string, int32, string, int32, int32, int32, int32, gh.Enum157, gh.Enum158, int32, int32, int32, int32, int32, int32, int32, int32, string, int32, int32, int32, string, int32, string, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, string, int32, int32, *int32, *string, *int32, *time.Time, *string, *string, *string, *string, *gh.Enum44, *gh.Enum45, *gh.Enum46, *string, *string, *int32, *int32, *string, *string, *string, *bool, *bool, *string, *string, *gh.Enum118, *string, *string, *string, *string, *string, *gh.Enum177, *time.Time, *bool, *int32, *string, *string, *string, *gh.Enum293, *gh.Enum295, exported.Pipeline)
	*/
	perPage := int32(0)
	before := "before"
	pl := runtime.NewPipeline("github.com/nickfloyd/source-generator/generated/go/gh",
		"v0.0.1", runtime.PipelineOptions{}, nil)
	client := gh.NewGitHubV3RESTAPIClient(
		nil, 0, 0, "", "", "", "", 0, 0, 0, "octokit", "octokit.net", 0, "octokit", 0,
		0, 0, "kfcampbell", "secretName", "kfcampbell", 0, 0, "codespaceName",
		0, "octokit.net", gh.Enum121Docker, "packageName", 0, "teamSlug", 0, 0, 0, 0,
		gh.Enum157DependabotAlerts, gh.Enum158EnableAll, 0, 0, 0, 0, 0, 0, 0, 0, "workflowId",
		0, 0, 0, "commitSHA", 0, "envName", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, "exportID",
		0, 0, &perPage, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &before, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, pl,
	)
	// fmt.Printf("client: %v\n", client)

	// get the license of the octokit.net repo
	opts := &gh.GitHubV3RESTAPIClientLicensesGetForRepoOptions{}
	res, err := client.LicensesGetForRepo(context.Background(), opts)
	if err != nil {
		log.Fatalf("error getting licenses for repo: %v", err)
	}
	fmt.Printf("license: %v \nurl: %v", *res.License.Name, *res.GitURL)
}
