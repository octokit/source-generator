package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespaceMachineable 
type CodespaceMachineable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpus()(*int32)
    GetDisplayName()(*string)
    GetMemoryInBytes()(*int32)
    GetName()(*string)
    GetOperatingSystem()(*string)
    GetPrebuildAvailability()(*CodespaceMachine_prebuild_availability)
    GetStorageInBytes()(*int32)
    SetCpus(value *int32)()
    SetDisplayName(value *string)()
    SetMemoryInBytes(value *int32)()
    SetName(value *string)()
    SetOperatingSystem(value *string)()
    SetPrebuildAvailability(value *CodespaceMachine_prebuild_availability)()
    SetStorageInBytes(value *int32)()
}
