package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithSecurity_productItemRequestBuilder builds and executes requests for operations under \orgs\{org}\{security_product}
type ItemWithSecurity_productItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemWithSecurity_productItemRequestBuilderInternal instantiates a new WithSecurity_productItemRequestBuilder and sets the default values.
func NewItemWithSecurity_productItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithSecurity_productItemRequestBuilder) {
    m := &ItemWithSecurity_productItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/{security_product}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemWithSecurity_productItemRequestBuilder instantiates a new WithSecurity_productItemRequestBuilder and sets the default values.
func NewItemWithSecurity_productItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithSecurity_productItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithSecurity_productItemRequestBuilderInternal(urlParams, requestAdapter)
}
