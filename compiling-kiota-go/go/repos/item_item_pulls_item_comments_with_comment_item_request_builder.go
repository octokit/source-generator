package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPullsItemCommentsWithComment_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pulls\{pull_number}\comments\{comment_id}
type ItemItemPullsItemCommentsWithComment_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    m := &ItemItemPullsItemCommentsWithComment_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/pulls/{pull_number}/comments/{comment_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemPullsItemCommentsWithComment_ItemRequestBuilder instantiates a new WithComment_ItemRequestBuilder and sets the default values.
func NewItemItemPullsItemCommentsWithComment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemCommentsWithComment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Replies the replies property
func (m *ItemItemPullsItemCommentsWithComment_ItemRequestBuilder) Replies()(*ItemItemPullsItemCommentsItemRepliesRequestBuilder) {
    return NewItemItemPullsItemCommentsItemRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
