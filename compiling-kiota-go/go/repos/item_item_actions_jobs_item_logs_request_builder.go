package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsJobsItemLogsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\jobs\{job_id}\logs
type ItemItemActionsJobsItemLogsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsJobsItemLogsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsJobsItemLogsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsJobsItemLogsRequestBuilderInternal instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsJobsItemLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsItemLogsRequestBuilder) {
    m := &ItemItemActionsJobsItemLogsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/jobs/{job_id}/logs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsJobsItemLogsRequestBuilder instantiates a new LogsRequestBuilder and sets the default values.
func NewItemItemActionsJobsItemLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsItemLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsJobsItemLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download a plain text file of logs for a workflow job. This link expires after 1 minute. Lookfor `Location:` in the response header to find the URL for the download. Anyone with read access to the repository canuse this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps musthave the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#download-job-logs-for-a-workflow-run
func (m *ItemItemActionsJobsItemLogsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsJobsItemLogsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation gets a redirect URL to download a plain text file of logs for a workflow job. This link expires after 1 minute. Lookfor `Location:` in the response header to find the URL for the download. Anyone with read access to the repository canuse this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps musthave the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsJobsItemLogsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsJobsItemLogsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
