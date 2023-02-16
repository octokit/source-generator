package repositories

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsWithEnvironment_nameItemRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}
type ItemEnvironmentsWithEnvironment_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    m := &ItemEnvironmentsWithEnvironment_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repositories/{repository_id}/environments/{environment_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEnvironmentsWithEnvironment_nameItemRequestBuilder instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsWithEnvironment_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Secrets the secrets property
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Secrets()(*ItemEnvironmentsItemSecretsRequestBuilder) {
    return NewItemEnvironmentsItemSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.repositories.item.environments.item.secrets.item collection
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) SecretsById(id string)(*ItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemEnvironmentsItemSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Variables the variables property
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Variables()(*ItemEnvironmentsItemVariablesRequestBuilder) {
    return NewItemEnvironmentsItemVariablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VariablesById gets an item from the github.com/octokit/kiota/.repositories.item.environments.item.variables.item collection
func (m *ItemEnvironmentsWithEnvironment_nameItemRequestBuilder) VariablesById(id string)(*ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["name"] = id
    }
    return NewItemEnvironmentsItemVariablesWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
