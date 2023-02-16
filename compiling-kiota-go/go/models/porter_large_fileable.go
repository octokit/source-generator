package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PorterLargeFileable 
type PorterLargeFileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOid()(*string)
    GetPath()(*string)
    GetRefName()(*string)
    GetSize()(*int32)
    SetOid(value *string)()
    SetPath(value *string)()
    SetRefName(value *string)()
    SetSize(value *int32)()
}
