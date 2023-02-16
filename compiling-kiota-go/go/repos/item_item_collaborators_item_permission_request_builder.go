package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCollaboratorsItemPermissionRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\collaborators\{username}\permission
type ItemItemCollaboratorsItemPermissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCollaboratorsItemPermissionRequestBuilderInternal instantiates a new PermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    m := &ItemItemCollaboratorsItemPermissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/collaborators/{username}/permission";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCollaboratorsItemPermissionRequestBuilder instantiates a new PermissionRequestBuilder and sets the default values.
func NewItemItemCollaboratorsItemPermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsItemPermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsItemPermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks the repository permission of a collaborator. The possible repository permissions are `admin`, `write`, `read`, and `none`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#get-repository-permissions-for-a-user
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepositoryCollaboratorPermissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRepositoryCollaboratorPermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepositoryCollaboratorPermissionable), nil
}
// ToGetRequestInformation checks the repository permission of a collaborator. The possible repository permissions are `admin`, `write`, `read`, and `none`.
func (m *ItemItemCollaboratorsItemPermissionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCollaboratorsItemPermissionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
