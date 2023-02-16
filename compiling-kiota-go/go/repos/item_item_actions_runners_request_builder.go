package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners
type ItemItemActionsRunnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunnersRequestBuilderGetQueryParameters lists all self-hosted runners configured in a repository. You must authenticate using an access token with the `repo` scope to use this endpoint.
type ItemItemActionsRunnersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemActionsRunnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunnersRequestBuilderGetQueryParameters
}
// NewItemItemActionsRunnersRequestBuilderInternal instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    m := &ItemItemActionsRunnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runners{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunnersRequestBuilder instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Downloads the downloads property
func (m *ItemItemActionsRunnersRequestBuilder) Downloads()(*ItemItemActionsRunnersDownloadsRequestBuilder) {
    return NewItemItemActionsRunnersDownloadsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get lists all self-hosted runners configured in a repository. You must authenticate using an access token with the `repo` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-self-hosted-runners-for-a-repository
func (m *ItemItemActionsRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemItemActionsRunnersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersResponseable), nil
}
// RegistrationToken the registrationToken property
func (m *ItemItemActionsRunnersRequestBuilder) RegistrationToken()(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoveToken the removeToken property
func (m *ItemItemActionsRunnersRequestBuilder) RemoveToken()(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation lists all self-hosted runners configured in a repository. You must authenticate using an access token with the `repo` scope to use this endpoint.
func (m *ItemItemActionsRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
