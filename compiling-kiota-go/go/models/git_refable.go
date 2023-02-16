package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitRefable 
type GitRefable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNodeId()(*string)
    GetObject()(GitRef_objectable)
    GetRef()(*string)
    GetUrl()(*string)
    SetNodeId(value *string)()
    SetObject(value GitRef_objectable)()
    SetRef(value *string)()
    SetUrl(value *string)()
}
