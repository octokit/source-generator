package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCommunityProfileRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\community\profile
type ItemItemCommunityProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCommunityProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommunityProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCommunityProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemItemCommunityProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityProfileRequestBuilder) {
    m := &ItemItemCommunityProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/community/profile";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCommunityProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemItemCommunityProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommunityProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommunityProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all community profile metrics for a repository. The repository cannot be a fork.The returned metrics include an overall health score, the repository description, the presence of documentation, thedetected code of conduct, the detected license, and the presence of ISSUE\_TEMPLATE, PULL\_REQUEST\_TEMPLATE,README, and CONTRIBUTING files.The `health_percentage` score is defined as a percentage of how many ofthese four documents are present: README, CONTRIBUTING, LICENSE, andCODE_OF_CONDUCT. For example, if all four documents are present, thenthe `health_percentage` is `100`. If only one is present, then the`health_percentage` is `25`.`content_reports_enabled` is only returned for organization-owned repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/community#get-community-profile-metrics
func (m *ItemItemCommunityProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommunityProfileRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CommunityProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCommunityProfileFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CommunityProfileable), nil
}
// ToGetRequestInformation returns all community profile metrics for a repository. The repository cannot be a fork.The returned metrics include an overall health score, the repository description, the presence of documentation, thedetected code of conduct, the detected license, and the presence of ISSUE\_TEMPLATE, PULL\_REQUEST\_TEMPLATE,README, and CONTRIBUTING files.The `health_percentage` score is defined as a percentage of how many ofthese four documents are present: README, CONTRIBUTING, LICENSE, andCODE_OF_CONDUCT. For example, if all four documents are present, thenthe `health_percentage` is `100`. If only one is present, then the`health_percentage` is `25`.`content_reports_enabled` is only returned for organization-owned repositories.
func (m *ItemItemCommunityProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommunityProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
