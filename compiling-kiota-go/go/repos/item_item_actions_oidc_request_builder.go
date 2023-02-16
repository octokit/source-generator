package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOidcRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\oidc
type ItemItemActionsOidcRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemActionsOidcRequestBuilderInternal instantiates a new OidcRequestBuilder and sets the default values.
func NewItemItemActionsOidcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcRequestBuilder) {
    m := &ItemItemActionsOidcRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/oidc";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsOidcRequestBuilder instantiates a new OidcRequestBuilder and sets the default values.
func NewItemItemActionsOidcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOidcRequestBuilderInternal(urlParams, requestAdapter)
}
// Customization the customization property
func (m *ItemItemActionsOidcRequestBuilder) Customization()(*ItemItemActionsOidcCustomizationRequestBuilder) {
    return NewItemItemActionsOidcCustomizationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
