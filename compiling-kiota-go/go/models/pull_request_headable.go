package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequest_headable 
type PullRequest_headable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabel()(*string)
    GetRef()(*string)
    GetRepo()(PullRequest_head_repoable)
    GetSha()(*string)
    GetUser()(PullRequest_head_userable)
    SetLabel(value *string)()
    SetRef(value *string)()
    SetRepo(value PullRequest_head_repoable)()
    SetSha(value *string)()
    SetUser(value PullRequest_head_userable)()
}
