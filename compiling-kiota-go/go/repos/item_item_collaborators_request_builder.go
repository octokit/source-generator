package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCollaboratorsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\collaborators
type ItemItemCollaboratorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCollaboratorsRequestBuilderGetQueryParameters for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
type ItemItemCollaboratorsRequestBuilderGetQueryParameters struct {
    // Filter collaborators returned by their affiliation. `outside` means all outside collaborators of an organization-owned repository. `direct` means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. `all` means all collaborators the authenticated user can see.
    Affiliation *string
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // Filter collaborators by the permissions they have on the repository. If not specified, all collaborators will be returned.
    Permission *string
}
// ItemItemCollaboratorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCollaboratorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCollaboratorsRequestBuilderGetQueryParameters
}
// NewItemItemCollaboratorsRequestBuilderInternal instantiates a new CollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    m := &ItemItemCollaboratorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/collaborators{?affiliation*,permission*,per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCollaboratorsRequestBuilder instantiates a new CollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#list-repository-collaborators
func (m *ItemItemCollaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCollaboratorsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Collaboratorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCollaboratorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Collaboratorable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Collaboratorable)
    }
    return val, nil
}
// ToGetRequestInformation for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
func (m *ItemItemCollaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCollaboratorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
