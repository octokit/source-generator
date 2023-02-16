package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRequestBuilder builds and executes requests for operations under \repos\{org}\{repo}\actions
type ItemItemActionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Artifacts the artifacts property
func (m *ItemItemActionsRequestBuilder) Artifacts()(*ItemItemActionsArtifactsRequestBuilder) {
    return NewItemItemActionsArtifactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ArtifactsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.artifacts.item collection
func (m *ItemItemActionsRequestBuilder) ArtifactsById(id string)(*ItemItemActionsArtifactsWithArtifact_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["artifact_id"] = id
    }
    return NewItemItemActionsArtifactsWithArtifact_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Cache the cache property
func (m *ItemItemActionsRequestBuilder) Cache()(*ItemItemActionsCacheRequestBuilder) {
    return NewItemItemActionsCacheRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Caches the caches property
func (m *ItemItemActionsRequestBuilder) Caches()(*ItemItemActionsCachesRequestBuilder) {
    return NewItemItemActionsCachesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CachesById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.caches.item collection
func (m *ItemItemActionsRequestBuilder) CachesById(id string)(*ItemItemActionsCachesWithCache_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cache_id"] = id
    }
    return NewItemItemActionsCachesWithCache_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemActionsRequestBuilderInternal instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    m := &ItemItemActionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{org}/{repo}/actions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRequestBuilder instantiates a new ActionsRequestBuilder and sets the default values.
func NewItemItemActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Jobs the jobs property
func (m *ItemItemActionsRequestBuilder) Jobs()(*ItemItemActionsJobsRequestBuilder) {
    return NewItemItemActionsJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// JobsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.jobs.item collection
func (m *ItemItemActionsRequestBuilder) JobsById(id string)(*ItemItemActionsJobsWithJob_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["job_id"] = id
    }
    return NewItemItemActionsJobsWithJob_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Oidc the oidc property
func (m *ItemItemActionsRequestBuilder) Oidc()(*ItemItemActionsOidcRequestBuilder) {
    return NewItemItemActionsOidcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Permissions the permissions property
func (m *ItemItemActionsRequestBuilder) Permissions()(*ItemItemActionsPermissionsRequestBuilder) {
    return NewItemItemActionsPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Required_workflows the required_workflows property
func (m *ItemItemActionsRequestBuilder) Required_workflows()(*ItemItemActionsRequired_workflowsRequestBuilder) {
    return NewItemItemActionsRequired_workflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Runners the runners property
func (m *ItemItemActionsRequestBuilder) Runners()(*ItemItemActionsRunnersRequestBuilder) {
    return NewItemItemActionsRunnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RunnersById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.runners.item collection
func (m *ItemItemActionsRequestBuilder) RunnersById(id string)(*ItemItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["runner_id"] = id
    }
    return NewItemItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Runs the runs property
func (m *ItemItemActionsRequestBuilder) Runs()(*ItemItemActionsRunsRequestBuilder) {
    return NewItemItemActionsRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RunsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.runs.item collection
func (m *ItemItemActionsRequestBuilder) RunsById(id string)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run_id"] = id
    }
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Secrets the secrets property
func (m *ItemItemActionsRequestBuilder) Secrets()(*ItemItemActionsSecretsRequestBuilder) {
    return NewItemItemActionsSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.secrets.item collection
func (m *ItemItemActionsRequestBuilder) SecretsById(id string)(*ItemItemActionsSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewItemItemActionsSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Variables the variables property
func (m *ItemItemActionsRequestBuilder) Variables()(*ItemItemActionsVariablesRequestBuilder) {
    return NewItemItemActionsVariablesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VariablesById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.variables.item collection
func (m *ItemItemActionsRequestBuilder) VariablesById(id string)(*ItemItemActionsVariablesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["name"] = id
    }
    return NewItemItemActionsVariablesWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Workflows the workflows property
func (m *ItemItemActionsRequestBuilder) Workflows()(*ItemItemActionsWorkflowsRequestBuilder) {
    return NewItemItemActionsWorkflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WorkflowsById gets an item from the github.com/octokit/kiota/.repos.item.item.actions.workflows.item collection
func (m *ItemItemActionsRequestBuilder) WorkflowsById(id string)(*ItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflow_id"] = id
    }
    return NewItemItemActionsWorkflowsWithWorkflow_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
