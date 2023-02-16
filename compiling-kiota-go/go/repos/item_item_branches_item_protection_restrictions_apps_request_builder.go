package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\restrictions\apps
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppsDeleteRequestBody composed type wrapper for classes appsDeleteRequestBodyMember1, string
type AppsDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type appsDeleteRequestBodyMember1
    appsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewAppsDeleteRequestBody instantiates a new appsDeleteRequestBody and sets the default values.
func NewAppsDeleteRequestBody()(*AppsDeleteRequestBody) {
    m := &AppsDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsDeleteRequestBody()
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
func (m *AppsDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppsDeleteRequestBodyMember1 gets the appsDeleteRequestBodyMember1 property value. Composed type representation for type appsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) GetAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able) {
    return m.appsDeleteRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsDeleteRequestBodyMember1())
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
func (m *AppsDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppsDeleteRequestBodyMember1 sets the appsDeleteRequestBodyMember1 property value. Composed type representation for type appsDeleteRequestBodyMember1
func (m *AppsDeleteRequestBody) SetAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)() {
    m.appsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPostRequestBody composed type wrapper for classes appsPostRequestBodyMember1, string
type AppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type appsPostRequestBodyMember1
    appsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewAppsPostRequestBody instantiates a new appsPostRequestBody and sets the default values.
func NewAppsPostRequestBody()(*AppsPostRequestBody) {
    m := &AppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPostRequestBody()
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
func (m *AppsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppsPostRequestBodyMember1 gets the appsPostRequestBodyMember1 property value. Composed type representation for type appsPostRequestBodyMember1
func (m *AppsPostRequestBody) GetAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able) {
    return m.appsPostRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPostRequestBodyMember1())
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
func (m *AppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppsPostRequestBodyMember1 sets the appsPostRequestBodyMember1 property value. Composed type representation for type appsPostRequestBodyMember1
func (m *AppsPostRequestBody) SetAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)() {
    m.appsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPutRequestBody composed type wrapper for classes appsPutRequestBodyMember1, string
type AppsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type appsPutRequestBodyMember1
    appsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewAppsPutRequestBody instantiates a new appsPutRequestBody and sets the default values.
func NewAppsPutRequestBody()(*AppsPutRequestBody) {
    m := &AppsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPutRequestBody()
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
func (m *AppsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppsPutRequestBodyMember1 gets the appsPutRequestBodyMember1 property value. Composed type representation for type appsPutRequestBodyMember1
func (m *AppsPutRequestBody) GetAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able) {
    return m.appsPutRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPutRequestBodyMember1())
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
func (m *AppsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppsPutRequestBodyMember1 sets the appsPutRequestBodyMember1 property value. Composed type representation for type appsPutRequestBodyMember1
func (m *AppsPutRequestBody) SetAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)() {
    m.appsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppsDeleteRequestBodyable 
type AppsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
// AppsPostRequestBodyable 
type AppsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)
    GetString()(*string)
    SetAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)()
    SetString(value *string)()
}
// AppsPutRequestBodyable 
type AppsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)
    GetString()(*string)
    SetAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)()
    SetString(value *string)()
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder instantiates a new AppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Delete(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable)
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#list-apps-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable)
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Post(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable)
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Put(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integrationable)
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only installed GitHub Apps with `write` access to the `contents` permission can be added as authorized actors on a protected branch.
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPutRequestInformation(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
