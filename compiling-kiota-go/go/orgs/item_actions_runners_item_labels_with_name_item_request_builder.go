package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRunnersItemLabelsWithNameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\{runner_id}\labels\{name}
type ItemActionsRunnersItemLabelsWithNameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    m := &ItemActionsRunnersItemLabelsWithNameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/runners/{runner_id}/labels/{name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRunnersItemLabelsWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemActionsRunnersItemLabelsWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a custom label from a self-hosted runner configuredin an organization. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#remove-a-custom-label-from-a-self-hosted-runner-for-an-organization
func (m *ItemActionsRunnersItemLabelsWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration)(ItemActionsRunnersItemLabelsItemWithNameResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersItemLabelsItemWithNameResponseable), nil
}
// ToDeleteRequestInformation remove a custom label from a self-hosted runner configuredin an organization. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemActionsRunnersItemLabelsWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
