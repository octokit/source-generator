package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestSimple_headable 
type PullRequestSimple_headable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabel()(*string)
    GetRef()(*string)
    GetRepo()(Repositoryable)
    GetSha()(*string)
    GetUser()(NullableSimpleUserable)
    SetLabel(value *string)()
    SetRef(value *string)()
    SetRepo(value Repositoryable)()
    SetSha(value *string)()
    SetUser(value NullableSimpleUserable)()
}
