package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\deliveries\{delivery_id}
type ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attempts the attempts property
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Attempts()(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    m := &ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a delivery for a webhook configured in an organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#get-a-webhook-delivery-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.HookDeliveryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateHookDeliveryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.HookDeliveryable), nil
}
// ToGetRequestInformation returns a delivery for a webhook configured in an organization.
func (m *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
