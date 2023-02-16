package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryInvitationable 
type RepositoryInvitationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpired()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetInvitee()(NullableSimpleUserable)
    GetInviter()(NullableSimpleUserable)
    GetNodeId()(*string)
    GetPermissions()(*RepositoryInvitation_permissions)
    GetRepository()(MinimalRepositoryable)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpired(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetInvitee(value NullableSimpleUserable)()
    SetInviter(value NullableSimpleUserable)()
    SetNodeId(value *string)()
    SetPermissions(value *RepositoryInvitation_permissions)()
    SetRepository(value MinimalRepositoryable)()
    SetUrl(value *string)()
}
