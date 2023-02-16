package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemAssigneesWithAssigneeItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\assignees\{assignee}
type ItemItemAssigneesWithAssigneeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    m := &ItemItemAssigneesWithAssigneeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/assignees/{assignee}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemAssigneesWithAssigneeItemRequestBuilder instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemAssigneesWithAssigneeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAssigneesWithAssigneeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#check-if-a-user-can-be-assigned
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation checks if a user has permission to be assigned to an issue in this repository.If the `assignee` can be assigned to issues in the repository, a `204` header with no content is returned.Otherwise a `404` status code is returned.
func (m *ItemItemAssigneesWithAssigneeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
