package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}\codespaces\{codespace_name}
type ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    m := &ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/members/{username}/codespaces/{codespace_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a user's codespace.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(ItemMembersItemCodespacesItemWithCodespace_nameResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "500": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMembersItemCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMembersItemCodespacesItemWithCodespace_nameResponseable), nil
}
// Stop the stop property
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) Stop()(*ItemMembersItemCodespacesItemStopRequestBuilder) {
    return NewItemMembersItemCodespacesItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation deletes a user's codespace.You must authenticate using an access token with the `admin:org` scope to use this endpoint.
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
