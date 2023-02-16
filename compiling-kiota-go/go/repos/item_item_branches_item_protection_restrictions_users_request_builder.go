package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\restrictions\users
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersDeleteRequestBody composed type wrapper for classes usersDeleteRequestBodyMember1, string
type UsersDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type usersDeleteRequestBodyMember1
    usersDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
}
// NewUsersDeleteRequestBody instantiates a new usersDeleteRequestBody and sets the default values.
func NewUsersDeleteRequestBody()(*UsersDeleteRequestBody) {
    m := &UsersDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersDeleteRequestBody()
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
func (m *UsersDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *UsersDeleteRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersDeleteRequestBodyMember1 gets the usersDeleteRequestBodyMember1 property value. Composed type representation for type usersDeleteRequestBodyMember1
func (m *UsersDeleteRequestBody) GetUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able) {
    return m.usersDeleteRequestBodyMember1
}
// Serialize serializes information the current object
func (m *UsersDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsersDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersDeleteRequestBodyMember1())
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
func (m *UsersDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersDeleteRequestBodyMember1 sets the usersDeleteRequestBodyMember1 property value. Composed type representation for type usersDeleteRequestBodyMember1
func (m *UsersDeleteRequestBody) SetUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)() {
    m.usersDeleteRequestBodyMember1 = value
}
// UsersPostRequestBody composed type wrapper for classes usersPostRequestBodyMember1, string
type UsersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type usersPostRequestBodyMember1
    usersPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
}
// NewUsersPostRequestBody instantiates a new usersPostRequestBody and sets the default values.
func NewUsersPostRequestBody()(*UsersPostRequestBody) {
    m := &UsersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersPostRequestBody()
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
func (m *UsersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *UsersPostRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersPostRequestBodyMember1 gets the usersPostRequestBodyMember1 property value. Composed type representation for type usersPostRequestBodyMember1
func (m *UsersPostRequestBody) GetUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able) {
    return m.usersPostRequestBodyMember1
}
// Serialize serializes information the current object
func (m *UsersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsersPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersPostRequestBodyMember1())
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
func (m *UsersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersPostRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersPostRequestBodyMember1 sets the usersPostRequestBodyMember1 property value. Composed type representation for type usersPostRequestBodyMember1
func (m *UsersPostRequestBody) SetUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)() {
    m.usersPostRequestBodyMember1 = value
}
// UsersPutRequestBody composed type wrapper for classes usersPutRequestBodyMember1, string
type UsersPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type usersPutRequestBodyMember1
    usersPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
}
// NewUsersPutRequestBody instantiates a new usersPutRequestBody and sets the default values.
func NewUsersPutRequestBody()(*UsersPutRequestBody) {
    m := &UsersPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUsersPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersPutRequestBody()
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
func (m *UsersPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *UsersPutRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersPutRequestBodyMember1 gets the usersPutRequestBodyMember1 property value. Composed type representation for type usersPutRequestBodyMember1
func (m *UsersPutRequestBody) GetUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able) {
    return m.usersPutRequestBodyMember1
}
// Serialize serializes information the current object
func (m *UsersPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUsersPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersPutRequestBodyMember1())
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
func (m *UsersPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersPutRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersPutRequestBodyMember1 sets the usersPutRequestBodyMember1 property value. Composed type representation for type usersPutRequestBodyMember1
func (m *UsersPutRequestBody) SetUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)() {
    m.usersPutRequestBodyMember1 = value
}
// UsersDeleteRequestBodyable 
type UsersDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)
    SetString(value *string)()
    SetUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)()
}
// UsersPostRequestBodyable 
type UsersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)
    SetString(value *string)()
    SetUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)()
}
// UsersPutRequestBodyable 
type UsersPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)
    SetString(value *string)()
    SetUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)()
}
// NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a user to push to this branch.| Type    | Description                                                                                                                                   || ------- | --------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames of the people who should no longer have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Delete(ctx context.Context, body UsersDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderDeleteRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable)
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the people who have push access to this branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#list-users-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable)
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified people push access for this branch.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Post(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPostRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable)
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of people that have push access to this branch. This removes all people that previously had push access and grants push access to the new list of people.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Put(ctx context.Context, body UsersPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPutRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable)
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a user to push to this branch.| Type    | Description                                                                                                                                   || ------- | --------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames of the people who should no longer have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body UsersDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the people who have push access to this branch.
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified people push access for this branch.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of people that have push access to this branch. This removes all people that previously had push access and grants push access to the new list of people.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToPutRequestInformation(ctx context.Context, body UsersPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
