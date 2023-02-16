package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgMembershipable 
type OrgMembershipable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrganization()(OrganizationSimpleable)
    GetOrganizationUrl()(*string)
    GetPermissions()(OrgMembership_permissionsable)
    GetRole()(*OrgMembership_role)
    GetState()(*OrgMembership_state)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetOrganization(value OrganizationSimpleable)()
    SetOrganizationUrl(value *string)()
    SetPermissions(value OrgMembership_permissionsable)()
    SetRole(value *OrgMembership_role)()
    SetState(value *OrgMembership_state)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
