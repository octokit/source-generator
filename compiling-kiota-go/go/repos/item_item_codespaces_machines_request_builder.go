package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCodespacesMachinesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces\machines
type ItemItemCodespacesMachinesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCodespacesMachinesRequestBuilderGetQueryParameters list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
type ItemItemCodespacesMachinesRequestBuilderGetQueryParameters struct {
    // IP for location auto-detection when proxying a request
    Client_ip *string
    // The location to check for available machines. Assigned by IP if not provided.
    Location *string
}
// ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesMachinesRequestBuilderGetQueryParameters
}
// NewItemItemCodespacesMachinesRequestBuilderInternal instantiates a new MachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    m := &ItemItemCodespacesMachinesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/codespaces/machines{?location*,client_ip*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodespacesMachinesRequestBuilder instantiates a new MachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesMachinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#list-available-machine-types-for-a-repository
func (m *ItemItemCodespacesMachinesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration)(ItemItemCodespacesMachinesResponseable, error) {
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
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesMachinesResponseable), nil
}
// ToGetRequestInformation list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
func (m *ItemItemCodespacesMachinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
