package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependency-graph\compare\{basehead}
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters struct {
    // The full path, relative to the repository root, of the dependency manifest file.
    Name *string
}
// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    m := &ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/dependency-graph/compare/{basehead}{?name*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/dependency-graph#get-a-diff-of-the-dependencies-between-commits
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependencyGraphDiffable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDependencyGraphDiffFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependencyGraphDiffable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependencyGraphDiffable)
    }
    return val, nil
}
// ToGetRequestInformation gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
