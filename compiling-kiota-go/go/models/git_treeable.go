package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitTreeable 
type GitTreeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSha()(*string)
    GetTree()([]GitTree_treeable)
    GetTruncated()(*bool)
    GetUrl()(*string)
    SetSha(value *string)()
    SetTree(value []GitTree_treeable)()
    SetTruncated(value *bool)()
    SetUrl(value *string)()
}
