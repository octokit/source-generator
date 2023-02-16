package appmanifests

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppManifestsRequestBuilder builds and executes requests for operations under \app-manifests
type AppManifestsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewAppManifestsRequestBuilderInternal instantiates a new AppManifestsRequestBuilder and sets the default values.
func NewAppManifestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManifestsRequestBuilder) {
    m := &AppManifestsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app-manifests";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAppManifestsRequestBuilder instantiates a new AppManifestsRequestBuilder and sets the default values.
func NewAppManifestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManifestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppManifestsRequestBuilderInternal(urlParams, requestAdapter)
}
