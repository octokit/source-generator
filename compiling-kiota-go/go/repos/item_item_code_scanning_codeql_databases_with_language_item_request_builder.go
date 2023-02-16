package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\codeql\databases\{language}
type ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal instantiates a new WithLanguageItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/code-scanning/codeql/databases/{language}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder instantiates a new WithLanguageItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a CodeQL database for a language in a repository.By default this endpoint returns JSON metadata about the CodeQL database. Todownload the CodeQL database binary content, set the `Accept` header of the requestto [`application/zip`](https://docs.github.com/rest/overview/media-types), and make sureyour HTTP client is configured to follow redirects or use the `Location` headerto make a second request to get the redirect URL.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.GitHub Apps must have the `contents` read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/code-scanning#get-codeql-database
func (m *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningCodeqlDatabaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningCodeqlDatabase503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodeScanningCodeqlDatabaseable), nil
}
// ToGetRequestInformation gets a CodeQL database for a language in a repository.By default this endpoint returns JSON metadata about the CodeQL database. Todownload the CodeQL database binary content, set the `Accept` header of the requestto [`application/zip`](https://docs.github.com/rest/overview/media-types), and make sureyour HTTP client is configured to follow redirects or use the `Location` headerto make a second request to get the redirect URL.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.GitHub Apps must have the `contents` read permission to use this endpoint.
func (m *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
