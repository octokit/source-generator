package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\required_workflows\{required_workflow_id}\repositories\{repository_id}
type ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderInternal instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) {
    m := &ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/required_workflows/{required_workflow_id}/repositories/{repository_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder instantiates a new WithRepository_ItemRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes a repository from a required workflow. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#remove-a-repository-from-selected-repositories-list-for-a-required-workflow
func (m *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Put adds a repository to a required workflow. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#add-a-repository-to-selected-repositories-list-for-a-required-workflow
func (m *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) Put(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation removes a repository from a required workflow. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation adds a repository to a required workflow. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesWithRepository_ItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
