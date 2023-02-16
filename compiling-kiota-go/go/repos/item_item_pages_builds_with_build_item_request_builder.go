package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemPagesBuildsWithBuild_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pages\builds\{build_id}
type ItemItemPagesBuildsWithBuild_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal instantiates a new WithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    m := &ItemItemPagesBuildsWithBuild_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/pages/builds/{build_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilder instantiates a new WithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get GitHub Pages build
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pages#get-github-pages-build
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PageBuildable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreatePageBuildFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PageBuildable), nil
}
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
