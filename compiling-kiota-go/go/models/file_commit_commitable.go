package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileCommit_commitable 
type FileCommit_commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(FileCommit_commit_authorable)
    GetCommitter()(FileCommit_commit_committerable)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]FileCommit_commit_parentsable)
    GetSha()(*string)
    GetTree()(FileCommit_commit_treeable)
    GetUrl()(*string)
    GetVerification()(FileCommit_commit_verificationable)
    SetAuthor(value FileCommit_commit_authorable)()
    SetCommitter(value FileCommit_commit_committerable)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []FileCommit_commit_parentsable)()
    SetSha(value *string)()
    SetTree(value FileCommit_commit_treeable)()
    SetUrl(value *string)()
    SetVerification(value FileCommit_commit_verificationable)()
}
