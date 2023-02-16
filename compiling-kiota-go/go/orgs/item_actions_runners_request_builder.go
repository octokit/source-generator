package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnersRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners
type ItemActionsRunnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRunnersRequestBuilderGetQueryParameters lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
type ItemActionsRunnersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemActionsRunnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsRunnersRequestBuilderGetQueryParameters
}
// NewItemActionsRunnersRequestBuilderInternal instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemActionsRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRequestBuilder) {
    m := &ItemActionsRunnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/runners{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRunnersRequestBuilder instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemActionsRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Downloads the downloads property
func (m *ItemActionsRunnersRequestBuilder) Downloads()(*ItemActionsRunnersDownloadsRequestBuilder) {
    return NewItemActionsRunnersDownloadsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-self-hosted-runners-for-an-organization
func (m *ItemActionsRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemActionsRunnersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersResponseable), nil
}
// RegistrationToken the registrationToken property
func (m *ItemActionsRunnersRequestBuilder) RegistrationToken()(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoveToken the removeToken property
func (m *ItemActionsRunnersRequestBuilder) RemoveToken()(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemActionsRunnersRemoveTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemActionsRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
