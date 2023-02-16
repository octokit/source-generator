package repositories

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemEnvironmentsItemVariablesWithNameItemRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\variables\{name}
type ItemEnvironmentsItemVariablesWithNameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemEnvironmentsItemVariablesWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesWithNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemEnvironmentsItemVariablesWithNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesWithNameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemEnvironmentsItemVariablesWithNameItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesWithNameItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEnvironmentsItemVariablesWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) {
    m := &ItemEnvironmentsItemVariablesWithNameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/variables/{name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEnvironmentsItemVariablesWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemVariablesWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes an environment variable using the variable name.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#delete-an-environment-variable
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get gets a specific variable in an environment. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `environments:read` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#get-an-environment-variable
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateActionsVariableFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable), nil
}
// Patch updates an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#update-an-environment-variable
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) Patch(ctx context.Context, body ItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation deletes an environment variable using the variable name.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation gets a specific variable in an environment. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `environments:read` repository permission to use this endpoint.
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation updates an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
func (m *ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesWithNameItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
