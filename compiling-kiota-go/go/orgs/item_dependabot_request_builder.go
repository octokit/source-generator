package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot
type ItemDependabotRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemDependabotRequestBuilder) Alerts()(*ItemDependabotAlertsRequestBuilder) {
    return NewItemDependabotAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemDependabotRequestBuilderInternal instantiates a new DependabotRequestBuilder and sets the default values.
func NewItemDependabotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotRequestBuilder) {
    m := &ItemDependabotRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/dependabot";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDependabotRequestBuilder instantiates a new DependabotRequestBuilder and sets the default values.
func NewItemDependabotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotRequestBuilderInternal(urlParams, requestAdapter)
}
// Secrets the secrets property
func (m *ItemDependabotRequestBuilder) Secrets()(*ItemDependabotSecretsRequestBuilder) {
    return NewItemDependabotSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.orgs.item.dependabot.secrets.item collection
func (m *ItemDependabotRequestBuilder) SecretsById(id string)(*ItemDependabotSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemDependabotSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
