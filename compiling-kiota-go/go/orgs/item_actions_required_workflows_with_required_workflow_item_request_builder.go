package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\required_workflows\{required_workflow_id}
type ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderInternal instantiates a new WithRequired_workflow_ItemRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) {
    m := &ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/required_workflows/{required_workflow_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder instantiates a new WithRequired_workflow_ItemRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a required workflow configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#delete-a-required-workflow
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a required workflow configured in an organization.You must authenticate using an access token with the `read:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-a-required-workflow
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRequiredWorkflowFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable), nil
}
// Patch update a required workflow in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#update-a-required-workflow
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) Patch(ctx context.Context, body ItemActionsRequired_workflowsItemWithRequired_workflow_PatchRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRequiredWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable), nil
}
// Repositories the repositories property
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) Repositories()(*ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) {
    return NewItemActionsRequired_workflowsItemRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RepositoriesById gets an item from the github.com/octokit/kiota/.orgs.item.actions.required_workflows.item.repositories.item collection
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) RepositoriesById(id string)(*ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["repository_id"] = id
    }
    return NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation deletes a required workflow configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get a required workflow configured in an organization.You must authenticate using an access token with the `read:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a required workflow in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemActionsRequired_workflowsItemWithRequired_workflow_PatchRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
