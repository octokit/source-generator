package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitSearchResultItemable 
type CommitSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetCommentsUrl()(*string)
    GetCommit()(CommitSearchResultItem_commitable)
    GetCommitter()(NullableGitUserable)
    GetHtmlUrl()(*string)
    GetNodeId()(*string)
    GetParents()([]CommitSearchResultItem_parentsable)
    GetRepository()(MinimalRepositoryable)
    GetScore()(*int64)
    GetSha()(*string)
    GetTextMatches()([]CommitSearchResultItem_text_matchesable)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetCommentsUrl(value *string)()
    SetCommit(value CommitSearchResultItem_commitable)()
    SetCommitter(value NullableGitUserable)()
    SetHtmlUrl(value *string)()
    SetNodeId(value *string)()
    SetParents(value []CommitSearchResultItem_parentsable)()
    SetRepository(value MinimalRepositoryable)()
    SetScore(value *int64)()
    SetSha(value *string)()
    SetTextMatches(value []CommitSearchResultItem_text_matchesable)()
    SetUrl(value *string)()
}
