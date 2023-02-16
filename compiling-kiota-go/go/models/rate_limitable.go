package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RateLimitable 
type RateLimitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLimit()(*int32)
    GetRemaining()(*int32)
    GetReset()(*int32)
    GetUsed()(*int32)
    SetLimit(value *int32)()
    SetRemaining(value *int32)()
    SetReset(value *int32)()
    SetUsed(value *int32)()
}
