package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRunnersRemoveTokenRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\remove-token
type ItemActionsRunnersRemoveTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRunnersRemoveTokenRequestBuilderInternal instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    m := &ItemActionsRunnersRemoveTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/runners/remove-token";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRunnersRemoveTokenRequestBuilder instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRemoveTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script to remove a self-hosted runner from an organization. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.#### Example using remove tokenTo remove your self-hosted runner from an organization, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#create-a-remove-token-for-an-organization
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.AuthenticationTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateAuthenticationTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.AuthenticationTokenable), nil
}
// ToPostRequestInformation returns a token that you can pass to the `config` script to remove a self-hosted runner from an organization. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.#### Example using remove tokenTo remove your self-hosted runner from an organization, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
