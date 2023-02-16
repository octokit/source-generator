package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PackagesWithPackage_typeItemRequestBuilder builds and executes requests for operations under \user\packages\{package_type}
type PackagesWithPackage_typeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewPackagesWithPackage_typeItemRequestBuilderInternal instantiates a new WithPackage_typeItemRequestBuilder and sets the default values.
func NewPackagesWithPackage_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesWithPackage_typeItemRequestBuilder) {
    m := &PackagesWithPackage_typeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/packages/{package_type}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPackagesWithPackage_typeItemRequestBuilder instantiates a new WithPackage_typeItemRequestBuilder and sets the default values.
func NewPackagesWithPackage_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesWithPackage_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesWithPackage_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
