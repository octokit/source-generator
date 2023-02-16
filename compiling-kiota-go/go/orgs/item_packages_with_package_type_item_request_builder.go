package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPackagesWithPackage_typeItemRequestBuilder builds and executes requests for operations under \orgs\{org}\packages\{package_type}
type ItemPackagesWithPackage_typeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemPackagesWithPackage_typeItemRequestBuilderInternal instantiates a new WithPackage_typeItemRequestBuilder and sets the default values.
func NewItemPackagesWithPackage_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    m := &ItemPackagesWithPackage_typeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/packages/{package_type}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemPackagesWithPackage_typeItemRequestBuilder instantiates a new WithPackage_typeItemRequestBuilder and sets the default values.
func NewItemPackagesWithPackage_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPackagesWithPackage_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
