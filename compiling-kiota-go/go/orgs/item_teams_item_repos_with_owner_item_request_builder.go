package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamsItemReposWithOwnerItemRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\repos\{owner}
type ItemTeamsItemReposWithOwnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    m := &ItemTeamsItemReposWithOwnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/teams/{team_slug}/repos/{owner}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamsItemReposWithOwnerItemRequestBuilder instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewItemTeamsItemReposWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
