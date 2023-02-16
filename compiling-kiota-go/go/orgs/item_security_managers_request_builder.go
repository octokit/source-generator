package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemSecurityManagersRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers
type ItemSecurityManagersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSecurityManagersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityManagersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityManagersRequestBuilderInternal instantiates a new SecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    m := &ItemSecurityManagersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/security-managers";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSecurityManagersRequestBuilder instantiates a new SecurityManagersRequestBuilder and sets the default values.
func NewItemSecurityManagersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an administrator or security manager for the organization, and you must use an access token with the `read:org` scope.GitHub Apps must have the `administration` organization read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#list-security-manager-teams
func (m *ItemSecurityManagersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecurityManagersRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamSimpleable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.TeamSimpleable)
    }
    return val, nil
}
// Teams the teams property
func (m *ItemSecurityManagersRequestBuilder) Teams()(*ItemSecurityManagersTeamsRequestBuilder) {
    return NewItemSecurityManagersTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TeamsById gets an item from the github.com/octokit/kiota/.orgs.item.securityManagers.teams.item collection
func (m *ItemSecurityManagersRequestBuilder) TeamsById(id string)(*ItemSecurityManagersTeamsWithTeam_slugItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team_slug"] = id
    }
    return NewItemSecurityManagersTeamsWithTeam_slugItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation lists teams that are security managers for an organization. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an administrator or security manager for the organization, and you must use an access token with the `read:org` scope.GitHub Apps must have the `administration` organization read permission to use this endpoint.
func (m *ItemSecurityManagersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityManagersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
