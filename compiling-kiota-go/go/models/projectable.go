package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Projectable 
type Projectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetColumnsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOrganizationPermission()(*Project_organization_permission)
    GetOwnerUrl()(*string)
    GetPrivate()(*bool)
    GetState()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetBody(value *string)()
    SetColumnsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOrganizationPermission(value *Project_organization_permission)()
    SetOwnerUrl(value *string)()
    SetPrivate(value *bool)()
    SetState(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
