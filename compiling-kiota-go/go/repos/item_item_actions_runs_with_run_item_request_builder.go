package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRunsWithRun_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}
type ItemItemActionsRunsWithRun_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool
}
// ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsWithRun_ItemRequestBuilderGetQueryParameters
}
// Approvals the approvals property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approvals()(*ItemItemActionsRunsItemApprovalsRequestBuilder) {
    return NewItemItemActionsRunsItemApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Approve the approve property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Approve()(*ItemItemActionsRunsItemApproveRequestBuilder) {
    return NewItemItemActionsRunsItemApproveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Artifacts the artifacts property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Artifacts()(*ItemItemActionsRunsItemArtifactsRequestBuilder) {
    return NewItemItemActionsRunsItemArtifactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Attempts the attempts property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Attempts()(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttemptsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.runs.item.attempts.item collection
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) AttemptsById(id string)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attempt_number"] = id
    }
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Cancel the cancel property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Cancel()(*ItemItemActionsRunsItemCancelRequestBuilder) {
    return NewItemItemActionsRunsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal instantiates a new WithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    m := &ItemItemActionsRunsWithRun_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}{?exclude_pull_requests*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunsWithRun_ItemRequestBuilder instantiates a new WithRun_ItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsWithRun_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific workflow run. Anyone with write access to the repository can use this endpoint. If the repository isprivate you must use an access token with the `repo` scope. GitHub Apps must have the `actions:write` permission to usethis endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#delete-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-a-workflow-run
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WorkflowRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateWorkflowRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WorkflowRunable), nil
}
// Jobs the jobs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Logs the logs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Logs()(*ItemItemActionsRunsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Pending_deployments the pending_deployments property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Pending_deployments()(*ItemItemActionsRunsItemPending_deploymentsRequestBuilder) {
    return NewItemItemActionsRunsItemPending_deploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Rerun the rerun property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Rerun()(*ItemItemActionsRunsItemRerunRequestBuilder) {
    return NewItemItemActionsRunsItemRerunRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RerunFailedJobs the rerunFailedJobs property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) RerunFailedJobs()(*ItemItemActionsRunsItemRerunFailedJobsRequestBuilder) {
    return NewItemItemActionsRunsItemRerunFailedJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Timing the timing property
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) Timing()(*ItemItemActionsRunsItemTimingRequestBuilder) {
    return NewItemItemActionsRunsItemTimingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete a specific workflow run. Anyone with write access to the repository can use this endpoint. If the repository isprivate you must use an access token with the `repo` scope. GitHub Apps must have the `actions:write` permission to usethis endpoint.
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation gets a specific workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsRunsWithRun_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsWithRun_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
