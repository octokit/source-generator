package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemEventsOrgsWithOrgItemRequestBuilder builds and executes requests for operations under \users\{username}\events\orgs\{org}
type ItemEventsOrgsWithOrgItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters this is the user's organization dashboard. You must be authenticated as the user to view this.
type ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters
}
// NewItemEventsOrgsWithOrgItemRequestBuilderInternal instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    m := &ItemEventsOrgsWithOrgItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}/events/orgs/{org}{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEventsOrgsWithOrgItemRequestBuilder instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsOrgsWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this is the user's organization dashboard. You must be authenticated as the user to view this.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/activity#list-organization-events-for-the-authenticated-user
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Eventable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Eventable)
    }
    return val, nil
}
// ToGetRequestInformation this is the user's organization dashboard. You must be authenticated as the user to view this.
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
