package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCompareRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\compare
type ItemItemCompareRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemCompareRequestBuilderInternal instantiates a new CompareRequestBuilder and sets the default values.
func NewItemItemCompareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareRequestBuilder) {
    m := &ItemItemCompareRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/compare";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCompareRequestBuilder instantiates a new CompareRequestBuilder and sets the default values.
func NewItemItemCompareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCompareRequestBuilderInternal(urlParams, requestAdapter)
}
