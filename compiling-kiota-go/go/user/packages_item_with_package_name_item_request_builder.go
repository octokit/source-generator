package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// PackagesItemWithPackage_nameItemRequestBuilder builds and executes requests for operations under \user\packages\{package_type}\{package_name}
type PackagesItemWithPackage_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PackagesItemWithPackage_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PackagesItemWithPackage_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PackagesItemWithPackage_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PackagesItemWithPackage_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPackagesItemWithPackage_nameItemRequestBuilderInternal instantiates a new WithPackage_nameItemRequestBuilder and sets the default values.
func NewPackagesItemWithPackage_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesItemWithPackage_nameItemRequestBuilder) {
    m := &PackagesItemWithPackage_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/packages/{package_type}/{package_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPackagesItemWithPackage_nameItemRequestBuilder instantiates a new WithPackage_nameItemRequestBuilder and sets the default values.
func NewPackagesItemWithPackage_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesItemWithPackage_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesItemWithPackage_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a package owned by the authenticated user. You cannot delete a public package if any version of the package has more than 5,000 downloads. In this scenario, contact GitHub support for further assistance.To use this endpoint, you must authenticate using an access token with the `read:packages` and `delete:packages` scopes.If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/packages#delete-a-package-for-the-authenticated-user
func (m *PackagesItemWithPackage_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PackagesItemWithPackage_nameItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get gets a specific package for a package owned by the authenticated user.To use this endpoint, you must authenticate using an access token with the `read:packages` scope. If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/packages#get-a-package-for-the-authenticated-user
func (m *PackagesItemWithPackage_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PackagesItemWithPackage_nameItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Packageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreatePackageFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Packageable), nil
}
// Restore the restore property
func (m *PackagesItemWithPackage_nameItemRequestBuilder) Restore()(*PackagesItemItemRestoreRequestBuilder) {
    return NewPackagesItemItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation deletes a package owned by the authenticated user. You cannot delete a public package if any version of the package has more than 5,000 downloads. In this scenario, contact GitHub support for further assistance.To use this endpoint, you must authenticate using an access token with the `read:packages` and `delete:packages` scopes.If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
func (m *PackagesItemWithPackage_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PackagesItemWithPackage_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets a specific package for a package owned by the authenticated user.To use this endpoint, you must authenticate using an access token with the `read:packages` scope. If the `package_type` belongs to a GitHub Packages registry that only supports repository-scoped permissions, your token must also include the `repo` scope. For the list of GitHub Packages registries that only support repository-scoped permissions, see "[About permissions for GitHub Packages](https://docs.github.com/packages/learn-github-packages/about-permissions-for-github-packages#permissions-for-repository-scoped-packages)."
func (m *PackagesItemWithPackage_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PackagesItemWithPackage_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Versions the versions property
func (m *PackagesItemWithPackage_nameItemRequestBuilder) Versions()(*PackagesItemItemVersionsRequestBuilder) {
    return NewPackagesItemItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VersionsById gets an item from the github.com/octokit/kiota/.user.packages.item.item.versions.item collection
func (m *PackagesItemWithPackage_nameItemRequestBuilder) VersionsById(id string)(*PackagesItemItemVersionsWithPackage_version_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["package_version_id"] = id
    }
    return NewPackagesItemItemVersionsWithPackage_version_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
