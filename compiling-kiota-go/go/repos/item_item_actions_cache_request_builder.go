package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsCacheRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\cache
type ItemItemActionsCacheRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemActionsCacheRequestBuilderInternal instantiates a new CacheRequestBuilder and sets the default values.
func NewItemItemActionsCacheRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCacheRequestBuilder) {
    m := &ItemItemActionsCacheRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/cache";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsCacheRequestBuilder instantiates a new CacheRequestBuilder and sets the default values.
func NewItemItemActionsCacheRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCacheRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsCacheRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
func (m *ItemItemActionsCacheRequestBuilder) Usage()(*ItemItemActionsCacheUsageRequestBuilder) {
    return NewItemItemActionsCacheUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
