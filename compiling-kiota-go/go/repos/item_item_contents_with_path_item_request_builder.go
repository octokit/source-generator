package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemContentsWithPathItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\contents\{path}
type ItemItemContentsWithPathItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemContentsWithPathItemRequestBuilderGetQueryParameters gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/reference/repos#custom-media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/reference/repos#custom-media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/reference/git#trees).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/reference/git#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.#### Size limitsIf the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported.#### If the content is a directoryThe response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule".#### If the content is a symlink If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself.#### If the content is a submoduleThe `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
type ItemItemContentsWithPathItemRequestBuilderGetQueryParameters struct {
    // The name of the commit/branch/tag. Default: the repository’s default branch (usually `master`)
    Ref *string
}
// ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemContentsWithPathItemRequestBuilderGetQueryParameters
}
// ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithPathResponse composed type wrapper for classes WithPathResponseMember1, contentFile, contentSymlink, contentSubmodule
type WithPathResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type contentFile
    contentFile ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentFileable
    // Composed type representation for type contentSubmodule
    contentSubmodule ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSubmoduleable
    // Composed type representation for type contentSymlink
    contentSymlink ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSymlinkable
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type WithPathResponseMember1
    withPathResponseMember1 ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WithPathResponseMember1able
}
// NewWithPathResponse instantiates a new WithPathResponse and sets the default values.
func NewWithPathResponse()(*WithPathResponse) {
    m := &WithPathResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWithPathResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithPathResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithPathResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WithPathResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentFile gets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathResponse) GetContentFile()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentFileable) {
    return m.contentFile
}
// GetContentSubmodule gets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathResponse) GetContentSubmodule()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSubmoduleable) {
    return m.contentSubmodule
}
// GetContentSymlink gets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathResponse) GetContentSymlink()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSymlinkable) {
    return m.contentSymlink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WithPathResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetWithPathResponseMember1 gets the withPathResponseMember1 property value. Composed type representation for type WithPathResponseMember1
func (m *WithPathResponse) GetWithPathResponseMember1()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WithPathResponseMember1able) {
    return m.withPathResponseMember1
}
// Serialize serializes information the current object
func (m *WithPathResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentFile() != nil {
        err := writer.WriteObjectValue("", m.GetContentFile())
        if err != nil {
            return err
        }
    } else if m.GetContentSubmodule() != nil {
        err := writer.WriteObjectValue("", m.GetContentSubmodule())
        if err != nil {
            return err
        }
    } else if m.GetContentSymlink() != nil {
        err := writer.WriteObjectValue("", m.GetContentSymlink())
        if err != nil {
            return err
        }
    } else if m.GetWithPathResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetWithPathResponseMember1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WithPathResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentFile sets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathResponse) SetContentFile(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentFileable)() {
    m.contentFile = value
}
// SetContentSubmodule sets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathResponse) SetContentSubmodule(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSubmoduleable)() {
    m.contentSubmodule = value
}
// SetContentSymlink sets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathResponse) SetContentSymlink(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSymlinkable)() {
    m.contentSymlink = value
}
// SetWithPathResponseMember1 sets the withPathResponseMember1 property value. Composed type representation for type WithPathResponseMember1
func (m *WithPathResponse) SetWithPathResponseMember1(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WithPathResponseMember1able)() {
    m.withPathResponseMember1 = value
}
// WithPathResponseable 
type WithPathResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFile()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentFileable)
    GetContentSubmodule()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSubmoduleable)
    GetContentSymlink()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSymlinkable)
    GetWithPathResponseMember1()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WithPathResponseMember1able)
    SetContentFile(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentFileable)()
    SetContentSubmodule(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSubmoduleable)()
    SetContentSymlink(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ContentSymlinkable)()
    SetWithPathResponseMember1(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.WithPathResponseMember1able)()
}
// NewItemItemContentsWithPathItemRequestBuilderInternal instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    m := &ItemItemContentsWithPathItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/contents/{path}{?ref*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemContentsWithPathItemRequestBuilder instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemContentsWithPathItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/reference/repos/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/repos#delete-a-file
func (m *ItemItemContentsWithPathItemRequestBuilder) Delete(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FileCommitable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "409": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateFileCommit503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FileCommitable), nil
}
// Get gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/reference/repos#custom-media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/reference/repos#custom-media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/reference/git#trees).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/reference/git#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.#### Size limitsIf the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported.#### If the content is a directoryThe response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule".#### If the content is a symlink If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself.#### If the content is a submoduleThe `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/repos#get-repository-content
func (m *ItemItemContentsWithPathItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration)(WithPathResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateWithPathResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithPathResponseable), nil
}
// Put creates a new file or replaces an existing file in a repository. You must authenticate using an access token with the `workflow` scope to use this endpoint.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/reference/repos/#delete-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/repos#create-or-update-file-contents
func (m *ItemItemContentsWithPathItemRequestBuilder) Put(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FileCommitable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "409": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FileCommitable), nil
}
// ToDeleteRequestInformation deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/reference/repos/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/reference/repos#custom-media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/reference/repos#custom-media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/reference/git#trees).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/reference/git#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.#### Size limitsIf the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported.#### If the content is a directoryThe response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule".#### If the content is a symlink If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself.#### If the content is a submoduleThe `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation creates a new file or replaces an existing file in a repository. You must authenticate using an access token with the `workflow` scope to use this endpoint.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/reference/repos/#delete-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
