package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithTemplate_ownerItemRequestBuilder builds and executes requests for operations under \repos\{template_owner}
type WithTemplate_ownerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithTemplate_ownerItemRequestBuilderInternal instantiates a new WithTemplate_ownerItemRequestBuilder and sets the default values.
func NewWithTemplate_ownerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTemplate_ownerItemRequestBuilder) {
    m := &WithTemplate_ownerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{template_owner}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithTemplate_ownerItemRequestBuilder instantiates a new WithTemplate_ownerItemRequestBuilder and sets the default values.
func NewWithTemplate_ownerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTemplate_ownerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithTemplate_ownerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ItemById gets an item from the github.com/octokit/kiota/.repos.item.item collection
func (m *WithTemplate_ownerItemRequestBuilder) ItemById(id string)(*ItemWithTemplate_repoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["template_repo"] = id
    }
    return NewItemWithTemplate_repoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
