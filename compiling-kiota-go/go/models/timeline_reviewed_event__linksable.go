package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineReviewedEvent__linksable 
type TimelineReviewedEvent__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(TimelineReviewedEvent__links_htmlable)
    GetPullRequest()(TimelineReviewedEvent__links_pull_requestable)
    SetHtml(value TimelineReviewedEvent__links_htmlable)()
    SetPullRequest(value TimelineReviewedEvent__links_pull_requestable)()
}
