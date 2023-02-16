package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsCacheUsageByRepositoryRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\cache\usage-by-repository
type ItemActionsCacheUsageByRepositoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
type ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsCacheUsageByRepositoryRequestBuilderGetQueryParameters
}
// NewItemActionsCacheUsageByRepositoryRequestBuilderInternal instantiates a new UsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    m := &ItemActionsCacheUsageByRepositoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/cache/usage-by-repository{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsCacheUsageByRepositoryRequestBuilder instantiates a new UsageByRepositoryRequestBuilder and sets the default values.
func NewItemActionsCacheUsageByRepositoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsCacheUsageByRepositoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsCacheUsageByRepositoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-repositories-with-github-actions-cache-usage-for-an-organization
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration)(ItemActionsCacheUsageByRepositoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsCacheUsageByRepositoryResponseable), nil
}
// ToGetRequestInformation lists repositories and their GitHub Actions cache usage for an organization.The data fetched using this API is refreshed approximately every 5 minutes, so values returned from this endpoint may take at least 5 minutes to get updated.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `organization_admistration:read` permission to use this endpoint.
func (m *ItemActionsCacheUsageByRepositoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsCacheUsageByRepositoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
