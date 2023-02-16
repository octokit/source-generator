package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTrafficRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic
type ItemItemTrafficRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Clones the clones property
func (m *ItemItemTrafficRequestBuilder) Clones()(*ItemItemTrafficClonesRequestBuilder) {
    return NewItemItemTrafficClonesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemTrafficRequestBuilderInternal instantiates a new TrafficRequestBuilder and sets the default values.
func NewItemItemTrafficRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficRequestBuilder) {
    m := &ItemItemTrafficRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/traffic";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemTrafficRequestBuilder instantiates a new TrafficRequestBuilder and sets the default values.
func NewItemItemTrafficRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficRequestBuilderInternal(urlParams, requestAdapter)
}
// Popular the popular property
func (m *ItemItemTrafficRequestBuilder) Popular()(*ItemItemTrafficPopularRequestBuilder) {
    return NewItemItemTrafficPopularRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Views the views property
func (m *ItemItemTrafficRequestBuilder) Views()(*ItemItemTrafficViewsRequestBuilder) {
    return NewItemItemTrafficViewsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
