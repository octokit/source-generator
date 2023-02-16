package search

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SearchRequestBuilder builds and executes requests for operations under \search
type SearchRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Code the code property
func (m *SearchRequestBuilder) Code()(*CodeRequestBuilder) {
    return NewCodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Commits the commits property
func (m *SearchRequestBuilder) Commits()(*CommitsRequestBuilder) {
    return NewCommitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewSearchRequestBuilderInternal instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    m := &SearchRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSearchRequestBuilder instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Issues the issues property
func (m *SearchRequestBuilder) Issues()(*IssuesRequestBuilder) {
    return NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Labels the labels property
func (m *SearchRequestBuilder) Labels()(*LabelsRequestBuilder) {
    return NewLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Repositories the repositories property
func (m *SearchRequestBuilder) Repositories()(*RepositoriesRequestBuilder) {
    return NewRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Topics the topics property
func (m *SearchRequestBuilder) Topics()(*TopicsRequestBuilder) {
    return NewTopicsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Users the users property
func (m *SearchRequestBuilder) Users()(*UsersRequestBuilder) {
    return NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
