package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MembershipsRequestBuilder builds and executes requests for operations under \user\memberships
type MembershipsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMembershipsRequestBuilderInternal instantiates a new MembershipsRequestBuilder and sets the default values.
func NewMembershipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsRequestBuilder) {
    m := &MembershipsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/memberships";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMembershipsRequestBuilder instantiates a new MembershipsRequestBuilder and sets the default values.
func NewMembershipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembershipsRequestBuilderInternal(urlParams, requestAdapter)
}
// Orgs the orgs property
func (m *MembershipsRequestBuilder) Orgs()(*MembershipsOrgsRequestBuilder) {
    return NewMembershipsOrgsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OrgsById gets an item from the github.com/octokit/kiota/.user.memberships.orgs.item collection
func (m *MembershipsRequestBuilder) OrgsById(id string)(*MembershipsOrgsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["org"] = id
    }
    return NewMembershipsOrgsWithOrgItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
