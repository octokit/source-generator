package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PorterAuthorable 
type PorterAuthorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetId()(*int32)
    GetImportUrl()(*string)
    GetName()(*string)
    GetRemoteId()(*string)
    GetRemoteName()(*string)
    GetUrl()(*string)
    SetEmail(value *string)()
    SetId(value *int32)()
    SetImportUrl(value *string)()
    SetName(value *string)()
    SetRemoteId(value *string)()
    SetRemoteName(value *string)()
    SetUrl(value *string)()
}
