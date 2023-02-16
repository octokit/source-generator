package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationInvitationable 
type OrganizationInvitationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetEmail()(*string)
    GetFailedAt()(*string)
    GetFailedReason()(*string)
    GetId()(*int32)
    GetInvitationSource()(*string)
    GetInvitationTeamsUrl()(*string)
    GetInviter()(SimpleUserable)
    GetLogin()(*string)
    GetNodeId()(*string)
    GetRole()(*string)
    GetTeamCount()(*int32)
    SetCreatedAt(value *string)()
    SetEmail(value *string)()
    SetFailedAt(value *string)()
    SetFailedReason(value *string)()
    SetId(value *int32)()
    SetInvitationSource(value *string)()
    SetInvitationTeamsUrl(value *string)()
    SetInviter(value SimpleUserable)()
    SetLogin(value *string)()
    SetNodeId(value *string)()
    SetRole(value *string)()
    SetTeamCount(value *int32)()
}
