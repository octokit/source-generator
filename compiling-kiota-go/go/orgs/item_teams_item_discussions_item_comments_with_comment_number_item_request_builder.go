package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\discussions\{discussion_number}\comments\{comment_number}
type ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal instantiates a new WithComment_numberItemRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    m := &ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder instantiates a new WithComment_numberItemRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#delete-a-discussion-comment
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a specific comment on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#get-a-discussion-comment
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamDiscussionCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamDiscussionCommentable), nil
}
// Patch edits the body text of a discussion comment. OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#update-a-discussion-comment
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Patch(ctx context.Context, body ItemTeamsItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyable, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamDiscussionCommentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamDiscussionCommentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamDiscussionCommentable), nil
}
// Reactions the reactions property
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) Reactions()(*ItemTeamsItemDiscussionsItemCommentsItemReactionsRequestBuilder) {
    return NewItemTeamsItemDiscussionsItemCommentsItemReactionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReactionsById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item.discussions.item.comments.item.reactions.item collection
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ReactionsById(id string)(*ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reaction_id"] = id
    }
    return NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation deletes a comment on a team discussion. OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get a specific comment on a team discussion. OAuth access tokens require the `read:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation edits the body text of a discussion comment. OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}/discussions/{discussion_number}/comments/{comment_number}`.
func (m *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemTeamsItemDiscussionsItemCommentsItemWithComment_numberPatchRequestBodyable, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsWithComment_numberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
