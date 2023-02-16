package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stats
type ItemItemStatsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Code_frequency the code_frequency property
func (m *ItemItemStatsRequestBuilder) Code_frequency()(*ItemItemStatsCode_frequencyRequestBuilder) {
    return NewItemItemStatsCode_frequencyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Commit_activity the commit_activity property
func (m *ItemItemStatsRequestBuilder) Commit_activity()(*ItemItemStatsCommit_activityRequestBuilder) {
    return NewItemItemStatsCommit_activityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemStatsRequestBuilderInternal instantiates a new StatsRequestBuilder and sets the default values.
func NewItemItemStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsRequestBuilder) {
    m := &ItemItemStatsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/stats";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemStatsRequestBuilder instantiates a new StatsRequestBuilder and sets the default values.
func NewItemItemStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Contributors the contributors property
func (m *ItemItemStatsRequestBuilder) Contributors()(*ItemItemStatsContributorsRequestBuilder) {
    return NewItemItemStatsContributorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Participation the participation property
func (m *ItemItemStatsRequestBuilder) Participation()(*ItemItemStatsParticipationRequestBuilder) {
    return NewItemItemStatsParticipationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Punch_card the punch_card property
func (m *ItemItemStatsRequestBuilder) Punch_card()(*ItemItemStatsPunch_cardRequestBuilder) {
    return NewItemItemStatsPunch_cardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
