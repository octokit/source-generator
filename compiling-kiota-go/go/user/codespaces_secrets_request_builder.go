package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesSecretsRequestBuilder builds and executes requests for operations under \user\codespaces\secrets
type CodespacesSecretsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CodespacesSecretsRequestBuilderGetQueryParameters lists all secrets available for a user's Codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
type CodespacesSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// CodespacesSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CodespacesSecretsRequestBuilderGetQueryParameters
}
// NewCodespacesSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewCodespacesSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsRequestBuilder) {
    m := &CodespacesSecretsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/codespaces/secrets{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCodespacesSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewCodespacesSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all secrets available for a user's Codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#list-secrets-for-the-authenticated-user
func (m *CodespacesSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesSecretsRequestBuilderGetRequestConfiguration)(CodespacesSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateCodespacesSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesSecretsResponseable), nil
}
// PublicKey the publicKey property
func (m *CodespacesSecretsRequestBuilder) PublicKey()(*CodespacesSecretsPublicKeyRequestBuilder) {
    return NewCodespacesSecretsPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation lists all secrets available for a user's Codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
func (m *CodespacesSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
