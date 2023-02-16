package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Commit_statsable 
type Commit_statsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditions()(*int32)
    GetDeletions()(*int32)
    GetTotal()(*int32)
    SetAdditions(value *int32)()
    SetDeletions(value *int32)()
    SetTotal(value *int32)()
}
