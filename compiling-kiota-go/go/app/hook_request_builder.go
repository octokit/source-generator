package app

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HookRequestBuilder builds and executes requests for operations under \app\hook
type HookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Config the config property
func (m *HookRequestBuilder) Config()(*HookConfigRequestBuilder) {
    return NewHookConfigRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewHookRequestBuilderInternal instantiates a new HookRequestBuilder and sets the default values.
func NewHookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookRequestBuilder) {
    m := &HookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/hook";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewHookRequestBuilder instantiates a new HookRequestBuilder and sets the default values.
func NewHookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHookRequestBuilderInternal(urlParams, requestAdapter)
}
// Deliveries the deliveries property
func (m *HookRequestBuilder) Deliveries()(*HookDeliveriesRequestBuilder) {
    return NewHookDeliveriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeliveriesById gets an item from the github.com/octokit/kiota/.app.hook.deliveries.item collection
func (m *HookRequestBuilder) DeliveriesById(id string)(*HookDeliveriesWithDelivery_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delivery_id"] = id
    }
    return NewHookDeliveriesWithDelivery_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
