package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnterprisesRequestBuilder builds and executes requests for operations under \enterprises
type EnterprisesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEnterprisesRequestBuilderInternal instantiates a new EnterprisesRequestBuilder and sets the default values.
func NewEnterprisesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisesRequestBuilder) {
    m := &EnterprisesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/enterprises";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEnterprisesRequestBuilder instantiates a new EnterprisesRequestBuilder and sets the default values.
func NewEnterprisesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterprisesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterprisesRequestBuilderInternal(urlParams, requestAdapter)
}
