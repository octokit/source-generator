package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsOidcCustomizationRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\oidc\customization
type ItemItemActionsOidcCustomizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemActionsOidcCustomizationRequestBuilderInternal instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemItemActionsOidcCustomizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcCustomizationRequestBuilder) {
    m := &ItemItemActionsOidcCustomizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/oidc/customization";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsOidcCustomizationRequestBuilder instantiates a new CustomizationRequestBuilder and sets the default values.
func NewItemItemActionsOidcCustomizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsOidcCustomizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsOidcCustomizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Sub the sub property
func (m *ItemItemActionsOidcCustomizationRequestBuilder) Sub()(*ItemItemActionsOidcCustomizationSubRequestBuilder) {
    return NewItemItemActionsOidcCustomizationSubRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
