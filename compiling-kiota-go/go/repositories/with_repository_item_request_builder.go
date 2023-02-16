package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithRepository_ItemRequestBuilder builds and executes requests for operations under \repositories\{repository_id}
type WithRepository_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWithRepository_ItemRequestBuilderInternal instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewWithRepository_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithRepository_ItemRequestBuilder) {
    m := &WithRepository_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repositories/{repository_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithRepository_ItemRequestBuilder instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewWithRepository_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithRepository_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithRepository_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Environments the environments property
func (m *WithRepository_ItemRequestBuilder) Environments()(*ItemEnvironmentsRequestBuilder) {
    return NewItemEnvironmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EnvironmentsById gets an item from the github.com/octokit/kiota/.repositories.item.environments.item collection
func (m *WithRepository_ItemRequestBuilder) EnvironmentsById(id string)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["environment_name"] = id
    }
    return NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
