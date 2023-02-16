package enterprises

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithEnterpriseItemRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}
type WithEnterpriseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithEnterpriseItemRequestBuilderInternal instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    m := &WithEnterpriseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/enterprises/{enterprise}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithEnterpriseItemRequestBuilder instantiates a new WithEnterpriseItemRequestBuilder and sets the default values.
func NewWithEnterpriseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithEnterpriseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithEnterpriseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Dependabot the dependabot property
func (m *WithEnterpriseItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretScanning the secretScanning property
func (m *WithEnterpriseItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
