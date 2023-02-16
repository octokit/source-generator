package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningAlertable 
type SecretScanningAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetLocationsUrl()(*string)
    GetNumber()(*int32)
    GetPushProtectionBypassed()(*bool)
    GetPushProtectionBypassedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPushProtectionBypassedBy()(NullableSimpleUserable)
    GetResolution()(*SecretScanningAlertResolution)
    GetResolutionComment()(*string)
    GetResolvedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResolvedBy()(NullableSimpleUserable)
    GetSecret()(*string)
    GetSecretType()(*string)
    GetSecretTypeDisplayName()(*string)
    GetState()(*SecretScanningAlertState)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetLocationsUrl(value *string)()
    SetNumber(value *int32)()
    SetPushProtectionBypassed(value *bool)()
    SetPushProtectionBypassedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPushProtectionBypassedBy(value NullableSimpleUserable)()
    SetResolution(value *SecretScanningAlertResolution)()
    SetResolutionComment(value *string)()
    SetResolvedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResolvedBy(value NullableSimpleUserable)()
    SetSecret(value *string)()
    SetSecretType(value *string)()
    SetSecretTypeDisplayName(value *string)()
    SetState(value *SecretScanningAlertState)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
