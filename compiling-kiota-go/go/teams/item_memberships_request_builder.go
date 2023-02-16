package teams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembershipsRequestBuilder builds and executes requests for operations under \teams\{team_id}\memberships
type ItemMembershipsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemMembershipsRequestBuilderInternal instantiates a new MembershipsRequestBuilder and sets the default values.
func NewItemMembershipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembershipsRequestBuilder) {
    m := &ItemMembershipsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/memberships";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMembershipsRequestBuilder instantiates a new MembershipsRequestBuilder and sets the default values.
func NewItemMembershipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembershipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembershipsRequestBuilderInternal(urlParams, requestAdapter)
}
