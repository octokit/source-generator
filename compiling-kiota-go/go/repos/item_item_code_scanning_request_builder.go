package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning
type ItemItemCodeScanningRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemItemCodeScanningRequestBuilder) Alerts()(*ItemItemCodeScanningAlertsRequestBuilder) {
    return NewItemItemCodeScanningAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AlertsById gets an item from the github.com/octokit/kiota/.repos.item.item.codeScanning.alerts.item collection
func (m *ItemItemCodeScanningRequestBuilder) AlertsById(id string)(*ItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert_number"] = id
    }
    return NewItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Analyses the analyses property
func (m *ItemItemCodeScanningRequestBuilder) Analyses()(*ItemItemCodeScanningAnalysesRequestBuilder) {
    return NewItemItemCodeScanningAnalysesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AnalysesById gets an item from the github.com/octokit/kiota/.repos.item.item.codeScanning.analyses.item collection
func (m *ItemItemCodeScanningRequestBuilder) AnalysesById(id string)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["analysis_id"] = id
    }
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Codeql the codeql property
func (m *ItemItemCodeScanningRequestBuilder) Codeql()(*ItemItemCodeScanningCodeqlRequestBuilder) {
    return NewItemItemCodeScanningCodeqlRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemCodeScanningRequestBuilderInternal instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    m := &ItemItemCodeScanningRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/code-scanning";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodeScanningRequestBuilder instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningRequestBuilderInternal(urlParams, requestAdapter)
}
// Sarifs the sarifs property
func (m *ItemItemCodeScanningRequestBuilder) Sarifs()(*ItemItemCodeScanningSarifsRequestBuilder) {
    return NewItemItemCodeScanningSarifsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SarifsById gets an item from the github.com/octokit/kiota/.repos.item.item.codeScanning.sarifs.item collection
func (m *ItemItemCodeScanningRequestBuilder) SarifsById(id string)(*ItemItemCodeScanningSarifsWithSarif_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sarif_id"] = id
    }
    return NewItemItemCodeScanningSarifsWithSarif_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
