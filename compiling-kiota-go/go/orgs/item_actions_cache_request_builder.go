package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsCacheRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\cache
type ItemActionsCacheRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemActionsCacheRequestBuilderInternal instantiates a new CacheRequestBuilder and sets the default values.
func NewItemActionsCacheRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheRequestBuilder) {
    m := &ItemActionsCacheRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/cache";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsCacheRequestBuilder instantiates a new CacheRequestBuilder and sets the default values.
func NewItemActionsCacheRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheRequestBuilderInternal(urlParams, requestAdapter)
}
// Usage the usage property
func (m *ItemActionsCacheRequestBuilder) Usage()(*ItemActionsCacheUsageRequestBuilder) {
    return NewItemActionsCacheUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UsageByRepository the usageByRepository property
func (m *ItemActionsCacheRequestBuilder) UsageByRepository()(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    return NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
