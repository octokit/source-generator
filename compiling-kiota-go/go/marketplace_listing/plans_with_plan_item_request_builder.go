package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PlansWithPlan_ItemRequestBuilder builds and executes requests for operations under \marketplace_listing\plans\{plan_id}
type PlansWithPlan_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accounts the accounts property
func (m *PlansWithPlan_ItemRequestBuilder) Accounts()(*PlansItemAccountsRequestBuilder) {
    return NewPlansItemAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewPlansWithPlan_ItemRequestBuilderInternal instantiates a new WithPlan_ItemRequestBuilder and sets the default values.
func NewPlansWithPlan_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansWithPlan_ItemRequestBuilder) {
    m := &PlansWithPlan_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/marketplace_listing/plans/{plan_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPlansWithPlan_ItemRequestBuilder instantiates a new WithPlan_ItemRequestBuilder and sets the default values.
func NewPlansWithPlan_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlansWithPlan_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlansWithPlan_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
