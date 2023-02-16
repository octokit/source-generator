package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeployKeyable 
type DeployKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedBy()(*string)
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetKey()(*string)
    GetLastUsed()(*string)
    GetReadOnly()(*bool)
    GetTitle()(*string)
    GetUrl()(*string)
    GetVerified()(*bool)
    SetAddedBy(value *string)()
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetKey(value *string)()
    SetLastUsed(value *string)()
    SetReadOnly(value *bool)()
    SetTitle(value *string)()
    SetUrl(value *string)()
    SetVerified(value *bool)()
}
