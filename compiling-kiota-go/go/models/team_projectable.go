package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamProjectable 
type TeamProjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetColumnsUrl()(*string)
    GetCreatedAt()(*string)
    GetCreator()(SimpleUserable)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOrganizationPermission()(*string)
    GetOwnerUrl()(*string)
    GetPermissions()(TeamProject_permissionsable)
    GetPrivate()(*bool)
    GetState()(*string)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetColumnsUrl(value *string)()
    SetCreatedAt(value *string)()
    SetCreator(value SimpleUserable)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOrganizationPermission(value *string)()
    SetOwnerUrl(value *string)()
    SetPermissions(value TeamProject_permissionsable)()
    SetPrivate(value *bool)()
    SetState(value *string)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
