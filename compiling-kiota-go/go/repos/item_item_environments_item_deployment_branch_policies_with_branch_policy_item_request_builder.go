package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\environments\{environment_name}\deployment-branch-policies\{branch_policy_id}
type ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal instantiates a new WithBranch_policy_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder instantiates a new WithBranch_policy_ItemRequestBuilder and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a deployment branch policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#delete-deployment-branch-policy
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get gets a deployment branch policy for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#get-deployment-branch-policy
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDeploymentBranchPolicyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable), nil
}
// Put updates a deployment branch policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/deployments/branch-policies#update-deployment-branch-policy
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) Put(ctx context.Context, body ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyNamePatternable, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderPutRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDeploymentBranchPolicyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable), nil
}
// ToDeleteRequestInformation deletes a deployment branch policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation gets a deployment branch policy for an environment.Anyone with read access to the repository can use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates a deployment branch policy for an environment.You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `administration:write` permission for the repository to use this endpoint.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyNamePatternable, requestConfiguration *ItemItemEnvironmentsItemDeploymentBranchPoliciesWithBranch_policy_ItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
