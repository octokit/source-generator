package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemGitMatchingRefsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\git\matching-refs
type ItemItemGitMatchingRefsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemGitMatchingRefsRequestBuilderInternal instantiates a new MatchingRefsRequestBuilder and sets the default values.
func NewItemItemGitMatchingRefsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitMatchingRefsRequestBuilder) {
    m := &ItemItemGitMatchingRefsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/git/matching-refs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemGitMatchingRefsRequestBuilder instantiates a new MatchingRefsRequestBuilder and sets the default values.
func NewItemItemGitMatchingRefsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitMatchingRefsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitMatchingRefsRequestBuilderInternal(urlParams, requestAdapter)
}
