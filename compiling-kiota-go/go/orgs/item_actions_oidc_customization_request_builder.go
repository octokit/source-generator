package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsOidcCustomizationRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\oidc\customization
type ItemActionsOidcCustomizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemActionsOidcCustomizationRequestBuilderInternal instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationRequestBuilder) {
    m := &ItemActionsOidcCustomizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/oidc/customization";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsOidcCustomizationRequestBuilder instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemActionsOidcCustomizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsOidcCustomizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsOidcCustomizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Sub the sub property
func (m *ItemActionsOidcCustomizationRequestBuilder) Sub()(*ItemActionsOidcCustomizationSubRequestBuilder) {
    return NewItemActionsOidcCustomizationSubRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
