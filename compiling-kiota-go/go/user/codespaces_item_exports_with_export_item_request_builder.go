package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// CodespacesItemExportsWithExport_ItemRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\exports\{export_id}
type CodespacesItemExportsWithExport_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    m := &CodespacesItemExportsWithExport_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/codespaces/{codespace_name}/exports/{export_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilder instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#get-details-about-a-codespace-export
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodespaceExportDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodespaceExportDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodespaceExportDetailsable), nil
}
// ToGetRequestInformation gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
