package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTrafficPopularRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic\popular
type ItemItemTrafficPopularRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemTrafficPopularRequestBuilderInternal instantiates a new PopularRequestBuilder and sets the default values.
func NewItemItemTrafficPopularRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularRequestBuilder) {
    m := &ItemItemTrafficPopularRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/traffic/popular";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemTrafficPopularRequestBuilder instantiates a new PopularRequestBuilder and sets the default values.
func NewItemItemTrafficPopularRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficPopularRequestBuilderInternal(urlParams, requestAdapter)
}
// Paths the paths property
func (m *ItemItemTrafficPopularRequestBuilder) Paths()(*ItemItemTrafficPopularPathsRequestBuilder) {
    return NewItemItemTrafficPopularPathsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Referrers the referrers property
func (m *ItemItemTrafficPopularRequestBuilder) Referrers()(*ItemItemTrafficPopularReferrersRequestBuilder) {
    return NewItemItemTrafficPopularReferrersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
