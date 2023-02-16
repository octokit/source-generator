package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCodeScanningRequestBuilder builds and executes requests for operations under \orgs\{org}\code-scanning
type ItemCodeScanningRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemCodeScanningRequestBuilder) Alerts()(*ItemCodeScanningAlertsRequestBuilder) {
    return NewItemCodeScanningAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCodeScanningRequestBuilderInternal instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemCodeScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningRequestBuilder) {
    m := &ItemCodeScanningRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/code-scanning";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCodeScanningRequestBuilder instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemCodeScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeScanningRequestBuilderInternal(urlParams, requestAdapter)
}
