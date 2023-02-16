package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\codeql
type ItemItemCodeScanningCodeqlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemCodeScanningCodeqlRequestBuilderInternal instantiates a new CodeqlRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/code-scanning/codeql";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCodeScanningCodeqlRequestBuilder instantiates a new CodeqlRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlRequestBuilderInternal(urlParams, requestAdapter)
}
// Databases the databases property
func (m *ItemItemCodeScanningCodeqlRequestBuilder) Databases()(*ItemItemCodeScanningCodeqlDatabasesRequestBuilder) {
    return NewItemItemCodeScanningCodeqlDatabasesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DatabasesById gets an item from the github.com/octokit/kiota/.repos.item.item.codeScanning.codeql.databases.item collection
func (m *ItemItemCodeScanningCodeqlRequestBuilder) DatabasesById(id string)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["language"] = id
    }
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
