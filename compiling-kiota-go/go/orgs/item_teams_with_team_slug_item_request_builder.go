package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemTeamsWithTeam_slugItemRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}
type ItemTeamsWithTeam_slugItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamsWithTeam_slugItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsWithTeam_slugItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamsWithTeam_slugItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsWithTeam_slugItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamsWithTeam_slugItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsWithTeam_slugItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamsWithTeam_slugItemRequestBuilderInternal instantiates a new WithTeam_slugItemRequestBuilder and sets the default values.
func NewItemTeamsWithTeam_slugItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsWithTeam_slugItemRequestBuilder) {
    m := &ItemTeamsWithTeam_slugItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/teams/{team_slug}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamsWithTeam_slugItemRequestBuilder instantiates a new WithTeam_slugItemRequestBuilder and sets the default values.
func NewItemTeamsWithTeam_slugItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsWithTeam_slugItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsWithTeam_slugItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete to delete a team, the authenticated user must be an organization owner or team maintainer.If you are an organization owner, deleting a parent team will delete all of its child teams as well.**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#delete-a-team
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Discussions the discussions property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Discussions()(*ItemTeamsItemDiscussionsRequestBuilder) {
    return NewItemTeamsItemDiscussionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DiscussionsById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item.discussions.item collection
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) DiscussionsById(id string)(*ItemTeamsItemDiscussionsWithDiscussion_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["discussion_number"] = id
    }
    return NewItemTeamsItemDiscussionsWithDiscussion_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get gets a team using the team's `slug`. GitHub generates the `slug` from the team `name`.**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#get-a-team-by-name
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable), nil
}
// Invitations the invitations property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Invitations()(*ItemTeamsItemInvitationsRequestBuilder) {
    return NewItemTeamsItemInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members the members property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Members()(*ItemTeamsItemMembersRequestBuilder) {
    return NewItemTeamsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Memberships the memberships property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Memberships()(*ItemTeamsItemMembershipsRequestBuilder) {
    return NewItemTeamsItemMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembershipsById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item.memberships.item collection
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) MembershipsById(id string)(*ItemTeamsItemMembershipsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemTeamsItemMembershipsWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch to edit a team, the authenticated user must either be an organization owner or a team maintainer.**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams#update-a-team
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Patch(ctx context.Context, body ItemTeamsItemWithTeam_slugPatchRequestBodyable, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable), nil
}
// Projects the projects property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Projects()(*ItemTeamsItemProjectsRequestBuilder) {
    return NewItemTeamsItemProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ProjectsById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item.projects.item collection
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) ProjectsById(id string)(*ItemTeamsItemProjectsWithProject_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["project_id"] = id
    }
    return NewItemTeamsItemProjectsWithProject_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Repos the repos property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Repos()(*ItemTeamsItemReposRequestBuilder) {
    return NewItemTeamsItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReposById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item.repos.item collection
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) ReposById(id string)(*ItemTeamsItemReposWithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["owner"] = id
    }
    return NewItemTeamsItemReposWithOwnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Teams the teams property
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) Teams()(*ItemTeamsItemTeamsRequestBuilder) {
    return NewItemTeamsItemTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation to delete a team, the authenticated user must be an organization owner or team maintainer.If you are an organization owner, deleting a parent team will delete all of its child teams as well.**Note:** You can also specify a team by `org_id` and `team_id` using the route `DELETE /organizations/{org_id}/team/{team_id}`.
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation gets a team using the team's `slug`. GitHub generates the `slug` from the team `name`.**Note:** You can also specify a team by `org_id` and `team_id` using the route `GET /organizations/{org_id}/team/{team_id}`.
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation to edit a team, the authenticated user must either be an organization owner or a team maintainer.**Note:** You can also specify a team by `org_id` and `team_id` using the route `PATCH /organizations/{org_id}/team/{team_id}`.
func (m *ItemTeamsWithTeam_slugItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemTeamsItemWithTeam_slugPatchRequestBodyable, requestConfiguration *ItemTeamsWithTeam_slugItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
