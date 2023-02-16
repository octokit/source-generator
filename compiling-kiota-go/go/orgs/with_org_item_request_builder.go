package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// WithOrgItemRequestBuilder builds and executes requests for operations under \orgs\{org}
type WithOrgItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationFull422Error composed type wrapper for classes validationError, validationErrorSimple
type OrganizationFull422Error struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type validationError
    validationError ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable
    // Composed type representation for type validationErrorSimple
    validationErrorSimple ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable
}
// NewOrganizationFull422Error instantiates a new OrganizationFull422Error and sets the default values.
func NewOrganizationFull422Error()(*OrganizationFull422Error) {
    m := &OrganizationFull422Error{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationFull422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationFull422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOrganizationFull422Error()
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
func (m *OrganizationFull422Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationFull422Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetValidationError gets the validationError property value. Composed type representation for type validationError
func (m *OrganizationFull422Error) GetValidationError()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable) {
    return m.validationError
}
// GetValidationErrorSimple gets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *OrganizationFull422Error) GetValidationErrorSimple()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable) {
    return m.validationErrorSimple
}
// Serialize serializes information the current object
func (m *OrganizationFull422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValidationError() != nil {
        err := writer.WriteObjectValue("", m.GetValidationError())
        if err != nil {
            return err
        }
    } else if m.GetValidationErrorSimple() != nil {
        err := writer.WriteObjectValue("", m.GetValidationErrorSimple())
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
func (m *OrganizationFull422Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValidationError sets the validationError property value. Composed type representation for type validationError
func (m *OrganizationFull422Error) SetValidationError(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable)() {
    m.validationError = value
}
// SetValidationErrorSimple sets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *OrganizationFull422Error) SetValidationErrorSimple(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable)() {
    m.validationErrorSimple = value
}
// WithOrgItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithOrgItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithOrgItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithOrgItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationFull422Errorable 
type OrganizationFull422Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValidationError()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable)
    GetValidationErrorSimple()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable)
    SetValidationError(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable)()
    SetValidationErrorSimple(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable)()
}
// Actions the actions property
func (m *WithOrgItemRequestBuilder) Actions()(*ItemActionsRequestBuilder) {
    return NewItemActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Blocks the blocks property
func (m *WithOrgItemRequestBuilder) Blocks()(*ItemBlocksRequestBuilder) {
    return NewItemBlocksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BlocksById gets an item from the github.com/octokit/kiota/.orgs.item.blocks.item collection
func (m *WithOrgItemRequestBuilder) BlocksById(id string)(*ItemBlocksWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemBlocksWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CodeScanning the codeScanning property
func (m *WithOrgItemRequestBuilder) CodeScanning()(*ItemCodeScanningRequestBuilder) {
    return NewItemCodeScanningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Codespaces the codespaces property
func (m *WithOrgItemRequestBuilder) Codespaces()(*ItemCodespacesRequestBuilder) {
    return NewItemCodespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewWithOrgItemRequestBuilderInternal instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    m := &WithOrgItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithOrgItemRequestBuilder instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Dependabot the dependabot property
func (m *WithOrgItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Events the events property
func (m *WithOrgItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Failed_invitations the failed_invitations property
func (m *WithOrgItemRequestBuilder) Failed_invitations()(*ItemFailed_invitationsRequestBuilder) {
    return NewItemFailed_invitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get to see many of the organization response values, you need to be an authenticated organization owner with the `admin:org` scope. When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/).GitHub Apps with the `Organization plan` permission can use this endpoint to retrieve information about an organization's GitHub plan. See "[Authenticating with GitHub Apps](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/)" for details. For an example response, see 'Response with GitHub plan information' below."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#get-an-organization
func (m *WithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.OrganizationFullable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.OrganizationFullable), nil
}
// Hooks the hooks property
func (m *WithOrgItemRequestBuilder) Hooks()(*ItemHooksRequestBuilder) {
    return NewItemHooksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// HooksById gets an item from the github.com/octokit/kiota/.orgs.item.hooks.item collection
func (m *WithOrgItemRequestBuilder) HooksById(id string)(*ItemHooksWithHook_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hook_id"] = id
    }
    return NewItemHooksWithHook_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Installation the installation property
func (m *WithOrgItemRequestBuilder) Installation()(*ItemInstallationRequestBuilder) {
    return NewItemInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Installations the installations property
func (m *WithOrgItemRequestBuilder) Installations()(*ItemInstallationsRequestBuilder) {
    return NewItemInstallationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InteractionLimits the interactionLimits property
func (m *WithOrgItemRequestBuilder) InteractionLimits()(*ItemInteractionLimitsRequestBuilder) {
    return NewItemInteractionLimitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Invitations the invitations property
func (m *WithOrgItemRequestBuilder) Invitations()(*ItemInvitationsRequestBuilder) {
    return NewItemInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InvitationsById gets an item from the github.com/octokit/kiota/.orgs.item.invitations.item collection
func (m *WithOrgItemRequestBuilder) InvitationsById(id string)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["invitation_id"] = id
    }
    return NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Issues the issues property
func (m *WithOrgItemRequestBuilder) Issues()(*ItemIssuesRequestBuilder) {
    return NewItemIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Members the members property
func (m *WithOrgItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById gets an item from the github.com/octokit/kiota/.orgs.item.members.item collection
func (m *WithOrgItemRequestBuilder) MembersById(id string)(*ItemMembersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemMembersWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Memberships the memberships property
func (m *WithOrgItemRequestBuilder) Memberships()(*ItemMembershipsRequestBuilder) {
    return NewItemMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembershipsById gets an item from the github.com/octokit/kiota/.orgs.item.memberships.item collection
func (m *WithOrgItemRequestBuilder) MembershipsById(id string)(*ItemMembershipsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemMembershipsWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Migrations the migrations property
func (m *WithOrgItemRequestBuilder) Migrations()(*ItemMigrationsRequestBuilder) {
    return NewItemMigrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MigrationsById gets an item from the github.com/octokit/kiota/.orgs.item.migrations.item collection
func (m *WithOrgItemRequestBuilder) MigrationsById(id string)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["migration_id"] = id
    }
    return NewItemMigrationsWithMigration_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Outside_collaborators the outside_collaborators property
func (m *WithOrgItemRequestBuilder) Outside_collaborators()(*ItemOutside_collaboratorsRequestBuilder) {
    return NewItemOutside_collaboratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Outside_collaboratorsById gets an item from the github.com/octokit/kiota/.orgs.item.outside_collaborators.item collection
func (m *WithOrgItemRequestBuilder) Outside_collaboratorsById(id string)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Packages the packages property
func (m *WithOrgItemRequestBuilder) Packages()(*ItemPackagesRequestBuilder) {
    return NewItemPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PackagesById gets an item from the github.com/octokit/kiota/.orgs.item.packages.item collection
func (m *WithOrgItemRequestBuilder) PackagesById(id string)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["package_type"] = id
    }
    return NewItemPackagesWithPackage_typeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch **Parameter Deprecation Notice:** GitHub will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Enables an authenticated organization owner with the `admin:org` scope to update the organization's profile and member privileges.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#update-an-organization
func (m *WithOrgItemRequestBuilder) Patch(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *WithOrgItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.OrganizationFullable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "409": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": CreateOrganizationFull422ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.OrganizationFullable), nil
}
// Projects the projects property
func (m *WithOrgItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Public_members the public_members property
func (m *WithOrgItemRequestBuilder) Public_members()(*ItemPublic_membersRequestBuilder) {
    return NewItemPublic_membersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Public_membersById gets an item from the github.com/octokit/kiota/.orgs.item.public_members.item collection
func (m *WithOrgItemRequestBuilder) Public_membersById(id string)(*ItemPublic_membersWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return NewItemPublic_membersWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Repos the repos property
func (m *WithOrgItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecretScanning the secretScanning property
func (m *WithOrgItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecurityManagers the securityManagers property
func (m *WithOrgItemRequestBuilder) SecurityManagers()(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Settings the settings property
func (m *WithOrgItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Teams the teams property
func (m *WithOrgItemRequestBuilder) Teams()(*ItemTeamsRequestBuilder) {
    return NewItemTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TeamsById gets an item from the github.com/octokit/kiota/.orgs.item.teams.item collection
func (m *WithOrgItemRequestBuilder) TeamsById(id string)(*ItemTeamsWithTeam_slugItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team_slug"] = id
    }
    return NewItemTeamsWithTeam_slugItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation to see many of the organization response values, you need to be an authenticated organization owner with the `admin:org` scope. When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/).GitHub Apps with the `Organization plan` permission can use this endpoint to retrieve information about an organization's GitHub plan. See "[Authenticating with GitHub Apps](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/)" for details. For an example response, see 'Response with GitHub plan information' below."
func (m *WithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Parameter Deprecation Notice:** GitHub will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Enables an authenticated organization owner with the `admin:org` scope to update the organization's profile and member privileges.
func (m *WithOrgItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *WithOrgItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
