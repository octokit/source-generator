package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsOidcRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\oidc
type ItemActionsOidcRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemActionsOidcRequestBuilderInternal instantiates a new OidcRequestBuilder and sets the default values.
func NewItemActionsOidcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcRequestBuilder) {
    m := &ItemActionsOidcRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/oidc";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsOidcRequestBuilder instantiates a new OidcRequestBuilder and sets the default values.
func NewItemActionsOidcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcRequestBuilderInternal(urlParams, requestAdapter)
}
// Customization the customization property
func (m *ItemActionsOidcRequestBuilder) Customization()(*ItemActionsOidcCustomizationRequestBuilder) {
    return NewItemActionsOidcCustomizationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
