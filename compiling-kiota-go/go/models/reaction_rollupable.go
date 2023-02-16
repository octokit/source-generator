package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReactionRollupable 
type ReactionRollupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfused()(*int32)
    GetEyes()(*int32)
    GetHeart()(*int32)
    GetHooray()(*int32)
    GetLaugh()(*int32)
    GetOne()(*int32)
    GetRocket()(*int32)
    GetTotalCount()(*int32)
    GetUrl()(*string)
    SetConfused(value *int32)()
    SetEyes(value *int32)()
    SetHeart(value *int32)()
    SetHooray(value *int32)()
    SetLaugh(value *int32)()
    SetOne(value *int32)()
    SetRocket(value *int32)()
    SetTotalCount(value *int32)()
    SetUrl(value *string)()
}
