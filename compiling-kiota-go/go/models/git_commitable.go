package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitCommitable 
type GitCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(GitCommit_authorable)
    GetCommitter()(GitCommit_committerable)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]GitCommit_parentsable)
    GetSha()(*string)
    GetTree()(GitCommit_treeable)
    GetUrl()(*string)
    GetVerification()(GitCommit_verificationable)
    SetAuthor(value GitCommit_authorable)()
    SetCommitter(value GitCommit_committerable)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []GitCommit_parentsable)()
    SetSha(value *string)()
    SetTree(value GitCommit_treeable)()
    SetUrl(value *string)()
    SetVerification(value GitCommit_verificationable)()
}
