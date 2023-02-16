package gitignore

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GitignoreRequestBuilder builds and executes requests for operations under \gitignore
type GitignoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewGitignoreRequestBuilderInternal instantiates a new GitignoreRequestBuilder and sets the default values.
func NewGitignoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GitignoreRequestBuilder) {
    m := &GitignoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/gitignore";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGitignoreRequestBuilder instantiates a new GitignoreRequestBuilder and sets the default values.
func NewGitignoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GitignoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGitignoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Templates the templates property
func (m *GitignoreRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemplatesById gets an item from the github.com/octokit/kiota/.gitignore.templates.item collection
func (m *GitignoreRequestBuilder) TemplatesById(id string)(*TemplatesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["name"] = id
    }
    return NewTemplatesWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
