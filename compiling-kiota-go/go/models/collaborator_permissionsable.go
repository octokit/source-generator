package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Collaborator_permissionsable 
type Collaborator_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmin()(*bool)
    GetMaintain()(*bool)
    GetPull()(*bool)
    GetPush()(*bool)
    GetTriage()(*bool)
    SetAdmin(value *bool)()
    SetMaintain(value *bool)()
    SetPull(value *bool)()
    SetPush(value *bool)()
    SetTriage(value *bool)()
}
