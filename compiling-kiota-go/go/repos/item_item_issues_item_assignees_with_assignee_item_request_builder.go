package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}\assignees\{assignee}
type ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    m := &ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}/assignees/{assignee}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder instantiates a new WithAssigneeItemRequestBuilder and sets the default values.
func NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks if a user has permission to be assigned to a specific issue.If the `assignee` can be assigned to this issue, a `204` status code with no content is returned.Otherwise a `404` status code is returned.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#check-if-a-user-can-be-assigned-to-a-issue
func (m *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(error) {
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
// ToGetRequestInformation checks if a user has permission to be assigned to a specific issue.If the `assignee` can be assigned to this issue, a `204` status code with no content is returned.Otherwise a `404` status code is returned.
func (m *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
