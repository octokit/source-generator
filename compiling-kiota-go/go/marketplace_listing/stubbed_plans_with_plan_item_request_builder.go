package marketplace_listing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StubbedPlansWithPlan_ItemRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed\plans\{plan_id}
type StubbedPlansWithPlan_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accounts the accounts property
func (m *StubbedPlansWithPlan_ItemRequestBuilder) Accounts()(*StubbedPlansItemAccountsRequestBuilder) {
    return NewStubbedPlansItemAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewStubbedPlansWithPlan_ItemRequestBuilderInternal instantiates a new WithPlan_ItemRequestBuilder and sets the default values.
func NewStubbedPlansWithPlan_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansWithPlan_ItemRequestBuilder) {
    m := &StubbedPlansWithPlan_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/marketplace_listing/stubbed/plans/{plan_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewStubbedPlansWithPlan_ItemRequestBuilder instantiates a new WithPlan_ItemRequestBuilder and sets the default values.
func NewStubbedPlansWithPlan_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansWithPlan_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedPlansWithPlan_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
