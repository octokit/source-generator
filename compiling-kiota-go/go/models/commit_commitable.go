package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Commit_commitable 
type Commit_commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableGitUserable)
    GetCommentCount()(*int32)
    GetCommitter()(NullableGitUserable)
    GetMessage()(*string)
    GetTree()(Commit_commit_treeable)
    GetUrl()(*string)
    GetVerification()(Verificationable)
    SetAuthor(value NullableGitUserable)()
    SetCommentCount(value *int32)()
    SetCommitter(value NullableGitUserable)()
    SetMessage(value *string)()
    SetTree(value Commit_commit_treeable)()
    SetUrl(value *string)()
    SetVerification(value Verificationable)()
}
