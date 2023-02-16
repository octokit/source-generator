package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsWorkflowsItemRunsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\workflows\{workflow_id}\runs
type ItemItemActionsWorkflowsItemRunsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
type ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters struct {
    // Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run.
    Actor *string
    // Returns workflow runs associated with a branch. Use the name of the branch of the `push`.
    Branch *string
    // Returns workflow runs with the `check_suite_id` that you specify.
    Check_suite_id *int32
    // Returns workflow runs created within the given date-time range. For more information on the syntax, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see "[Events that trigger workflows](https://docs.github.com/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows)."
    Event *string
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool
    // Only returns workflow runs that are associated with the specified `head_sha`.
    Head_sha *string
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub can set a status of `waiting` or `requested`.
    Status *string
}
// ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    m := &ItemItemActionsWorkflowsItemRunsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs{?actor*,branch*,event*,status*,per_page*,page*,created*,exclude_pull_requests*,check_suite_id*,head_sha*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilder instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-workflow-runs
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(ItemItemActionsWorkflowsItemRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsWorkflowsItemRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsWorkflowsItemRunsResponseable), nil
}
// ToGetRequestInformation list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
