package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HookDeliveryItemable 
type HookDeliveryItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetDeliveredAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDuration()(*int64)
    GetEvent()(*string)
    GetGuid()(*string)
    GetId()(*int32)
    GetInstallationId()(*int32)
    GetRedelivery()(*bool)
    GetRepositoryId()(*int32)
    GetStatus()(*string)
    GetStatusCode()(*int32)
    SetAction(value *string)()
    SetDeliveredAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDuration(value *int64)()
    SetEvent(value *string)()
    SetGuid(value *string)()
    SetId(value *int32)()
    SetInstallationId(value *int32)()
    SetRedelivery(value *bool)()
    SetRepositoryId(value *int32)()
    SetStatus(value *string)()
    SetStatusCode(value *int32)()
}
