package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Actorable 
type Actorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetDisplayLogin()(*string)
    GetGravatarId()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetDisplayLogin(value *string)()
    SetGravatarId(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetUrl(value *string)()
}
