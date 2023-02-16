package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamFullable 
type TeamFullable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLdapDn()(*string)
    GetMembersCount()(*int32)
    GetMembersUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetOrganization()(TeamOrganizationable)
    GetParent()(NullableTeamSimpleable)
    GetPermission()(*string)
    GetPrivacy()(*TeamFull_privacy)
    GetReposCount()(*int32)
    GetRepositoriesUrl()(*string)
    GetSlug()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLdapDn(value *string)()
    SetMembersCount(value *int32)()
    SetMembersUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOrganization(value TeamOrganizationable)()
    SetParent(value NullableTeamSimpleable)()
    SetPermission(value *string)()
    SetPrivacy(value *TeamFull_privacy)()
    SetReposCount(value *int32)()
    SetRepositoriesUrl(value *string)()
    SetSlug(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
