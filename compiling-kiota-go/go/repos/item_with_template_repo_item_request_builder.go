package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithTemplate_repoItemRequestBuilder builds and executes requests for operations under \repos\{template_owner}\{template_repo}
type ItemWithTemplate_repoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemWithTemplate_repoItemRequestBuilderInternal instantiates a new WithTemplate_repoItemRequestBuilder and sets the default values.
func NewItemWithTemplate_repoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithTemplate_repoItemRequestBuilder) {
    m := &ItemWithTemplate_repoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{template_owner}/{template_repo}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemWithTemplate_repoItemRequestBuilder instantiates a new WithTemplate_repoItemRequestBuilder and sets the default values.
func NewItemWithTemplate_repoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithTemplate_repoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithTemplate_repoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Generate the generate property
func (m *ItemWithTemplate_repoItemRequestBuilder) Generate()(*ItemItemGenerateRequestBuilder) {
    return NewItemItemGenerateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
