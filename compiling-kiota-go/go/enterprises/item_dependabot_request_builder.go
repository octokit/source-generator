package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\dependabot
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
    m.urlTemplate = "{+baseurl}/enterprises/{enterprise}/dependabot";
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
