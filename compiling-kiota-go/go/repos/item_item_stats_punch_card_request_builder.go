package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsPunch_cardRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stats\punch_card
type ItemItemStatsPunch_cardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatsPunch_cardRequestBuilderInternal instantiates a new Punch_cardRequestBuilder and sets the default values.
func NewItemItemStatsPunch_cardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsPunch_cardRequestBuilder) {
    m := &ItemItemStatsPunch_cardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/stats/punch_card";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemStatsPunch_cardRequestBuilder instantiates a new Punch_cardRequestBuilder and sets the default values.
func NewItemItemStatsPunch_cardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsPunch_cardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsPunch_cardRequestBuilderInternal(urlParams, requestAdapter)
}
// Get each array contains the day number, hour number, and number of commits:*   `0-6`: Sunday - Saturday*   `0-23`: Hour of day*   Number of commitsFor example, `[2, 14, 25]` indicates that there were 25 total commits, during the 2:00pm hour on Tuesdays. All times are based on the time zone of individual commits.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/statistics/repos#get-the-hourly-commit-count-for-each-day
func (m *ItemItemStatsPunch_cardRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration)([]int32, error) {
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
// ToGetRequestInformation each array contains the day number, hour number, and number of commits:*   `0-6`: Sunday - Saturday*   `0-23`: Hour of day*   Number of commitsFor example, `[2, 14, 25]` indicates that there were 25 total commits, during the 2:00pm hour on Tuesdays. All times are based on the time zone of individual commits.
func (m *ItemItemStatsPunch_cardRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
