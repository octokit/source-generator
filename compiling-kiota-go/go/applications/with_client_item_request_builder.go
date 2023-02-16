package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithClient_ItemRequestBuilder builds and executes requests for operations under \applications\{client_id}
type WithClient_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithClient_ItemRequestBuilderInternal instantiates a new WithClient_ItemRequestBuilder and sets the default values.
func NewWithClient_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClient_ItemRequestBuilder) {
    m := &WithClient_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{client_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithClient_ItemRequestBuilder instantiates a new WithClient_ItemRequestBuilder and sets the default values.
func NewWithClient_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithClient_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithClient_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Grant the grant property
func (m *WithClient_ItemRequestBuilder) Grant()(*ItemGrantRequestBuilder) {
    return NewItemGrantRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Token the token property
func (m *WithClient_ItemRequestBuilder) Token()(*ItemTokenRequestBuilder) {
    return NewItemTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
