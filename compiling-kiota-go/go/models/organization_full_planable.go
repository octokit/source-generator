package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationFull_planable 
type OrganizationFull_planable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilledSeats()(*int32)
    GetName()(*string)
    GetPrivateRepos()(*int32)
    GetSeats()(*int32)
    GetSpace()(*int32)
    SetFilledSeats(value *int32)()
    SetName(value *string)()
    SetPrivateRepos(value *int32)()
    SetSeats(value *int32)()
    SetSpace(value *int32)()
}
