package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCodespacesSecretsRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces\secrets
type ItemCodespacesSecretsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCodespacesSecretsRequestBuilderGetQueryParameters lists all Codespaces secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
type ItemCodespacesSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemCodespacesSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCodespacesSecretsRequestBuilderGetQueryParameters
}
// NewItemCodespacesSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemCodespacesSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesSecretsRequestBuilder) {
    m := &ItemCodespacesSecretsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/codespaces/secrets{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCodespacesSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewItemCodespacesSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all Codespaces secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#list-organization-secrets
func (m *ItemCodespacesSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(ItemCodespacesSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemCodespacesSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodespacesSecretsResponseable), nil
}
// PublicKey the publicKey property
func (m *ItemCodespacesSecretsRequestBuilder) PublicKey()(*ItemCodespacesSecretsPublicKeyRequestBuilder) {
    return NewItemCodespacesSecretsPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation lists all Codespaces secrets available at the organization-level without revealing their encrypted values.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemCodespacesSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCodespacesSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
