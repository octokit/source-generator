package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitComparisonable 
type CommitComparisonable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAheadBy()(*int32)
    GetBaseCommit()(Commitable)
    GetBehindBy()(*int32)
    GetCommits()([]Commitable)
    GetDiffUrl()(*string)
    GetFiles()([]DiffEntryable)
    GetHtmlUrl()(*string)
    GetMergeBaseCommit()(Commitable)
    GetPatchUrl()(*string)
    GetPermalinkUrl()(*string)
    GetStatus()(*CommitComparison_status)
    GetTotalCommits()(*int32)
    GetUrl()(*string)
    SetAheadBy(value *int32)()
    SetBaseCommit(value Commitable)()
    SetBehindBy(value *int32)()
    SetCommits(value []Commitable)()
    SetDiffUrl(value *string)()
    SetFiles(value []DiffEntryable)()
    SetHtmlUrl(value *string)()
    SetMergeBaseCommit(value Commitable)()
    SetPatchUrl(value *string)()
    SetPermalinkUrl(value *string)()
    SetStatus(value *CommitComparison_status)()
    SetTotalCommits(value *int32)()
    SetUrl(value *string)()
}
