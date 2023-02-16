package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRunnersRegistrationTokenRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\registration-token
type ItemActionsRunnersRegistrationTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRunnersRegistrationTokenRequestBuilderInternal instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    m := &ItemActionsRunnersRegistrationTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/runners/registration-token";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRunnersRegistrationTokenRequestBuilder instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.#### Example using registration tokenConfigure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.```./config.sh --url https://github.com/octo-org --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#create-a-registration-token-for-an-organization
func (m *ItemActionsRunnersRegistrationTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.AuthenticationTokenable, error) {
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
// ToPostRequestInformation returns a token that you can pass to the `config` script. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.#### Example using registration tokenConfigure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.```./config.sh --url https://github.com/octo-org --token TOKEN```
func (m *ItemActionsRunnersRegistrationTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
