package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReviewComment__linksable 
type ReviewComment__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtml()(Linkable)
    GetPullRequest()(Linkable)
    GetSelf()(Linkable)
    SetHtml(value Linkable)()
    SetPullRequest(value Linkable)()
    SetSelf(value Linkable)()
}
