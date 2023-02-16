package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphCompareRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependency-graph\compare
type ItemItemDependencyGraphCompareRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemDependencyGraphCompareRequestBuilderInternal instantiates a new CompareRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareRequestBuilder) {
    m := &ItemItemDependencyGraphCompareRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/dependency-graph/compare";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDependencyGraphCompareRequestBuilder instantiates a new CompareRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphCompareRequestBuilderInternal(urlParams, requestAdapter)
}
