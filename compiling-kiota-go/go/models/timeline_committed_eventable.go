package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineCommittedEventable 
type TimelineCommittedEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(TimelineCommittedEvent_authorable)
    GetCommitter()(TimelineCommittedEvent_committerable)
    GetEvent()(*string)
    GetHtmlUrl()(*string)
    GetMessage()(*string)
    GetNodeId()(*string)
    GetParents()([]TimelineCommittedEvent_parentsable)
    GetSha()(*string)
    GetTree()(TimelineCommittedEvent_treeable)
    GetUrl()(*string)
    GetVerification()(TimelineCommittedEvent_verificationable)
    SetAuthor(value TimelineCommittedEvent_authorable)()
    SetCommitter(value TimelineCommittedEvent_committerable)()
    SetEvent(value *string)()
    SetHtmlUrl(value *string)()
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetParents(value []TimelineCommittedEvent_parentsable)()
    SetSha(value *string)()
    SetTree(value TimelineCommittedEvent_treeable)()
    SetUrl(value *string)()
    SetVerification(value TimelineCommittedEvent_verificationable)()
}
