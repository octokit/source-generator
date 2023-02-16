package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedAccountsRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed\accounts
type StubbedAccountsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewStubbedAccountsRequestBuilderInternal instantiates a new AccountsRequestBuilder and sets the default values.
func NewStubbedAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedAccountsRequestBuilder) {
    m := &StubbedAccountsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/marketplace_listing/stubbed/accounts";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewStubbedAccountsRequestBuilder instantiates a new AccountsRequestBuilder and sets the default values.
func NewStubbedAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
