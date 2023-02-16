package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitTagable 
type GitTagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    GetNodeId()(*string)
    GetObject()(GitTag_objectable)
    GetSha()(*string)
    GetTag()(*string)
    GetTagger()(GitTag_taggerable)
    GetUrl()(*string)
    GetVerification()(Verificationable)
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetObject(value GitTag_objectable)()
    SetSha(value *string)()
    SetTag(value *string)()
    SetTagger(value GitTag_taggerable)()
    SetUrl(value *string)()
    SetVerification(value Verificationable)()
}
