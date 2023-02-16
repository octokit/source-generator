package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsBillingRequestBuilder builds and executes requests for operations under \users\{username}\settings\billing
type ItemSettingsBillingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Actions the actions property
func (m *ItemSettingsBillingRequestBuilder) Actions()(*ItemSettingsBillingActionsRequestBuilder) {
    return NewItemSettingsBillingActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemSettingsBillingRequestBuilderInternal instantiates a new BillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    m := &ItemSettingsBillingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}/settings/billing";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSettingsBillingRequestBuilder instantiates a new BillingRequestBuilder and sets the default values.
func NewItemSettingsBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Packages the packages property
func (m *ItemSettingsBillingRequestBuilder) Packages()(*ItemSettingsBillingPackagesRequestBuilder) {
    return NewItemSettingsBillingPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharedStorage the sharedStorage property
func (m *ItemSettingsBillingRequestBuilder) SharedStorage()(*ItemSettingsBillingSharedStorageRequestBuilder) {
    return NewItemSettingsBillingSharedStorageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
