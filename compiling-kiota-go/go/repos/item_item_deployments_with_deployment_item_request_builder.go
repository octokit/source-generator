package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemDeploymentsWithDeployment_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\deployments\{deployment_id}
type ItemItemDeploymentsWithDeployment_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemDeploymentsWithDeployment_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDeploymentsWithDeployment_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemDeploymentsWithDeployment_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDeploymentsWithDeployment_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal instantiates a new WithDeployment_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    m := &ItemItemDeploymentsWithDeployment_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/deployments/{deployment_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemDeploymentsWithDeployment_ItemRequestBuilder instantiates a new WithDeployment_ItemRequestBuilder and sets the default values.
func NewItemItemDeploymentsWithDeployment_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDeploymentsWithDeployment_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete if the repository only has one deployment, you can delete the deployment regardless of its status. If the repository has more than one deployment, you can only delete inactive deployments. This ensures that repositories with multiple deployments will always have an active deployment. Anyone with `repo` or `repo_deployment` scopes can delete a deployment.To set a deployment as inactive, you must:*   Create a new deployment that is active so that the system has a record of the current state, then delete the previously active deployment.*   Mark the active deployment as inactive by adding any non-successful deployment status.For more information, see "[Create a deployment](https://docs.github.com/rest/deployments/deployments/#create-a-deployment)" and "[Create a deployment status](https://docs.github.com/rest/deployments/deployment-statuses#create-a-deployment-status)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/deployments#delete-a-deployment
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemDeploymentsWithDeployment_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a deployment
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/deployments#get-a-deployment
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDeploymentsWithDeployment_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Deploymentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Deploymentable), nil
}
// Statuses the statuses property
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) Statuses()(*ItemItemDeploymentsItemStatusesRequestBuilder) {
    return NewItemItemDeploymentsItemStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// StatusesById gets an item from the github.com/octokit/kiota/.repos.item.item.deployments.item.statuses.item collection
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) StatusesById(id string)(*ItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["status_id"] = id
    }
    return NewItemItemDeploymentsItemStatusesWithStatus_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation if the repository only has one deployment, you can delete the deployment regardless of its status. If the repository has more than one deployment, you can only delete inactive deployments. This ensures that repositories with multiple deployments will always have an active deployment. Anyone with `repo` or `repo_deployment` scopes can delete a deployment.To set a deployment as inactive, you must:*   Create a new deployment that is active so that the system has a record of the current state, then delete the previously active deployment.*   Mark the active deployment as inactive by adding any non-successful deployment status.For more information, see "[Create a deployment](https://docs.github.com/rest/deployments/deployments/#create-a-deployment)" and "[Create a deployment status](https://docs.github.com/rest/deployments/deployment-statuses#create-a-deployment-status)."
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemDeploymentsWithDeployment_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemDeploymentsWithDeployment_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDeploymentsWithDeployment_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
