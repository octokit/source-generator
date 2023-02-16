package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchRestrictionPolicy_appsable 
type BranchRestrictionPolicy_appsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetDescription()(*string)
    GetEvents()([]string)
    GetExternalUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwner()(BranchRestrictionPolicy_apps_ownerable)
    GetPermissions()(BranchRestrictionPolicy_apps_permissionsable)
    GetSlug()(*string)
    GetUpdatedAt()(*string)
    SetCreatedAt(value *string)()
    SetDescription(value *string)()
    SetEvents(value []string)()
    SetExternalUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwner(value BranchRestrictionPolicy_apps_ownerable)()
    SetPermissions(value BranchRestrictionPolicy_apps_permissionsable)()
    SetSlug(value *string)()
    SetUpdatedAt(value *string)()
}
