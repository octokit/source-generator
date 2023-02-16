package networks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithRepoItemRequestBuilder builds and executes requests for operations under \networks\{owner}\{repo}
type ItemWithRepoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemWithRepoItemRequestBuilderInternal instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    m := &ItemWithRepoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/networks/{owner}/{repo}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemWithRepoItemRequestBuilder instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithRepoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithRepoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
func (m *ItemWithRepoItemRequestBuilder) Events()(*ItemItemEventsRequestBuilder) {
    return NewItemItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
