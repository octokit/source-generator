package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequest_head_repo_licenseable 
type PullRequest_head_repo_licenseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetSpdxId()(*string)
    GetUrl()(*string)
    SetKey(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetSpdxId(value *string)()
    SetUrl(value *string)()
}
