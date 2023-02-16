package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableIntegrationable 
type NullableIntegrationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEvents()([]string)
    GetExternalUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetInstallationsCount()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOwner()(NullableSimpleUserable)
    GetPem()(*string)
    GetPermissions()(NullableIntegration_permissionsable)
    GetSlug()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetWebhookSecret()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEvents(value []string)()
    SetExternalUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetInstallationsCount(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOwner(value NullableSimpleUserable)()
    SetPem(value *string)()
    SetPermissions(value NullableIntegration_permissionsable)()
    SetSlug(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetWebhookSecret(value *string)()
}
