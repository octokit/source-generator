package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemCodespacesRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces
type ItemCodespacesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCodespacesRequestBuilderGetQueryParameters lists the codespaces associated to a specified organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
type ItemCodespacesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemCodespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCodespacesRequestBuilderGetQueryParameters
}
// Billing the billing property
func (m *ItemCodespacesRequestBuilder) Billing()(*ItemCodespacesBillingRequestBuilder) {
    return NewItemCodespacesBillingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCodespacesRequestBuilderInternal instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemCodespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesRequestBuilder) {
    m := &ItemCodespacesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/codespaces{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCodespacesRequestBuilder instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemCodespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the codespaces associated to a specified organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#list-in-organization
func (m *ItemCodespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCodespacesRequestBuilderGetRequestConfiguration)(ItemCodespacesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "500": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemCodespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCodespacesResponseable), nil
}
// Secrets the secrets property
func (m *ItemCodespacesRequestBuilder) Secrets()(*ItemCodespacesSecretsRequestBuilder) {
    return NewItemCodespacesSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.orgs.item.codespaces.secrets.item collection
func (m *ItemCodespacesRequestBuilder) SecretsById(id string)(*ItemCodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation lists the codespaces associated to a specified organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemCodespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCodespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
