package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsItemCheckSuitesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\commits\{ref}\check-suites
type ItemItemCommitsItemCheckSuitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.
type ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters struct {
    // Filters check suites by GitHub App `id`.
    App_id *int32
    // Returns check runs with the specified `name`.
    Check_name *string
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters
}
// NewItemItemCommitsItemCheckSuitesRequestBuilderInternal instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    m := &ItemItemCommitsItemCheckSuitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/commits/{ref}/check-suites{?app_id*,check_name*,per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCommitsItemCheckSuitesRequestBuilder instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#list-check-suites-for-a-git-reference
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration)(ItemItemCommitsItemCheckSuitesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckSuitesResponseable), nil
}
// ToGetRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.Lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth Apps and authenticated users must have the `repo` scope to get check suites in a private repository.
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
