package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    m := &ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder instantiates a new WithEnvironment_nameItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete you must authenticate using an access token with the repo scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#delete-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentBranchPolicies the deploymentBranchPolicies property
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) DeploymentBranchPolicies()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilder) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeploymentBranchPoliciesById gets an item from the github.com/octokit/kiota/.repos.item.item.environments.item.deploymentBranchPolicies.item collection
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) DeploymentBranchPoliciesById(id string)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["branch_policy_id"] = id
    }
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get **Note:** To get information about name patterns that branches must match in order to deploy to this environment, see "[Get a deployment branch policy](/rest/deployments/branch-policies#get-a-deployment-branch-policy)."Anyone with read access to the repository can use this endpoint. If therepository is private, you must use an access token with the `repo` scope. GitHubApps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#get-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Environmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEnvironmentFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Environmentable), nil
}
// Put create or update an environment with protection rules, such as required reviewers. For more information about environment protection rules, see "[Environments](/actions/reference/environments#environment-protection-rules)."**Note:** To create or update name patterns that branches must match in order to deploy to this environment, see "[Deployment branch policies](/rest/deployments/branch-policies)."**Note:** To create or update secrets for an environment, see "[Secrets](/rest/reference/actions#secrets)."You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/environments#create-or-update-an-environment
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) Put(ctx context.Context, body ItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyable, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Environmentable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Environmentable), nil
}
// ToDeleteRequestInformation you must authenticate using an access token with the repo scope to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation **Note:** To get information about name patterns that branches must match in order to deploy to this environment, see "[Get a deployment branch policy](/rest/deployments/branch-policies#get-a-deployment-branch-policy)."Anyone with read access to the repository can use this endpoint. If therepository is private, you must use an access token with the `repo` scope. GitHubApps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation create or update an environment with protection rules, such as required reviewers. For more information about environment protection rules, see "[Environments](/actions/reference/environments#environment-protection-rules)."**Note:** To create or update name patterns that branches must match in order to deploy to this environment, see "[Deployment branch policies](/rest/deployments/branch-policies)."**Note:** To create or update secrets for an environment, see "[Secrets](/rest/reference/actions#secrets)."You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
func (m *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyable, requestConfiguration *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
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
