package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\restrictions\teams
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsDeleteRequestBody composed type wrapper for classes teamsDeleteRequestBodyMember1, string
type TeamsDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type teamsDeleteRequestBodyMember1
    teamsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
}
// NewTeamsDeleteRequestBody instantiates a new teamsDeleteRequestBody and sets the default values.
func NewTeamsDeleteRequestBody()(*TeamsDeleteRequestBody) {
    m := &TeamsDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsDeleteRequestBody()
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
func (m *TeamsDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *TeamsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsDeleteRequestBodyMember1 gets the teamsDeleteRequestBodyMember1 property value. Composed type representation for type teamsDeleteRequestBodyMember1
func (m *TeamsDeleteRequestBody) GetTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able) {
    return m.teamsDeleteRequestBodyMember1
}
// Serialize serializes information the current object
func (m *TeamsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTeamsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsDeleteRequestBodyMember1())
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
func (m *TeamsDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsDeleteRequestBodyMember1 sets the teamsDeleteRequestBodyMember1 property value. Composed type representation for type teamsDeleteRequestBodyMember1
func (m *TeamsDeleteRequestBody) SetTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)() {
    m.teamsDeleteRequestBodyMember1 = value
}
// TeamsPostRequestBody composed type wrapper for classes teamsPostRequestBodyMember1, string
type TeamsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type teamsPostRequestBodyMember1
    teamsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
}
// NewTeamsPostRequestBody instantiates a new teamsPostRequestBody and sets the default values.
func NewTeamsPostRequestBody()(*TeamsPostRequestBody) {
    m := &TeamsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsPostRequestBody()
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
func (m *TeamsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *TeamsPostRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsPostRequestBodyMember1 gets the teamsPostRequestBodyMember1 property value. Composed type representation for type teamsPostRequestBodyMember1
func (m *TeamsPostRequestBody) GetTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able) {
    return m.teamsPostRequestBodyMember1
}
// Serialize serializes information the current object
func (m *TeamsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTeamsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsPostRequestBodyMember1())
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
func (m *TeamsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsPostRequestBodyMember1 sets the teamsPostRequestBodyMember1 property value. Composed type representation for type teamsPostRequestBodyMember1
func (m *TeamsPostRequestBody) SetTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)() {
    m.teamsPostRequestBodyMember1 = value
}
// TeamsPutRequestBody composed type wrapper for classes teamsPutRequestBodyMember1, string
type TeamsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
    // Composed type representation for type teamsPutRequestBodyMember1
    teamsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
}
// NewTeamsPutRequestBody instantiates a new teamsPutRequestBody and sets the default values.
func NewTeamsPutRequestBody()(*TeamsPutRequestBody) {
    m := &TeamsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsPutRequestBody()
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
func (m *TeamsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *TeamsPutRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsPutRequestBodyMember1 gets the teamsPutRequestBodyMember1 property value. Composed type representation for type teamsPutRequestBodyMember1
func (m *TeamsPutRequestBody) GetTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able) {
    return m.teamsPutRequestBodyMember1
}
// Serialize serializes information the current object
func (m *TeamsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTeamsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsPutRequestBodyMember1())
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
func (m *TeamsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsPutRequestBodyMember1 sets the teamsPutRequestBodyMember1 property value. Composed type representation for type teamsPutRequestBodyMember1
func (m *TeamsPutRequestBody) SetTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)() {
    m.teamsPutRequestBodyMember1 = value
}
// TeamsDeleteRequestBodyable 
type TeamsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)
    SetString(value *string)()
    SetTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)()
}
// TeamsPostRequestBodyable 
type TeamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)
    SetString(value *string)()
    SetTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)()
}
// TeamsPutRequestBodyable 
type TeamsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetString()(*string)
    GetTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)
    SetString(value *string)()
    SetTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)()
}
// NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a team to push to this branch. You can also remove push access for child teams.| Type    | Description                                                                                                                                         || ------- | --------------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Teams that should no longer have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Delete(ctx context.Context, body TeamsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderDeleteRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable)
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the teams who have push access to this branch. The list includes child teams.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#list-teams-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable)
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified teams push access for this branch. You can also give push access to child teams.| Type    | Description                                                                                                                                || ------- | ------------------------------------------------------------------------------------------------------------------------------------------ || `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Post(ctx context.Context, body TeamsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPostRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable)
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of teams that have push access to this branch. This removes all teams that previously had push access and grants push access to the new list of teams. Team restrictions include child teams.| Type    | Description                                                                                                                                || ------- | ------------------------------------------------------------------------------------------------------------------------------------------ || `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Put(ctx context.Context, body TeamsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPutRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Teamable)
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a team to push to this branch. You can also remove push access for child teams.| Type    | Description                                                                                                                                         || ------- | --------------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Teams that should no longer have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body TeamsDeleteRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the teams who have push access to this branch. The list includes child teams.
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified teams push access for this branch. You can also give push access to child teams.| Type    | Description                                                                                                                                || ------- | ------------------------------------------------------------------------------------------------------------------------------------------ || `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamsPostRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of teams that have push access to this branch. This removes all teams that previously had push access and grants push access to the new list of teams. Team restrictions include child teams.| Type    | Description                                                                                                                                || ------- | ------------------------------------------------------------------------------------------------------------------------------------------ || `array` | The teams that can have push access. Use the team's `slug`. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToPutRequestInformation(ctx context.Context, body TeamsPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
