package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Event_payloadable 
type Event_payloadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetComment()(IssueCommentable)
    GetIssue()(Issueable)
    GetPages()([]Event_payload_pagesable)
    SetAction(value *string)()
    SetComment(value IssueCommentable)()
    SetIssue(value Issueable)()
    SetPages(value []Event_payload_pagesable)()
}
