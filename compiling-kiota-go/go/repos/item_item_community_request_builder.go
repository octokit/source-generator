package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommunityRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\community
type ItemItemCommunityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemCommunityRequestBuilderInternal instantiates a new CommunityRequestBuilder and sets the default values.
func NewItemItemCommunityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityRequestBuilder) {
    m := &ItemItemCommunityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/community";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCommunityRequestBuilder instantiates a new CommunityRequestBuilder and sets the default values.
func NewItemItemCommunityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommunityRequestBuilderInternal(urlParams, requestAdapter)
}
// Profile the profile property
func (m *ItemItemCommunityRequestBuilder) Profile()(*ItemItemCommunityProfileRequestBuilder) {
    return NewItemItemCommunityProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
