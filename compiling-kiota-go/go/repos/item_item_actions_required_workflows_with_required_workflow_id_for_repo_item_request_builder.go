package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder builds and executes requests for operations under \repos\{org}\{repo}\actions\required_workflows\{required_workflow_id_for_repo}
type ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderInternal instantiates a new WithRequired_workflow_id_for_repoItemRequestBuilder and sets the default values.
func NewItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) {
    m := &ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{org}/{repo}/actions/required_workflows/{required_workflow_id_for_repo}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder instantiates a new WithRequired_workflow_id_for_repoItemRequestBuilder and sets the default values.
func NewItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a specific required workflow present in a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-repository-required-workflow
func (m *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRepoRequiredWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable), nil
}
// Runs the runs property
func (m *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) Runs()(*ItemItemActionsRequired_workflowsItemRunsRequestBuilder) {
    return NewItemItemActionsRequired_workflowsItemRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Timing the timing property
func (m *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) Timing()(*ItemItemActionsRequired_workflowsItemTimingRequestBuilder) {
    return NewItemItemActionsRequired_workflowsItemTimingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation gets a specific required workflow present in a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRequired_workflowsWithRequired_workflow_id_for_repoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
