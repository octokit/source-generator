package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestSimple__linksable 
type PullRequestSimple__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(Linkable)
    GetCommits()(Linkable)
    GetHtml()(Linkable)
    GetIssue()(Linkable)
    GetReviewComment()(Linkable)
    GetReviewComments()(Linkable)
    GetSelf()(Linkable)
    GetStatuses()(Linkable)
    SetComments(value Linkable)()
    SetCommits(value Linkable)()
    SetHtml(value Linkable)()
    SetIssue(value Linkable)()
    SetReviewComment(value Linkable)()
    SetReviewComments(value Linkable)()
    SetSelf(value Linkable)()
    SetStatuses(value Linkable)()
}
