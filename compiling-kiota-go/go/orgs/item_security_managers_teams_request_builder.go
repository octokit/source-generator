package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSecurityManagersTeamsRequestBuilder builds and executes requests for operations under \orgs\{org}\security-managers\teams
type ItemSecurityManagersTeamsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemSecurityManagersTeamsRequestBuilderInternal instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemSecurityManagersTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersTeamsRequestBuilder) {
    m := &ItemSecurityManagersTeamsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/security-managers/teams";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSecurityManagersTeamsRequestBuilder instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemSecurityManagersTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityManagersTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityManagersTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
