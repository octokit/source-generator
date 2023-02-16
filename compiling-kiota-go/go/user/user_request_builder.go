package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// UserRequestBuilder builds and executes requests for operations under \user
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserResponse composed type wrapper for classes privateUser, publicUser
type UserResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type privateUser
    privateUser ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable
    // Composed type representation for type publicUser
    publicUser ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewUserResponse instantiates a new userResponse and sets the default values.
func NewUserResponse()(*UserResponse) {
    m := &UserResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserResponse()
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
func (m *UserResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *UserResponse) GetPrivateUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *UserResponse) GetPublicUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *UserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPrivateUser() != nil {
        err := writer.WriteObjectValue("", m.GetPrivateUser())
        if err != nil {
            return err
        }
    } else if m.GetPublicUser() != nil {
        err := writer.WriteObjectValue("", m.GetPublicUser())
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
func (m *UserResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateUser sets the privateUser property value. Composed type representation for type privateUser
func (m *UserResponse) SetPrivateUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *UserResponse) SetPublicUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)() {
    m.publicUser = value
}
// UserResponseable 
type UserResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)
    GetPublicUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)
    SetPrivateUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)()
    SetPublicUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)()
}
// Blocks the blocks property
func (m *UserRequestBuilder) Blocks()(*BlocksRequestBuilder) {
    return NewBlocksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BlocksById gets an item from the github.com/octokit/kiota/.user.blocks.item collection
func (m *UserRequestBuilder) BlocksById(id string)(*BlocksWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewBlocksWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Codespaces the codespaces property
func (m *UserRequestBuilder) Codespaces()(*CodespacesRequestBuilder) {
    return NewCodespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CodespacesById gets an item from the github.com/octokit/kiota/.user.codespaces.item collection
func (m *UserRequestBuilder) CodespacesById(id string)(*CodespacesWithCodespace_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["codespace_name"] = id
    }
    return NewCodespacesWithCodespace_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Email the email property
func (m *UserRequestBuilder) Email()(*EmailRequestBuilder) {
    return NewEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Emails the emails property
func (m *UserRequestBuilder) Emails()(*EmailsRequestBuilder) {
    return NewEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Followers the followers property
func (m *UserRequestBuilder) Followers()(*FollowersRequestBuilder) {
    return NewFollowersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Following the following property
func (m *UserRequestBuilder) Following()(*FollowingRequestBuilder) {
    return NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FollowingById gets an item from the github.com/octokit/kiota/.user.following.item collection
func (m *UserRequestBuilder) FollowingById(id string)(*FollowingWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewFollowingWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get if the authenticated user is authenticated through basic authentication or OAuth with the `user` scope, then the response lists public and private profile information.If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#get-the-authenticated-user
func (m *UserRequestBuilder) Get(ctx context.Context, requestConfiguration *UserRequestBuilderGetRequestConfiguration)(UserResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserResponseable), nil
}
// Gpg_keys the gpg_keys property
func (m *UserRequestBuilder) Gpg_keys()(*Gpg_keysRequestBuilder) {
    return NewGpg_keysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Gpg_keysById gets an item from the github.com/octokit/kiota/.user.gpg_keys.item collection
func (m *UserRequestBuilder) Gpg_keysById(id string)(*Gpg_keysWithGpg_key_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["gpg_key_id"] = id
    }
    return NewGpg_keysWithGpg_key_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Installations the installations property
func (m *UserRequestBuilder) Installations()(*InstallationsRequestBuilder) {
    return NewInstallationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstallationsById gets an item from the github.com/octokit/kiota/.user.installations.item collection
func (m *UserRequestBuilder) InstallationsById(id string)(*InstallationsWithInstallation_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["installation_id"] = id
    }
    return NewInstallationsWithInstallation_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InteractionLimits the interactionLimits property
func (m *UserRequestBuilder) InteractionLimits()(*InteractionLimitsRequestBuilder) {
    return NewInteractionLimitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Issues the issues property
func (m *UserRequestBuilder) Issues()(*IssuesRequestBuilder) {
    return NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Keys the keys property
func (m *UserRequestBuilder) Keys()(*KeysRequestBuilder) {
    return NewKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// KeysById gets an item from the github.com/octokit/kiota/.user.keys.item collection
func (m *UserRequestBuilder) KeysById(id string)(*KeysWithKey_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["key_id"] = id
    }
    return NewKeysWithKey_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Marketplace_purchases the marketplace_purchases property
func (m *UserRequestBuilder) Marketplace_purchases()(*Marketplace_purchasesRequestBuilder) {
    return NewMarketplace_purchasesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Memberships the memberships property
func (m *UserRequestBuilder) Memberships()(*MembershipsRequestBuilder) {
    return NewMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Migrations the migrations property
func (m *UserRequestBuilder) Migrations()(*MigrationsRequestBuilder) {
    return NewMigrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MigrationsById gets an item from the github.com/octokit/kiota/.user.migrations.item collection
func (m *UserRequestBuilder) MigrationsById(id string)(*MigrationsWithMigration_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["migration_id"] = id
    }
    return NewMigrationsWithMigration_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Orgs the orgs property
func (m *UserRequestBuilder) Orgs()(*OrgsRequestBuilder) {
    return NewOrgsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Packages the packages property
func (m *UserRequestBuilder) Packages()(*PackagesRequestBuilder) {
    return NewPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PackagesById gets an item from the github.com/octokit/kiota/.user.packages.item collection
func (m *UserRequestBuilder) PackagesById(id string)(*PackagesWithPackage_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["package_type"] = id
    }
    return NewPackagesWithPackage_typeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch **Note:** If your email is set to private and you send an `email` parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users/#update-the-authenticated-user
func (m *UserRequestBuilder) Patch(ctx context.Context, body UserPatchRequestBodyable, requestConfiguration *UserRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreatePrivateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable), nil
}
// Projects the projects property
func (m *UserRequestBuilder) Projects()(*ProjectsRequestBuilder) {
    return NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Public_emails the public_emails property
func (m *UserRequestBuilder) Public_emails()(*Public_emailsRequestBuilder) {
    return NewPublic_emailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Repos the repos property
func (m *UserRequestBuilder) Repos()(*ReposRequestBuilder) {
    return NewReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Repository_invitations the repository_invitations property
func (m *UserRequestBuilder) Repository_invitations()(*Repository_invitationsRequestBuilder) {
    return NewRepository_invitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Repository_invitationsById gets an item from the github.com/octokit/kiota/.user.repository_invitations.item collection
func (m *UserRequestBuilder) Repository_invitationsById(id string)(*Repository_invitationsWithInvitation_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["invitation_id"] = id
    }
    return NewRepository_invitationsWithInvitation_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Ssh_signing_keys the ssh_signing_keys property
func (m *UserRequestBuilder) Ssh_signing_keys()(*Ssh_signing_keysRequestBuilder) {
    return NewSsh_signing_keysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ssh_signing_keysById gets an item from the github.com/octokit/kiota/.user.ssh_signing_keys.item collection
func (m *UserRequestBuilder) Ssh_signing_keysById(id string)(*Ssh_signing_keysWithSsh_signing_key_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ssh_signing_key_id"] = id
    }
    return NewSsh_signing_keysWithSsh_signing_key_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Starred the starred property
func (m *UserRequestBuilder) Starred()(*StarredRequestBuilder) {
    return NewStarredRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// StarredById gets an item from the github.com/octokit/kiota/.user.starred.item collection
func (m *UserRequestBuilder) StarredById(id string)(*StarredWithOwnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["owner"] = id
    }
    return NewStarredWithOwnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Subscriptions the subscriptions property
func (m *UserRequestBuilder) Subscriptions()(*SubscriptionsRequestBuilder) {
    return NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Teams the teams property
func (m *UserRequestBuilder) Teams()(*TeamsRequestBuilder) {
    return NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation if the authenticated user is authenticated through basic authentication or OAuth with the `user` scope, then the response lists public and private profile information.If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.
func (m *UserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Note:** If your email is set to private and you send an `email` parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
func (m *UserRequestBuilder) ToPatchRequestInformation(ctx context.Context, body UserPatchRequestBodyable, requestConfiguration *UserRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
