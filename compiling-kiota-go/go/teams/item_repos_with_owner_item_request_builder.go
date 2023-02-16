package teams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemReposWithOwnerItemRequestBuilder builds and executes requests for operations under \teams\{team_id}\repos\{owner}
type ItemReposWithOwnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemReposWithOwnerItemRequestBuilderInternal instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewItemReposWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposWithOwnerItemRequestBuilder) {
    m := &ItemReposWithOwnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/repos/{owner}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemReposWithOwnerItemRequestBuilder instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewItemReposWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReposWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
