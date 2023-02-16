package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\workflows\{workflow_id}
type ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal instantiates a new WithWorkflow_ItemRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    m := &ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/workflows/{workflow_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder instantiates a new WithWorkflow_ItemRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Disable the disable property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Disable()(*ItemItemActionsWorkflowsItemDisableRequestBuilder) {
    return NewItemItemActionsWorkflowsItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Dispatches the dispatches property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Dispatches()(*ItemItemActionsWorkflowsItemDispatchesRequestBuilder) {
    return NewItemItemActionsWorkflowsItemDispatchesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Enable the enable property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Enable()(*ItemItemActionsWorkflowsItemEnableRequestBuilder) {
    return NewItemItemActionsWorkflowsItemEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get gets a specific workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-a-workflow
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateWorkflowFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Workflowable), nil
}
// Runs the runs property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Runs()(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    return NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Timing the timing property
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) Timing()(*ItemItemActionsWorkflowsItemTimingRequestBuilder) {
    return NewItemItemActionsWorkflowsItemTimingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation gets a specific workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
