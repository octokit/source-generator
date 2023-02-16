package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCodeScanningAlertsItemInstancesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\alerts\{alert_number}\instances
type ItemItemCodeScanningAlertsItemInstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetQueryParameters lists all instances of the specified code scanning alert.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.GitHub Apps must have the `security_events` read permission to use this endpoint.
type ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string
}
// ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetQueryParameters
}
// NewItemItemCodeScanningAlertsItemInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemInstancesRequestBuilder) {
    m := &ItemItemCodeScanningAlertsItemInstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/code-scanning/alerts/{alert_number}/instances{?page*,per_page*,ref*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodeScanningAlertsItemInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAlertsItemInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all instances of the specified code scanning alert.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.GitHub Apps must have the `security_events` read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/code-scanning#list-instances-of-a-code-scanning-alert
func (m *ItemItemCodeScanningAlertsItemInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAlertInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateInstances503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningAlertInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAlertInstanceable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningAlertInstanceable)
    }
    return val, nil
}
// ToGetRequestInformation lists all instances of the specified code scanning alert.You must use an access token with the `security_events` scope to use this endpoint with private repos,the `public_repo` scope also grants permission to read security events on public repos only.GitHub Apps must have the `security_events` read permission to use this endpoint.
func (m *ItemItemCodeScanningAlertsItemInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningAlertsItemInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
