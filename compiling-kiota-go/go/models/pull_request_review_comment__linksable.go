package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestReviewComment__linksable 
type PullRequestReviewComment__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(PullRequestReviewComment__links_htmlable)
    GetPullRequest()(PullRequestReviewComment__links_pull_requestable)
    GetSelf()(PullRequestReviewComment__links_selfable)
    SetHtml(value PullRequestReviewComment__links_htmlable)()
    SetPullRequest(value PullRequestReviewComment__links_pull_requestable)()
    SetSelf(value PullRequestReviewComment__links_selfable)()
}
