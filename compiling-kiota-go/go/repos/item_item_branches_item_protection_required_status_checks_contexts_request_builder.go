package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\required_status_checks\contexts
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContextsDeleteRequestBody composed type wrapper for classes contextsDeleteRequestBodyMember1, string
type ContextsDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type contextsDeleteRequestBodyMember1
    contextsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewContextsDeleteRequestBody instantiates a new contextsDeleteRequestBody and sets the default values.
func NewContextsDeleteRequestBody()(*ContextsDeleteRequestBody) {
    m := &ContextsDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContextsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContextsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsDeleteRequestBody()
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
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContextsDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContextsDeleteRequestBodyMember1 gets the contextsDeleteRequestBodyMember1 property value. Composed type representation for type contextsDeleteRequestBodyMember1
func (m *ContextsDeleteRequestBody) GetContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able) {
    return m.contextsDeleteRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContextsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *ContextsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *ContextsDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContextsDeleteRequestBodyMember1 sets the contextsDeleteRequestBodyMember1 property value. Composed type representation for type contextsDeleteRequestBodyMember1
func (m *ContextsDeleteRequestBody) SetContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)() {
    m.contextsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// ContextsPostRequestBody composed type wrapper for classes contextsPostRequestBodyMember1, string
type ContextsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type contextsPostRequestBodyMember1
    contextsPostRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewContextsPostRequestBody instantiates a new contextsPostRequestBody and sets the default values.
func NewContextsPostRequestBody()(*ContextsPostRequestBody) {
    m := &ContextsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContextsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContextsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsPostRequestBody()
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
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContextsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContextsPostRequestBodyMember1 gets the contextsPostRequestBodyMember1 property value. Composed type representation for type contextsPostRequestBodyMember1
func (m *ContextsPostRequestBody) GetContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able) {
    return m.contextsPostRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContextsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *ContextsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *ContextsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContextsPostRequestBodyMember1 sets the contextsPostRequestBodyMember1 property value. Composed type representation for type contextsPostRequestBodyMember1
func (m *ContextsPostRequestBody) SetContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)() {
    m.contextsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// ContextsPutRequestBody composed type wrapper for classes contextsPutRequestBodyMember1, string
type ContextsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type contextsPutRequestBodyMember1
    contextsPutRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewContextsPutRequestBody instantiates a new contextsPutRequestBody and sets the default values.
func NewContextsPutRequestBody()(*ContextsPutRequestBody) {
    m := &ContextsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContextsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContextsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsPutRequestBody()
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
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContextsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContextsPutRequestBodyMember1 gets the contextsPutRequestBodyMember1 property value. Composed type representation for type contextsPutRequestBodyMember1
func (m *ContextsPutRequestBody) GetContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able) {
    return m.contextsPutRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContextsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *ContextsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *ContextsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContextsPutRequestBodyMember1 sets the contextsPutRequestBodyMember1 property value. Composed type representation for type contextsPutRequestBodyMember1
func (m *ContextsPutRequestBody) SetContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)() {
    m.contextsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContextsDeleteRequestBodyable 
type ContextsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
// ContextsPostRequestBodyable 
type ContextsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)
    GetString()(*string)
    SetContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)()
    SetString(value *string)()
}
// ContextsPutRequestBodyable 
type ContextsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)
    GetString()(*string)
    SetContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)()
    SetString(value *string)()
}
// NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal instantiates a new ContextsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder instantiates a new ContextsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Delete(ctx context.Context, body ContextsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderDeleteRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-all-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderGetRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Post(ctx context.Context, body ContextsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPostRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Put(ctx context.Context, body ContextsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPutRequestConfiguration)([]string, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ContextsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContextsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ContextsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
