package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationsWithInstallation_ItemRequestBuilder builds and executes requests for operations under \user\installations\{installation_id}
type InstallationsWithInstallation_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewInstallationsWithInstallation_ItemRequestBuilderInternal instantiates a new WithInstallation_ItemRequestBuilder and sets the default values.
func NewInstallationsWithInstallation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsWithInstallation_ItemRequestBuilder) {
    m := &InstallationsWithInstallation_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/installations/{installation_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInstallationsWithInstallation_ItemRequestBuilder instantiates a new WithInstallation_ItemRequestBuilder and sets the default values.
func NewInstallationsWithInstallation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationsWithInstallation_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationsWithInstallation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Repositories the repositories property
func (m *InstallationsWithInstallation_ItemRequestBuilder) Repositories()(*InstallationsItemRepositoriesRequestBuilder) {
    return NewInstallationsItemRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RepositoriesById gets an item from the github.com/octokit/kiota/.user.installations.item.repositories.item collection
func (m *InstallationsWithInstallation_ItemRequestBuilder) RepositoriesById(id string)(*InstallationsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["repository_id"] = id
    }
    return NewInstallationsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
