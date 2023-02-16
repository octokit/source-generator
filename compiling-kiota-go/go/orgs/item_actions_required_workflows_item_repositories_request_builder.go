package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRequired_workflowsItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\required_workflows\{required_workflow_id}\repositories
type ItemActionsRequired_workflowsItemRepositoriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRequired_workflowsItemRepositoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsItemRepositoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemActionsRequired_workflowsItemRepositoriesRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsItemRepositoriesRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRequired_workflowsItemRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) {
    m := &ItemActionsRequired_workflowsItemRepositoriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/required_workflows/{required_workflow_id}/repositories";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRequired_workflowsItemRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequired_workflowsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the selected repositories that are configured for a required workflow in an organization. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation https://docs.github.com/rest/reference/actions#list-selected-repositories-required-workflows]
// 
// [API method documentation https://docs.github.com/rest/reference/actions#list-selected-repositories-required-workflows]: https://docs.github.com/rest/reference/actions#list-selected-repositories-required-workflows
func (m *ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesRequestBuilderGetRequestConfiguration)(ItemActionsRequired_workflowsItemRepositoriesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemActionsRequired_workflowsItemRepositoriesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRequired_workflowsItemRepositoriesResponseable), nil
}
// Put sets the repositories for a required workflow that is required for selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#set-selected-repositories-for-a-required-workflow
func (m *ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) Put(ctx context.Context, body ItemActionsRequired_workflowsItemRepositoriesPutRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation lists the selected repositories that are configured for a required workflow in an organization. To use this endpoint, the required workflow must be configured to run on selected repositories.You must authenticate using an access token with the `read:org` scope to use this endpoint. GitHub Apps must have the `administration` organization permission to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation sets the repositories for a required workflow that is required for selected repositories.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsItemRepositoriesRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsRequired_workflowsItemRepositoriesPutRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsItemRepositoriesRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
