package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions
type ItemActionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Cache the cache property
func (m *ItemActionsRequestBuilder) Cache()(*ItemActionsCacheRequestBuilder) {
    return NewItemActionsCacheRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemActionsRequestBuilderInternal instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequestBuilder) {
    m := &ItemActionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRequestBuilder instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Oidc the oidc property
func (m *ItemActionsRequestBuilder) Oidc()(*ItemActionsOidcRequestBuilder) {
    return NewItemActionsOidcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Permissions the permissions property
func (m *ItemActionsRequestBuilder) Permissions()(*ItemActionsPermissionsRequestBuilder) {
    return NewItemActionsPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Required_workflows the required_workflows property
func (m *ItemActionsRequestBuilder) Required_workflows()(*ItemActionsRequired_workflowsRequestBuilder) {
    return NewItemActionsRequired_workflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Required_workflowsById gets an item from the github.com/octokit/kiota/.orgs.item.actions.required_workflows.item collection
func (m *ItemActionsRequestBuilder) Required_workflowsById(id string)(*ItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["required_workflow_id"] = id
    }
    return NewItemActionsRequired_workflowsWithRequired_workflow_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Runners the runners property
func (m *ItemActionsRequestBuilder) Runners()(*ItemActionsRunnersRequestBuilder) {
    return NewItemActionsRunnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RunnersById gets an item from the github.com/octokit/kiota/.orgs.item.actions.runners.item collection
func (m *ItemActionsRequestBuilder) RunnersById(id string)(*ItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["runner_id"] = id
    }
    return NewItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Secrets the secrets property
func (m *ItemActionsRequestBuilder) Secrets()(*ItemActionsSecretsRequestBuilder) {
    return NewItemActionsSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.orgs.item.actions.secrets.item collection
func (m *ItemActionsRequestBuilder) SecretsById(id string)(*ItemActionsSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemActionsSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Variables the variables property
func (m *ItemActionsRequestBuilder) Variables()(*ItemActionsVariablesRequestBuilder) {
    return NewItemActionsVariablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VariablesById gets an item from the github.com/octokit/kiota/.orgs.item.actions.variables.item collection
func (m *ItemActionsRequestBuilder) VariablesById(id string)(*ItemActionsVariablesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["name"] = id
    }
    return NewItemActionsVariablesWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
