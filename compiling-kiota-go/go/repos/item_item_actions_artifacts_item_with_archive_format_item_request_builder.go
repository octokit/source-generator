package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\artifacts\{artifact_id}\{archive_format}
type ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal instantiates a new WithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    m := &ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/artifacts/{artifact_id}/{archive_format}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder instantiates a new WithArchive_formatItemRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#download-an-artifact
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download an archive for a repository. This URL expires after 1 minute. Look for `Location:` inthe response header to find the URL for the download. The `:archive_format` must be `zip`. Anyone with read access tothe repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsItemWithArchive_formatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
