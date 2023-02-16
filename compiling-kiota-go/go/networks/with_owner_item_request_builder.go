package networks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithOwnerItemRequestBuilder builds and executes requests for operations under \networks\{owner}
type WithOwnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithOwnerItemRequestBuilderInternal instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOwnerItemRequestBuilder) {
    m := &WithOwnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/networks/{owner}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithOwnerItemRequestBuilder instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
