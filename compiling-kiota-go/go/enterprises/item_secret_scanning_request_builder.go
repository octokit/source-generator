package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSecretScanningRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\secret-scanning
type ItemSecretScanningRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Alerts the alerts property
func (m *ItemSecretScanningRequestBuilder) Alerts()(*ItemSecretScanningAlertsRequestBuilder) {
    return NewItemSecretScanningAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemSecretScanningRequestBuilderInternal instantiates a new SecretScanningRequestBuilder and sets the default values.
func NewItemSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningRequestBuilder) {
    m := &ItemSecretScanningRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/enterprises/{enterprise}/secret-scanning";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSecretScanningRequestBuilder instantiates a new SecretScanningRequestBuilder and sets the default values.
func NewItemSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
