package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsPermissionsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\permissions
type ItemItemActionsPermissionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsPermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsPermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsPermissionsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsPermissionsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Access the access property
func (m *ItemItemActionsPermissionsRequestBuilder) Access()(*ItemItemActionsPermissionsAccessRequestBuilder) {
    return NewItemItemActionsPermissionsAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemActionsPermissionsRequestBuilderInternal instantiates a new PermissionsRequestBuilder and sets the default values.
func NewItemItemActionsPermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsPermissionsRequestBuilder) {
    m := &ItemItemActionsPermissionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/permissions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsPermissionsRequestBuilder instantiates a new PermissionsRequestBuilder and sets the default values.
func NewItemItemActionsPermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsPermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsPermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the GitHub Actions permissions policy for a repository, including whether GitHub Actions is enabled and the actions and reusable workflows allowed to run in the repository.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration` repository permission to use this API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-github-actions-permissions-for-a-repository
func (m *ItemItemActionsPermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsPermissionsRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsRepositoryPermissionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateActionsRepositoryPermissionsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsRepositoryPermissionsable), nil
}
// Put sets the GitHub Actions permissions policy for enabling GitHub Actions and allowed actions and reusable workflows in the repository.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration` repository permission to use this API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#set-github-actions-permissions-for-a-repository
func (m *ItemItemActionsPermissionsRequestBuilder) Put(ctx context.Context, body ItemItemActionsPermissionsPutRequestBodyable, requestConfiguration *ItemItemActionsPermissionsRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// SelectedActions the selectedActions property
func (m *ItemItemActionsPermissionsRequestBuilder) SelectedActions()(*ItemItemActionsPermissionsSelectedActionsRequestBuilder) {
    return NewItemItemActionsPermissionsSelectedActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation gets the GitHub Actions permissions policy for a repository, including whether GitHub Actions is enabled and the actions and reusable workflows allowed to run in the repository.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration` repository permission to use this API.
func (m *ItemItemActionsPermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsPermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation sets the GitHub Actions permissions policy for enabling GitHub Actions and allowed actions and reusable workflows in the repository.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration` repository permission to use this API.
func (m *ItemItemActionsPermissionsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemActionsPermissionsPutRequestBodyable, requestConfiguration *ItemItemActionsPermissionsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
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
// Workflow the workflow property
func (m *ItemItemActionsPermissionsRequestBuilder) Workflow()(*ItemItemActionsPermissionsWorkflowRequestBuilder) {
    return NewItemItemActionsPermissionsWorkflowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
