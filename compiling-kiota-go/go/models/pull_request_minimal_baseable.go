package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestMinimal_baseable 
type PullRequestMinimal_baseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRef()(*string)
    GetRepo()(PullRequestMinimal_base_repoable)
    GetSha()(*string)
    SetRef(value *string)()
    SetRepo(value PullRequestMinimal_base_repoable)()
    SetSha(value *string)()
}
