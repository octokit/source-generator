package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTarballRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\tarball
type ItemItemTarballRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemTarballRequestBuilderInternal instantiates a new TarballRequestBuilder and sets the default values.
func NewItemItemTarballRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballRequestBuilder) {
    m := &ItemItemTarballRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/tarball";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemTarballRequestBuilder instantiates a new TarballRequestBuilder and sets the default values.
func NewItemItemTarballRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTarballRequestBuilderInternal(urlParams, requestAdapter)
}
