package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites\{check_suite_id}
type ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckRuns the checkRuns property
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) CheckRuns()(*ItemItemCheckSuitesItemCheckRunsRequestBuilder) {
    return NewItemItemCheckSuitesItemCheckRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal instantiates a new WithCheck_suite_ItemRequestBuilder and sets the default values.
func NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    m := &ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/check-suites/{check_suite_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder instantiates a new WithCheck_suite_ItemRequestBuilder and sets the default values.
func NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Gets a single check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#get-a-check-suite
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCheckSuiteFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable), nil
}
// Rerequest the rerequest property
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) Rerequest()(*ItemItemCheckSuitesItemRerequestRequestBuilder) {
    return NewItemItemCheckSuitesItemRerequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Gets a single check suite using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.
func (m *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
