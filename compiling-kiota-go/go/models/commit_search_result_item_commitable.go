package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitSearchResultItem_commitable 
type CommitSearchResultItem_commitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(CommitSearchResultItem_commit_authorable)
    GetCommentCount()(*int32)
    GetCommitter()(NullableGitUserable)
    GetMessage()(*string)
    GetTree()(CommitSearchResultItem_commit_treeable)
    GetUrl()(*string)
    GetVerification()(Verificationable)
    SetAuthor(value CommitSearchResultItem_commit_authorable)()
    SetCommentCount(value *int32)()
    SetCommitter(value NullableGitUserable)()
    SetMessage(value *string)()
    SetTree(value CommitSearchResultItem_commit_treeable)()
    SetUrl(value *string)()
    SetVerification(value Verificationable)()
}
