package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Marketplace_listingRequestBuilder builds and executes requests for operations under \marketplace_listing
type Marketplace_listingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accounts the accounts property
func (m *Marketplace_listingRequestBuilder) Accounts()(*AccountsRequestBuilder) {
    return NewAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccountsById gets an item from the github.com/octokit/kiota/.marketplace_listing.accounts.item collection
func (m *Marketplace_listingRequestBuilder) AccountsById(id string)(*AccountsWithAccount_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account_id"] = id
    }
    return NewAccountsWithAccount_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMarketplace_listingRequestBuilderInternal instantiates a new Marketplace_listingRequestBuilder and sets the default values.
func NewMarketplace_listingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_listingRequestBuilder) {
    m := &Marketplace_listingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/marketplace_listing";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMarketplace_listingRequestBuilder instantiates a new Marketplace_listingRequestBuilder and sets the default values.
func NewMarketplace_listingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_listingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMarketplace_listingRequestBuilderInternal(urlParams, requestAdapter)
}
// Plans the plans property
func (m *Marketplace_listingRequestBuilder) Plans()(*PlansRequestBuilder) {
    return NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PlansById gets an item from the github.com/octokit/kiota/.marketplace_listing.plans.item collection
func (m *Marketplace_listingRequestBuilder) PlansById(id string)(*PlansWithPlan_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plan_id"] = id
    }
    return NewPlansWithPlan_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Stubbed the stubbed property
func (m *Marketplace_listingRequestBuilder) Stubbed()(*StubbedRequestBuilder) {
    return NewStubbedRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
