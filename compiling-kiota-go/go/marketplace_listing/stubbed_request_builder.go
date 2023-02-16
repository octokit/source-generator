package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed
type StubbedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accounts the accounts property
func (m *StubbedRequestBuilder) Accounts()(*StubbedAccountsRequestBuilder) {
    return NewStubbedAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccountsById gets an item from the github.com/octokit/kiota/.marketplace_listing.stubbed.accounts.item collection
func (m *StubbedRequestBuilder) AccountsById(id string)(*StubbedAccountsWithAccount_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account_id"] = id
    }
    return NewStubbedAccountsWithAccount_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewStubbedRequestBuilderInternal instantiates a new StubbedRequestBuilder and sets the default values.
func NewStubbedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedRequestBuilder) {
    m := &StubbedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/marketplace_listing/stubbed";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewStubbedRequestBuilder instantiates a new StubbedRequestBuilder and sets the default values.
func NewStubbedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedRequestBuilderInternal(urlParams, requestAdapter)
}
// Plans the plans property
func (m *StubbedRequestBuilder) Plans()(*StubbedPlansRequestBuilder) {
    return NewStubbedPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PlansById gets an item from the github.com/octokit/kiota/.marketplace_listing.stubbed.plans.item collection
func (m *StubbedRequestBuilder) PlansById(id string)(*StubbedPlansWithPlan_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plan_id"] = id
    }
    return NewStubbedPlansWithPlan_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
