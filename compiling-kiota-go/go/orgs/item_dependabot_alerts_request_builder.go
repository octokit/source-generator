package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemDependabotAlertsRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\alerts
type ItemDependabotAlertsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDependabotAlertsRequestBuilderGetQueryParameters lists Dependabot alerts for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
type ItemDependabotAlertsRequestBuilderGetQueryParameters struct {
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
    // A comma-separated list of package names. If specified, only alerts for these packages will be returned.
    Package *string
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
// ItemDependabotAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDependabotAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDependabotAlertsRequestBuilderGetQueryParameters
}
// NewItemDependabotAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemDependabotAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotAlertsRequestBuilder) {
    m := &ItemDependabotAlertsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/dependabot/alerts{?state*,severity*,ecosystem*,package*,scope*,sort*,direction*,before*,after*,first*,last*,per_page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDependabotAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemDependabotAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists Dependabot alerts for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/alerts#list-dependabot-alerts-for-an-organization
func (m *ItemDependabotAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDependabotAlertsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertWithRepositoryable, error) {
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
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDependabotAlertWithRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertWithRepositoryable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotAlertWithRepositoryable)
    }
    return val, nil
}
// ToGetRequestInformation lists Dependabot alerts for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
func (m *ItemDependabotAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDependabotAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
