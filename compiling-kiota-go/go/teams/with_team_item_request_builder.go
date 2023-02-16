package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// WithTeam_ItemRequestBuilder builds and executes requests for operations under \teams\{team_id}
type WithTeam_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WithTeam_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithTeam_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithTeam_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithTeam_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithTeam_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithTeam_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWithTeam_ItemRequestBuilderInternal instantiates a new WithTeam_ItemRequestBuilder and sets the default values.
func NewWithTeam_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTeam_ItemRequestBuilder) {
    m := &WithTeam_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithTeam_ItemRequestBuilder instantiates a new WithTeam_ItemRequestBuilder and sets the default values.
func NewWithTeam_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTeam_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithTeam_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a team](https://docs.github.com/rest/reference/teams#delete-a-team) endpoint.To delete a team, the authenticated user must be an organization owner or team maintainer.If you are an organization owner, deleting a parent team will delete all of its child teams as well.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams/#delete-a-team-legacy
func (m *WithTeam_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithTeam_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Discussions the discussions property
func (m *WithTeam_ItemRequestBuilder) Discussions()(*ItemDiscussionsRequestBuilder) {
    return NewItemDiscussionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DiscussionsById gets an item from the github.com/octokit/kiota/.teams.item.discussions.item collection
func (m *WithTeam_ItemRequestBuilder) DiscussionsById(id string)(*ItemDiscussionsWithDiscussion_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["discussion_number"] = id
    }
    return NewItemDiscussionsWithDiscussion_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the [Get a team by name](https://docs.github.com/rest/reference/teams#get-a-team-by-name) endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams/#get-a-team-legacy
func (m *WithTeam_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithTeam_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable, error) {
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
func (m *WithTeam_ItemRequestBuilder) Invitations()(*ItemInvitationsRequestBuilder) {
    return NewItemInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members the members property
func (m *WithTeam_ItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById gets an item from the github.com/octokit/kiota/.teams.item.members.item collection
func (m *WithTeam_ItemRequestBuilder) MembersById(id string)(*ItemMembersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemMembersWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Memberships the memberships property
func (m *WithTeam_ItemRequestBuilder) Memberships()(*ItemMembershipsRequestBuilder) {
    return NewItemMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembershipsById gets an item from the github.com/octokit/kiota/.teams.item.memberships.item collection
func (m *WithTeam_ItemRequestBuilder) MembershipsById(id string)(*ItemMembershipsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemMembershipsWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a team](https://docs.github.com/rest/reference/teams#update-a-team) endpoint.To edit a team, the authenticated user must either be an organization owner or a team maintainer.**Note:** With nested teams, the `privacy` for parent teams cannot be `secret`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/teams/#update-a-team-legacy
func (m *WithTeam_ItemRequestBuilder) Patch(ctx context.Context, body ItemWithTeam_PatchRequestBodyable, requestConfiguration *WithTeam_ItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamFullable, error) {
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
func (m *WithTeam_ItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ProjectsById gets an item from the github.com/octokit/kiota/.teams.item.projects.item collection
func (m *WithTeam_ItemRequestBuilder) ProjectsById(id string)(*ItemProjectsWithProject_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["project_id"] = id
    }
    return NewItemProjectsWithProject_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Repos the repos property
func (m *WithTeam_ItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReposById gets an item from the github.com/octokit/kiota/.teams.item.repos.item collection
func (m *WithTeam_ItemRequestBuilder) ReposById(id string)(*ItemReposWithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["owner"] = id
    }
    return NewItemReposWithOwnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Teams the teams property
func (m *WithTeam_ItemRequestBuilder) Teams()(*ItemTeamsRequestBuilder) {
    return NewItemTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Delete a team](https://docs.github.com/rest/reference/teams#delete-a-team) endpoint.To delete a team, the authenticated user must be an organization owner or team maintainer.If you are an organization owner, deleting a parent team will delete all of its child teams as well.
func (m *WithTeam_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithTeam_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the [Get a team by name](https://docs.github.com/rest/reference/teams#get-a-team-by-name) endpoint.
func (m *WithTeam_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithTeam_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Deprecation Notice:** This endpoint route is deprecated and will be removed from the Teams API. We recommend migrating your existing code to use the new [Update a team](https://docs.github.com/rest/reference/teams#update-a-team) endpoint.To edit a team, the authenticated user must either be an organization owner or a team maintainer.**Note:** With nested teams, the `privacy` for parent teams cannot be `secret`.
func (m *WithTeam_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithTeam_PatchRequestBodyable, requestConfiguration *WithTeam_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
