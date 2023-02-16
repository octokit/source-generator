package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationSimpleable 
type OrganizationSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetDescription()(*string)
    GetEventsUrl()(*string)
    GetHooksUrl()(*string)
    GetId()(*int32)
    GetIssuesUrl()(*string)
    GetLogin()(*string)
    GetMembersUrl()(*string)
    GetNodeId()(*string)
    GetPublicMembersUrl()(*string)
    GetReposUrl()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetDescription(value *string)()
    SetEventsUrl(value *string)()
    SetHooksUrl(value *string)()
    SetId(value *int32)()
    SetIssuesUrl(value *string)()
    SetLogin(value *string)()
    SetMembersUrl(value *string)()
    SetNodeId(value *string)()
    SetPublicMembersUrl(value *string)()
    SetReposUrl(value *string)()
    SetUrl(value *string)()
}
