package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemZipballRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\zipball
type ItemItemZipballRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemZipballRequestBuilderInternal instantiates a new ZipballRequestBuilder and sets the default values.
func NewItemItemZipballRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballRequestBuilder) {
    m := &ItemItemZipballRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/zipball";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemZipballRequestBuilder instantiates a new ZipballRequestBuilder and sets the default values.
func NewItemItemZipballRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemZipballRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemZipballRequestBuilderInternal(urlParams, requestAdapter)
}
