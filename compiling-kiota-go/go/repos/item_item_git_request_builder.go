package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemGitRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\git
type ItemItemGitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Blobs the blobs property
func (m *ItemItemGitRequestBuilder) Blobs()(*ItemItemGitBlobsRequestBuilder) {
    return NewItemItemGitBlobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BlobsById gets an item from the github.com/octokit/kiota/.repos.item.item.git.blobs.item collection
func (m *ItemItemGitRequestBuilder) BlobsById(id string)(*ItemItemGitBlobsWithFile_shaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["file_sha"] = id
    }
    return NewItemItemGitBlobsWithFile_shaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Commits the commits property
func (m *ItemItemGitRequestBuilder) Commits()(*ItemItemGitCommitsRequestBuilder) {
    return NewItemItemGitCommitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CommitsById gets an item from the github.com/octokit/kiota/.repos.item.item.git.commits.item collection
func (m *ItemItemGitRequestBuilder) CommitsById(id string)(*ItemItemGitCommitsWithCommit_shaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commit_sha"] = id
    }
    return NewItemItemGitCommitsWithCommit_shaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemGitRequestBuilderInternal instantiates a new GitRequestBuilder and sets the default values.
func NewItemItemGitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRequestBuilder) {
    m := &ItemItemGitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/git";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemGitRequestBuilder instantiates a new GitRequestBuilder and sets the default values.
func NewItemItemGitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitRequestBuilderInternal(urlParams, requestAdapter)
}
// MatchingRefs the matchingRefs property
func (m *ItemItemGitRequestBuilder) MatchingRefs()(*ItemItemGitMatchingRefsRequestBuilder) {
    return NewItemItemGitMatchingRefsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MatchingRefsById gets an item from the github.com/octokit/kiota/.repos.item.item.git.matchingRefs.item collection
func (m *ItemItemGitRequestBuilder) MatchingRefsById(id string)(*ItemItemGitMatchingRefsWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ref"] = id
    }
    return NewItemItemGitMatchingRefsWithRefItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Ref the ref property
func (m *ItemItemGitRequestBuilder) Ref()(*ItemItemGitRefRequestBuilder) {
    return NewItemItemGitRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RefById gets an item from the github.com/octokit/kiota/.repos.item.item.git.ref.item collection
func (m *ItemItemGitRequestBuilder) RefById(id string)(*ItemItemGitRefWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ref"] = id
    }
    return NewItemItemGitRefWithRefItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Refs the refs property
func (m *ItemItemGitRequestBuilder) Refs()(*ItemItemGitRefsRequestBuilder) {
    return NewItemItemGitRefsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RefsById gets an item from the github.com/octokit/kiota/.repos.item.item.git.refs.item collection
func (m *ItemItemGitRequestBuilder) RefsById(id string)(*ItemItemGitRefsWithRefItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ref"] = id
    }
    return NewItemItemGitRefsWithRefItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tags the tags property
func (m *ItemItemGitRequestBuilder) Tags()(*ItemItemGitTagsRequestBuilder) {
    return NewItemItemGitTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById gets an item from the github.com/octokit/kiota/.repos.item.item.git.tags.item collection
func (m *ItemItemGitRequestBuilder) TagsById(id string)(*ItemItemGitTagsWithTag_shaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tag_sha"] = id
    }
    return NewItemItemGitTagsWithTag_shaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Trees the trees property
func (m *ItemItemGitRequestBuilder) Trees()(*ItemItemGitTreesRequestBuilder) {
    return NewItemItemGitTreesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TreesById gets an item from the github.com/octokit/kiota/.repos.item.item.git.trees.item collection
func (m *ItemItemGitRequestBuilder) TreesById(id string)(*ItemItemGitTreesWithTree_shaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tree_sha"] = id
    }
    return NewItemItemGitTreesWithTree_shaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
