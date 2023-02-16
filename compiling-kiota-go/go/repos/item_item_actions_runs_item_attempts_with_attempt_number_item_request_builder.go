package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}\attempts\{attempt_number}
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters struct {
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool
}
// ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetQueryParameters
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal instantiates a new WithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}{?exclude_pull_requests*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder instantiates a new WithAttempt_numberItemRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#get-a-workflow-run-attempt
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WorkflowRunable, error) {
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
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Jobs()(*ItemItemActionsRunsItemAttemptsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Logs the logs property
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) Logs()(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation gets a specific workflow run attempt. Anyone with read access to the repositorycan use this endpoint. If the repository is private you must use an access tokenwith the `repo` scope. GitHub Apps must have the `actions:read` permission touse this endpoint.
func (m *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
