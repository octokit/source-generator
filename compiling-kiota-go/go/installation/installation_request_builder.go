package installation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationRequestBuilder builds and executes requests for operations under \installation
type InstallationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewInstallationRequestBuilderInternal instantiates a new InstallationRequestBuilder and sets the default values.
func NewInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestBuilder) {
    m := &InstallationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/installation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInstallationRequestBuilder instantiates a new InstallationRequestBuilder and sets the default values.
func NewInstallationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// Repositories the repositories property
func (m *InstallationRequestBuilder) Repositories()(*RepositoriesRequestBuilder) {
    return NewRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Token the token property
func (m *InstallationRequestBuilder) Token()(*TokenRequestBuilder) {
    return NewTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
