package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemIssuesWithIssue_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}
type ItemItemIssuesWithIssue_numberItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignees the assignees property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Assignees()(*ItemItemIssuesItemAssigneesRequestBuilder) {
    return NewItemItemIssuesItemAssigneesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssigneesById gets an item from the github.com/octokit/kiota/.repos.item.item.issues.item.assignees.item collection
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) AssigneesById(id string)(*ItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignee"] = id
    }
    return NewItemItemIssuesItemAssigneesWithAssigneeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Comments the comments property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Comments()(*ItemItemIssuesItemCommentsRequestBuilder) {
    return NewItemItemIssuesItemCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal instantiates a new WithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    m := &ItemItemIssuesWithIssue_numberItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemIssuesWithIssue_numberItemRequestBuilder instantiates a new WithIssue_numberItemRequestBuilder and sets the default values.
func NewItemItemIssuesWithIssue_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesWithIssue_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Events()(*ItemItemIssuesItemEventsRequestBuilder) {
    return NewItemItemIssuesItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the API returns a [`301 Moved Permanently` status](https://docs.github.com/rest/overview/resources-in-the-rest-api#http-redirects-redirects) if the issue was[transferred](https://docs.github.com/articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/webhooks/event-payloads/#issues) webhook.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/reference/pulls#list-pull-requests)" endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#get-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Issueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Issueable), nil
}
// Labels the labels property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Labels()(*ItemItemIssuesItemLabelsRequestBuilder) {
    return NewItemItemIssuesItemLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LabelsById gets an item from the github.com/octokit/kiota/.repos.item.item.issues.item.labels.item collection
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) LabelsById(id string)(*ItemItemIssuesItemLabelsWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["name"] = id
    }
    return NewItemItemIssuesItemLabelsWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Lock the lock property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Lock()(*ItemItemIssuesItemLockRequestBuilder) {
    return NewItemItemIssuesItemLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch issue owners and users with push access can edit an issue.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#update-an-issue
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Issueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIssue503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Issueable), nil
}
// Reactions the reactions property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Reactions()(*ItemItemIssuesItemReactionsRequestBuilder) {
    return NewItemItemIssuesItemReactionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReactionsById gets an item from the github.com/octokit/kiota/.repos.item.item.issues.item.reactions.item collection
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ReactionsById(id string)(*ItemItemIssuesItemReactionsWithReaction_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reaction_id"] = id
    }
    return NewItemItemIssuesItemReactionsWithReaction_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Timeline the timeline property
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) Timeline()(*ItemItemIssuesItemTimelineRequestBuilder) {
    return NewItemItemIssuesItemTimelineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation the API returns a [`301 Moved Permanently` status](https://docs.github.com/rest/overview/resources-in-the-rest-api#http-redirects-redirects) if the issue was[transferred](https://docs.github.com/articles/transferring-an-issue-to-another-repository/) to another repository. Ifthe issue was transferred to or deleted from a repository where the authenticated user lacks read access, the APIreturns a `404 Not Found` status. If the issue was deleted from a repository where the authenticated user has readaccess, the API returns a `410 Gone` status. To receive webhook events for transferred and deleted issues, subscribeto the [`issues`](https://docs.github.com/webhooks/event-payloads/#issues) webhook.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/reference/pulls#list-pull-requests)" endpoint.
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation issue owners and users with push access can edit an issue.
func (m *ItemItemIssuesWithIssue_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemIssuesItemWithIssue_numberPatchRequestBodyable, requestConfiguration *ItemItemIssuesWithIssue_numberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
