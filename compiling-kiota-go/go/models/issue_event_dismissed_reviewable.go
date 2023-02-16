package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueEventDismissedReviewable 
type IssueEventDismissedReviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDismissalCommitId()(*string)
    GetDismissalMessage()(*string)
    GetReviewId()(*int32)
    GetState()(*string)
    SetDismissalCommitId(value *string)()
    SetDismissalMessage(value *string)()
    SetReviewId(value *int32)()
    SetState(value *string)()
}
