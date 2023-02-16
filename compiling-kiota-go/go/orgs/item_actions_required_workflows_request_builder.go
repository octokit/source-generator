package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRequired_workflowsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\required_workflows
type ItemActionsRequired_workflowsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemActionsRequired_workflowsRequestBuilderGetQueryParameters list all required workflows in an organization.You must authenticate using an access token with the `read:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
type ItemActionsRequired_workflowsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemActionsRequired_workflowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsRequired_workflowsRequestBuilderGetQueryParameters
}
// ItemActionsRequired_workflowsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRequired_workflowsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRequired_workflowsRequestBuilderInternal instantiates a new Required_workflowsRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsRequestBuilder) {
    m := &ItemActionsRequired_workflowsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/actions/required_workflows{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemActionsRequired_workflowsRequestBuilder instantiates a new Required_workflowsRequestBuilder and sets the default values.
func NewItemActionsRequired_workflowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRequired_workflowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRequired_workflowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all required workflows in an organization.You must authenticate using an access token with the `read:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-required-workflows
func (m *ItemActionsRequired_workflowsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsRequestBuilderGetRequestConfiguration)(ItemActionsRequired_workflowsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemActionsRequired_workflowsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRequired_workflowsResponseable), nil
}
// Post create a required workflow in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#create-a-required-workflow
func (m *ItemActionsRequired_workflowsRequestBuilder) Post(ctx context.Context, body ItemActionsRequired_workflowsPostRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRequiredWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable), nil
}
// ToGetRequestInformation list all required workflows in an organization.You must authenticate using an access token with the `read:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRequired_workflowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create a required workflow in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.For more information, see "[Required Workflows](https://docs.github.com/actions/using-workflows/required-workflows)."
func (m *ItemActionsRequired_workflowsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemActionsRequired_workflowsPostRequestBodyable, requestConfiguration *ItemActionsRequired_workflowsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
