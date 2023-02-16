package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Commitable 
type Commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetCommentsUrl()(*string)
    GetCommit()(Commit_commitable)
    GetCommitter()(NullableSimpleUserable)
    GetFiles()([]DiffEntryable)
    GetHtmlUrl()(*string)
    GetNodeId()(*string)
    GetParents()([]Commit_parentsable)
    GetSha()(*string)
    GetStats()(Commit_statsable)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetCommentsUrl(value *string)()
    SetCommit(value Commit_commitable)()
    SetCommitter(value NullableSimpleUserable)()
    SetFiles(value []DiffEntryable)()
    SetHtmlUrl(value *string)()
    SetNodeId(value *string)()
    SetParents(value []Commit_parentsable)()
    SetSha(value *string)()
    SetStats(value Commit_statsable)()
    SetUrl(value *string)()
}
