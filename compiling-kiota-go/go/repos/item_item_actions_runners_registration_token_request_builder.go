package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRunnersRegistrationTokenRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\registration-token
type ItemItemActionsRunnersRegistrationTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    m := &ItemItemActionsRunnersRegistrationTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runners/registration-token";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunnersRegistrationTokenRequestBuilder instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script. The token expires after one hour. You must authenticateusing an access token with the `repo` scope to use this endpoint.#### Example using registration token Configure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.```./config.sh --url https://github.com/octo-org/octo-repo-artifacts --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#create-a-registration-token-for-a-repository
func (m *ItemItemActionsRunnersRegistrationTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.AuthenticationTokenable, error) {
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
// ToPostRequestInformation returns a token that you can pass to the `config` script. The token expires after one hour. You must authenticateusing an access token with the `repo` scope to use this endpoint.#### Example using registration token Configure your self-hosted runner, replacing `TOKEN` with the registration token provided by this endpoint.```./config.sh --url https://github.com/octo-org/octo-repo-artifacts --token TOKEN```
func (m *ItemItemActionsRunnersRegistrationTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
