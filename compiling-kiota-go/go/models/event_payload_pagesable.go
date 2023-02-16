package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Event_payload_pagesable 
type Event_payload_pagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetHtmlUrl()(*string)
    GetPageName()(*string)
    GetSha()(*string)
    GetSummary()(*string)
    GetTitle()(*string)
    SetAction(value *string)()
    SetHtmlUrl(value *string)()
    SetPageName(value *string)()
    SetSha(value *string)()
    SetSummary(value *string)()
    SetTitle(value *string)()
}
