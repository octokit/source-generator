package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesPostRequestBodyable 
type ItemItemCodespacesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientIp()(*string)
    GetDevcontainerPath()(*string)
    GetDisplayName()(*string)
    GetIdleTimeoutMinutes()(*int32)
    GetLocation()(*string)
    GetMachine()(*string)
    GetMultiRepoPermissionsOptOut()(*bool)
    GetRef()(*string)
    GetRetentionPeriodMinutes()(*int32)
    GetWorkingDirectory()(*string)
    SetClientIp(value *string)()
    SetDevcontainerPath(value *string)()
    SetDisplayName(value *string)()
    SetIdleTimeoutMinutes(value *int32)()
    SetLocation(value *string)()
    SetMachine(value *string)()
    SetMultiRepoPermissionsOptOut(value *bool)()
    SetRef(value *string)()
    SetRetentionPeriodMinutes(value *int32)()
    SetWorkingDirectory(value *string)()
}
