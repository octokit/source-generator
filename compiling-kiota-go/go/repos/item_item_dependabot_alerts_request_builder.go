package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemDependabotAlertsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependabot\alerts
type ItemItemDependabotAlertsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemDependabotAlertsRequestBuilderGetQueryParameters you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
type ItemItemDependabotAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for results after this cursor.
    After *string
    // A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for results before this cursor.
    Before *string
    // The direction to sort the results by.
    Direction *string
    // A comma-separated list of ecosystems. If specified, only alerts for these ecosystems will be returned.Can be: `composer`, `go`, `maven`, `npm`, `nuget`, `pip`, `pub`, `rubygems`, `rust`
    Ecosystem *string
    // **Deprecated**. The number of results per page (max 100), starting from the first matching result.This parameter must not be used in combination with `last`.Instead, use `per_page` in combination with `after` to fetch the first page of results.
    First *int32
    // **Deprecated**. The number of results per page (max 100), starting from the last matching result.This parameter must not be used in combination with `first`.Instead, use `per_page` in combination with `before` to fetch the last page of results.
    Last *int32
    // A comma-separated list of full manifest paths. If specified, only alerts for these manifests will be returned.
    Manifest *string
    // A comma-separated list of package names. If specified, only alerts for these packages will be returned.
    Package *string
    // **Deprecated**. Page number of the results to fetch. Use cursor-based pagination with `before` or `after` instead.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.
    Scope *string
    // A comma-separated list of severities. If specified, only alerts with these severities will be returned.Can be: `low`, `medium`, `high`, `critical`
    Severity *string
    // The property by which to sort the results.`created` means when the alert was created.`updated` means when the alert's state last changed.
    Sort *string
    // A comma-separated list of states. If specified, only alerts with these states will be returned.Can be: `dismissed`, `fixed`, `open`
    State *string
}
// ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemDependabotAlertsRequestBuilderGetQueryParameters
}
// NewItemItemDependabotAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    m := &ItemItemDependabotAlertsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/dependabot/alerts{?state*,severity*,ecosystem*,package*,manifest*,scope*,sort*,direction*,page*,per_page*,before*,after*,first*,last*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDependabotAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependabotAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/dependabot#list-dependabot-alerts-for-a-repository
func (m *ItemItemDependabotAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDependabotAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertable)
    }
    return val, nil
}
// ToGetRequestInformation you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
func (m *ItemItemDependabotAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
