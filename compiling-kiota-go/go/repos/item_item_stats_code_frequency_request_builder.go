package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsCode_frequencyRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stats\code_frequency
type ItemItemStatsCode_frequencyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemStatsCode_frequencyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatsCode_frequencyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatsCode_frequencyRequestBuilderInternal instantiates a new Code_frequencyRequestBuilder and sets the default values.
func NewItemItemStatsCode_frequencyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCode_frequencyRequestBuilder) {
    m := &ItemItemStatsCode_frequencyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/stats/code_frequency";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemStatsCode_frequencyRequestBuilder instantiates a new Code_frequencyRequestBuilder and sets the default values.
func NewItemItemStatsCode_frequencyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsCode_frequencyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsCode_frequencyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a weekly aggregate of the number of additions and deletions pushed to a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-the-weekly-commit-activity
func (m *ItemItemStatsCode_frequencyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStatsCode_frequencyRequestBuilderGetRequestConfiguration)([]int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        val[i] = *(v.(*int32))
    }
    return val, nil
}
// ToGetRequestInformation returns a weekly aggregate of the number of additions and deletions pushed to a repository.
func (m *ItemItemStatsCode_frequencyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStatsCode_frequencyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
