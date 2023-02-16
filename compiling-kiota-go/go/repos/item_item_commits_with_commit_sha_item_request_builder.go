package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsWithCommit_shaItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\commits\{commit_sha}
type ItemItemCommitsWithCommit_shaItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BranchesWhereHead the branchesWhereHead property
func (m *ItemItemCommitsWithCommit_shaItemRequestBuilder) BranchesWhereHead()(*ItemItemCommitsItemBranchesWhereHeadRequestBuilder) {
    return NewItemItemCommitsItemBranchesWhereHeadRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Comments the comments property
func (m *ItemItemCommitsWithCommit_shaItemRequestBuilder) Comments()(*ItemItemCommitsItemCommentsRequestBuilder) {
    return NewItemItemCommitsItemCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemCommitsWithCommit_shaItemRequestBuilderInternal instantiates a new WithCommit_shaItemRequestBuilder and sets the default values.
func NewItemItemCommitsWithCommit_shaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsWithCommit_shaItemRequestBuilder) {
    m := &ItemItemCommitsWithCommit_shaItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/commits/{commit_sha}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCommitsWithCommit_shaItemRequestBuilder instantiates a new WithCommit_shaItemRequestBuilder and sets the default values.
func NewItemItemCommitsWithCommit_shaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsWithCommit_shaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsWithCommit_shaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Pulls the pulls property
func (m *ItemItemCommitsWithCommit_shaItemRequestBuilder) Pulls()(*ItemItemCommitsItemPullsRequestBuilder) {
    return NewItemItemCommitsItemPullsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
