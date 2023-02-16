package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// CodespacesRequestBuilder builds and executes requests for operations under \user\codespaces
type CodespacesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CodespacesPostRequestBody composed type wrapper for classes codespacesPostRequestBodyMember1, codespacesPostRequestBodyMember2
type CodespacesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type codespacesPostRequestBodyMember1
    codespacesPostRequestBodyMember1 CodespacesPostRequestBodyMember1able
    // Composed type representation for type codespacesPostRequestBodyMember2
    codespacesPostRequestBodyMember2 CodespacesPostRequestBodyMember2able
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewCodespacesPostRequestBody instantiates a new codespacesPostRequestBody and sets the default values.
func NewCodespacesPostRequestBody()(*CodespacesPostRequestBody) {
    m := &CodespacesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCodespacesPostRequestBody()
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
func (m *CodespacesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodespacesPostRequestBodyMember1 gets the codespacesPostRequestBodyMember1 property value. Composed type representation for type codespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able) {
    return m.codespacesPostRequestBodyMember1
}
// GetCodespacesPostRequestBodyMember2 gets the codespacesPostRequestBodyMember2 property value. Composed type representation for type codespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able) {
    return m.codespacesPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// Serialize serializes information the current object
func (m *CodespacesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCodespacesPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCodespacesPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyMember2())
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
func (m *CodespacesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodespacesPostRequestBodyMember1 sets the codespacesPostRequestBodyMember1 property value. Composed type representation for type codespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)() {
    m.codespacesPostRequestBodyMember1 = value
}
// SetCodespacesPostRequestBodyMember2 sets the codespacesPostRequestBodyMember2 property value. Composed type representation for type codespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)() {
    m.codespacesPostRequestBodyMember2 = value
}
// CodespacesRequestBuilderGetQueryParameters lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
type CodespacesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // ID of the Repository to filter on
    Repository_id *int32
}
// CodespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CodespacesRequestBuilderGetQueryParameters
}
// CodespacesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesPostRequestBodyable 
type CodespacesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able)
    GetCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able)
    SetCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)()
    SetCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)()
}
// NewCodespacesRequestBuilderInternal instantiates a new CodespacesRequestBuilder and sets the default values.
func NewCodespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesRequestBuilder) {
    m := &CodespacesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/codespaces{?per_page*,page*,repository_id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCodespacesRequestBuilder instantiates a new CodespacesRequestBuilder and sets the default values.
func NewCodespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#list-codespaces-for-the-authenticated-user
func (m *CodespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesRequestBuilderGetRequestConfiguration)(CodespacesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "500": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateCodespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesResponseable), nil
}
// Post creates a new codespace, owned by the authenticated user.This endpoint requires either a `repository_id` OR a `pull_request` but not both.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#create-a-codespace-for-the-authenticated-user
func (m *CodespacesRequestBuilder) Post(ctx context.Context, body CodespacesPostRequestBodyable, requestConfiguration *CodespacesRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Codespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodespace503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Codespaceable), nil
}
// Secrets the secrets property
func (m *CodespacesRequestBuilder) Secrets()(*CodespacesSecretsRequestBuilder) {
    return NewCodespacesSecretsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretsById gets an item from the github.com/octokit/kiota/.user.codespaces.secrets.item collection
func (m *CodespacesRequestBuilder) SecretsById(id string)(*CodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secret_name"] = id
    }
    return NewCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
func (m *CodespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a new codespace, owned by the authenticated user.This endpoint requires either a `repository_id` OR a `pull_request` but not both.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
func (m *CodespacesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CodespacesPostRequestBodyable, requestConfiguration *CodespacesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
