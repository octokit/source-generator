package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CombinedCommitStatusable 
type CombinedCommitStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitUrl()(*string)
    GetRepository()(MinimalRepositoryable)
    GetSha()(*string)
    GetState()(*string)
    GetStatuses()([]SimpleCommitStatusable)
    GetTotalCount()(*int32)
    GetUrl()(*string)
    SetCommitUrl(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetSha(value *string)()
    SetState(value *string)()
    SetStatuses(value []SimpleCommitStatusable)()
    SetTotalCount(value *int32)()
    SetUrl(value *string)()
}
