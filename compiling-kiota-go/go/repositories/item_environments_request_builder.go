package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments
type ItemEnvironmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemEnvironmentsRequestBuilderInternal instantiates a new EnvironmentsRequestBuilder and sets the default values.
func NewItemEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsRequestBuilder) {
    m := &ItemEnvironmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repositories/{repository_id}/environments";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEnvironmentsRequestBuilder instantiates a new EnvironmentsRequestBuilder and sets the default values.
func NewItemEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
