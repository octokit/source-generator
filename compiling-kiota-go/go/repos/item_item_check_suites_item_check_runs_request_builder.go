package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCheckSuitesItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites\{check_suite_id}\check-runs
type ItemItemCheckSuitesItemCheckRunsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.
type ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters struct {
    // Returns check runs with the specified `name`.
    Check_name *string
    // Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs.
    Filter *string
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // Returns check runs with the specified `status`.
    Status *string
}
// ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetQueryParameters
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    m := &ItemItemCheckSuitesItemCheckRunsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs{?check_name*,status*,filter*,per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCheckSuitesItemCheckRunsRequestBuilder instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#list-check-runs-in-a-check-suite
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration)(ItemItemCheckSuitesItemCheckRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCheckSuitesItemCheckRunsResponseable), nil
}
// ToGetRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Lists check runs for a check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.
func (m *ItemItemCheckSuitesItemCheckRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCheckSuitesItemCheckRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
