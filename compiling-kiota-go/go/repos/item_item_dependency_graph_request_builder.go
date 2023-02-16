package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependency-graph
type ItemItemDependencyGraphRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Compare the compare property
func (m *ItemItemDependencyGraphRequestBuilder) Compare()(*ItemItemDependencyGraphCompareRequestBuilder) {
    return NewItemItemDependencyGraphCompareRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CompareById gets an item from the github.com/octokit/kiota/.repos.item.item.dependencyGraph.compare.item collection
func (m *ItemItemDependencyGraphRequestBuilder) CompareById(id string)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["basehead"] = id
    }
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemDependencyGraphRequestBuilderInternal instantiates a new DependencyGraphRequestBuilder and sets the default values.
func NewItemItemDependencyGraphRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphRequestBuilder) {
    m := &ItemItemDependencyGraphRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/dependency-graph";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDependencyGraphRequestBuilder instantiates a new DependencyGraphRequestBuilder and sets the default values.
func NewItemItemDependencyGraphRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphRequestBuilderInternal(urlParams, requestAdapter)
}
// Snapshots the snapshots property
func (m *ItemItemDependencyGraphRequestBuilder) Snapshots()(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    return NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
