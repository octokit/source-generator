package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemSecretScanningRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\secret-scanning
type ItemItemSecretScanningRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemItemSecretScanningRequestBuilder) Alerts()(*ItemItemSecretScanningAlertsRequestBuilder) {
    return NewItemItemSecretScanningAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AlertsById gets an item from the github.com/octokit/kiota/.repos.item.item.secretScanning.alerts.item collection
func (m *ItemItemSecretScanningRequestBuilder) AlertsById(id string)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert_number"] = id
    }
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemSecretScanningRequestBuilderInternal instantiates a new SecretScanningRequestBuilder and sets the default values.
func NewItemItemSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningRequestBuilder) {
    m := &ItemItemSecretScanningRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/secret-scanning";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemSecretScanningRequestBuilder instantiates a new SecretScanningRequestBuilder and sets the default values.
func NewItemItemSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
