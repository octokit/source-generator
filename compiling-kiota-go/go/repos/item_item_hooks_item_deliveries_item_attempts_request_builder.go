package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\hooks\{hook_id}\deliveries\{delivery_id}\attempts
type ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    m := &ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}/attempts";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilder instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post redeliver a webhook delivery for a webhook configured in a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/webhooks/repo-deliveries#redeliver-a-delivery-for-a-repository-webhook
func (m *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(ItemItemHooksItemDeliveriesItemAttemptsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemHooksItemDeliveriesItemAttemptsResponseable), nil
}
// ToPostRequestInformation redeliver a webhook delivery for a webhook configured in a repository.
func (m *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
