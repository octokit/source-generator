package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependabotRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependabot
type ItemItemDependabotRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemItemDependabotRequestBuilder) Alerts()(*ItemItemDependabotAlertsRequestBuilder) {
    return NewItemItemDependabotAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AlertsById gets an item from the github.com/octokit/kiota/.repos.item.item.dependabot.alerts.item collection
func (m *ItemItemDependabotRequestBuilder) AlertsById(id string)(*ItemItemDependabotAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert_number"] = id
    }
    return NewItemItemDependabotAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemDependabotRequestBuilderInternal instantiates a new DependabotRequestBuilder and sets the default values.
func NewItemItemDependabotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotRequestBuilder) {
    m := &ItemItemDependabotRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/dependabot";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDependabotRequestBuilder instantiates a new DependabotRequestBuilder and sets the default values.
func NewItemItemDependabotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependabotRequestBuilderInternal(urlParams, requestAdapter)
}
// Secrets the secrets property
func (m *ItemItemDependabotRequestBuilder) Secrets()(*ItemItemDependabotSecretsRequestBuilder) {
    return NewItemItemDependabotSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.repos.item.item.dependabot.secrets.item collection
func (m *ItemItemDependabotRequestBuilder) SecretsById(id string)(*ItemItemDependabotSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemItemDependabotSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
