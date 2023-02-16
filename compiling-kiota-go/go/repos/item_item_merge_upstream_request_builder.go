package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemMergeUpstreamRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\merge-upstream
type ItemItemMergeUpstreamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemMergeUpstreamRequestBuilderInternal instantiates a new MergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    m := &ItemItemMergeUpstreamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/merge-upstream";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemMergeUpstreamRequestBuilder instantiates a new MergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemMergeUpstreamRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sync a branch of a forked repository to keep it up-to-date with the upstream repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branches#sync-a-fork-branch-with-the-upstream-repository
func (m *ItemItemMergeUpstreamRequestBuilder) Post(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MergedUpstreamable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateMergedUpstreamFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MergedUpstreamable), nil
}
// ToPostRequestInformation sync a branch of a forked repository to keep it up-to-date with the upstream repository.
func (m *ItemItemMergeUpstreamRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
