package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\comments\{comment_id}\reactions\{reaction_id}
type ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderInternal instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    m := &ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/pulls/comments/:comment_id/reactions/:reaction_id.`Delete a reaction to a [pull request review comment](https://docs.github.com/rest/reference/pulls#review-comments).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/reactions#delete-a-pull-request-comment-reaction
func (m *ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation **Note:** You can also specify a repository by `repository_id` using the route `DELETE /repositories/:repository_id/pulls/comments/:comment_id/reactions/:reaction_id.`Delete a reaction to a [pull request review comment](https://docs.github.com/rest/reference/pulls#review-comments).
func (m *ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
