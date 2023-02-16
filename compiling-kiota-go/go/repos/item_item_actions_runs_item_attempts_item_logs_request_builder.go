package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}\attempts\{attempt_number}\logs
type ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunsItemAttemptsItemLogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemAttemptsItemLogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilder instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsItemLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download an archive of log files for a specific workflow run attempt. This link expires after1 minute. Look for `Location:` in the response header to find the URL for the download. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#download-workflow-run-attempt-logs
func (m *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download an archive of log files for a specific workflow run attempt. This link expires after1 minute. Look for `Location:` in the response header to find the URL for the download. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemAttemptsItemLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
