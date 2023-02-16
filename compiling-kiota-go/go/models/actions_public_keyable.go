package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsPublicKeyable 
type ActionsPublicKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetKey()(*string)
    GetKeyId()(*string)
    GetTitle()(*string)
    GetUrl()(*string)
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetKeyId(value *string)()
    SetTitle(value *string)()
    SetUrl(value *string)()
}
