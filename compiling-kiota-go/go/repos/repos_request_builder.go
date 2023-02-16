package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151 "github.com/octokit/kiota/orgs"
)

// ReposRequestBuilder builds and executes requests for operations under \repos
type ReposRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposRequestBuilder) {
    m := &ReposRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReposRequestBuilderInternal(urlParams, requestAdapter)
}
// ReposByOrg gets an item from the github.com/octokit/kiota/.repos.item collection
func (m *ReposRequestBuilder) ReposByOrg(id *string)(*if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.WithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != nil {
        urlTplParams["org"] = *id
    }
    return if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.NewWithOrgItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ReposByOwner gets an item from the github.com/octokit/kiota/.repos.item collection
func (m *ReposRequestBuilder) ReposByOwner(id string)(*if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["owner"] = id
    }
    return if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ReposByTemplate_owner gets an item from the github.com/octokit/kiota/.repos.item collection
func (m *ReposRequestBuilder) ReposByTemplate_owner(id string)(*WithTemplate_ownerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["template_owner"] = id
    }
    return NewWithTemplate_ownerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
