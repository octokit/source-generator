package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemReleasesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\releases
type ItemItemReleasesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemReleasesRequestBuilderGetQueryParameters this returns a list of releases, which does not include regular Git tags that have not been associated with a release. To get a list of Git tags, use the [Repository Tags API](https://docs.github.com/rest/reference/repos#list-repository-tags).Information about published releases are available to everyone. Only users with push access will receive listings for draft releases.
type ItemItemReleasesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemReleasesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemReleasesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemReleasesRequestBuilderGetQueryParameters
}
// ItemItemReleasesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemReleasesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assets the assets property
func (m *ItemItemReleasesRequestBuilder) Assets()(*ItemItemReleasesAssetsRequestBuilder) {
    return NewItemItemReleasesAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssetsById gets an item from the github.com/octokit/kiota/.repos.item.item.releases.assets.item collection
func (m *ItemItemReleasesRequestBuilder) AssetsById(id string)(*ItemItemReleasesAssetsWithAsset_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["asset_id"] = id
    }
    return NewItemItemReleasesAssetsWithAsset_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemReleasesRequestBuilderInternal instantiates a new ReleasesRequestBuilder and sets the default values.
func NewItemItemReleasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesRequestBuilder) {
    m := &ItemItemReleasesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/releases{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemReleasesRequestBuilder instantiates a new ReleasesRequestBuilder and sets the default values.
func NewItemItemReleasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReleasesRequestBuilderInternal(urlParams, requestAdapter)
}
// GenerateNotes the generateNotes property
func (m *ItemItemReleasesRequestBuilder) GenerateNotes()(*ItemItemReleasesGenerateNotesRequestBuilder) {
    return NewItemItemReleasesGenerateNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get this returns a list of releases, which does not include regular Git tags that have not been associated with a release. To get a list of Git tags, use the [Repository Tags API](https://docs.github.com/rest/reference/repos#list-repository-tags).Information about published releases are available to everyone. Only users with push access will receive listings for draft releases.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/repos#list-releases
func (m *ItemItemReleasesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemReleasesRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Releaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateReleaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Releaseable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Releaseable)
    }
    return val, nil
}
// Latest the latest property
func (m *ItemItemReleasesRequestBuilder) Latest()(*ItemItemReleasesLatestRequestBuilder) {
    return NewItemItemReleasesLatestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Post users with push access to the repository can create a release.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. See "[Secondary rate limits](https://docs.github.com/rest/overview/resources-in-the-rest-api#secondary-rate-limits)" and "[Dealing with secondary rate limits](https://docs.github.com/rest/guides/best-practices-for-integrators#dealing-with-secondary-rate-limits)" for details.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/releases/releases#create-a-release
func (m *ItemItemReleasesRequestBuilder) Post(ctx context.Context, body ItemItemReleasesPostRequestBodyable, requestConfiguration *ItemItemReleasesRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Releaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateReleaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Releaseable), nil
}
// Tags the tags property
func (m *ItemItemReleasesRequestBuilder) Tags()(*ItemItemReleasesTagsRequestBuilder) {
    return NewItemItemReleasesTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById gets an item from the github.com/octokit/kiota/.repos.item.item.releases.tags.item collection
func (m *ItemItemReleasesRequestBuilder) TagsById(id string)(*ItemItemReleasesTagsWithTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tag"] = id
    }
    return NewItemItemReleasesTagsWithTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation this returns a list of releases, which does not include regular Git tags that have not been associated with a release. To get a list of Git tags, use the [Repository Tags API](https://docs.github.com/rest/reference/repos#list-repository-tags).Information about published releases are available to everyone. Only users with push access will receive listings for draft releases.
func (m *ItemItemReleasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemReleasesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation users with push access to the repository can create a release.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications). Creating content too quickly using this endpoint may result in secondary rate limiting. See "[Secondary rate limits](https://docs.github.com/rest/overview/resources-in-the-rest-api#secondary-rate-limits)" and "[Dealing with secondary rate limits](https://docs.github.com/rest/guides/best-practices-for-integrators#dealing-with-secondary-rate-limits)" for details.
func (m *ItemItemReleasesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemReleasesPostRequestBodyable, requestConfiguration *ItemItemReleasesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
