package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\comments\{comment_id}\reactions\{reaction_id}
type ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    m := &ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/comments/:comment_id/reactions/:reaction_id`.Delete a reaction to a [commit comment](https://docs.github.com/rest/reference/repos#comments).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/reactions#delete-a-commit-comment-reaction
func (m *ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation **Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/comments/:comment_id/reactions/:reaction_id`.Delete a reaction to a [commit comment](https://docs.github.com/rest/reference/repos#comments).
func (m *ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
