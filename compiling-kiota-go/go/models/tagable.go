package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Tagable 
type Tagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(Tag_commitable)
    GetName()(*string)
    GetNodeId()(*string)
    GetTarballUrl()(*string)
    GetZipballUrl()(*string)
    SetCommit(value Tag_commitable)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetTarballUrl(value *string)()
    SetZipballUrl(value *string)()
}
