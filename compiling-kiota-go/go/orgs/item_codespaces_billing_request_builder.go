package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemCodespacesBillingRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces\billing
type ItemCodespacesBillingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCodespacesBillingRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesBillingRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCodespacesBillingRequestBuilderInternal instantiates a new BillingRequestBuilder and sets the default values.
func NewItemCodespacesBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesBillingRequestBuilder) {
    m := &ItemCodespacesBillingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/codespaces/billing";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCodespacesBillingRequestBuilder instantiates a new BillingRequestBuilder and sets the default values.
func NewItemCodespacesBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Put sets which users can access codespaces in an organization. This is synonymous with granting or revoking codespaces billing permissions for users according to the visibility.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#set-codespaces-billing
func (m *ItemCodespacesBillingRequestBuilder) Put(ctx context.Context, body ItemCodespacesBillingPutRequestBodyable, requestConfiguration *ItemCodespacesBillingRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
        "500": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Selected_users the selected_users property
func (m *ItemCodespacesBillingRequestBuilder) Selected_users()(*ItemCodespacesBillingSelected_usersRequestBuilder) {
    return NewItemCodespacesBillingSelected_usersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToPutRequestInformation sets which users can access codespaces in an organization. This is synonymous with granting or revoking codespaces billing permissions for users according to the visibility.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemCodespacesBillingRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemCodespacesBillingPutRequestBodyable, requestConfiguration *ItemCodespacesBillingRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
