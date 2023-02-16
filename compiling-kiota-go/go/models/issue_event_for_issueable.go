package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueEventForIssueable 
type IssueEventForIssueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedToProjectIssueEvent()(AddedToProjectIssueEventable)
    GetAssignedIssueEvent()(AssignedIssueEventable)
    GetConvertedNoteToIssueIssueEvent()(ConvertedNoteToIssueIssueEventable)
    GetDemilestonedIssueEvent()(DemilestonedIssueEventable)
    GetLabeledIssueEvent()(LabeledIssueEventable)
    GetLockedIssueEvent()(LockedIssueEventable)
    GetMilestonedIssueEvent()(MilestonedIssueEventable)
    GetMovedColumnInProjectIssueEvent()(MovedColumnInProjectIssueEventable)
    GetRemovedFromProjectIssueEvent()(RemovedFromProjectIssueEventable)
    GetRenamedIssueEvent()(RenamedIssueEventable)
    GetReviewDismissedIssueEvent()(ReviewDismissedIssueEventable)
    GetReviewRequestedIssueEvent()(ReviewRequestedIssueEventable)
    GetReviewRequestRemovedIssueEvent()(ReviewRequestRemovedIssueEventable)
    GetUnassignedIssueEvent()(UnassignedIssueEventable)
    GetUnlabeledIssueEvent()(UnlabeledIssueEventable)
    SetAddedToProjectIssueEvent(value AddedToProjectIssueEventable)()
    SetAssignedIssueEvent(value AssignedIssueEventable)()
    SetConvertedNoteToIssueIssueEvent(value ConvertedNoteToIssueIssueEventable)()
    SetDemilestonedIssueEvent(value DemilestonedIssueEventable)()
    SetLabeledIssueEvent(value LabeledIssueEventable)()
    SetLockedIssueEvent(value LockedIssueEventable)()
    SetMilestonedIssueEvent(value MilestonedIssueEventable)()
    SetMovedColumnInProjectIssueEvent(value MovedColumnInProjectIssueEventable)()
    SetRemovedFromProjectIssueEvent(value RemovedFromProjectIssueEventable)()
    SetRenamedIssueEvent(value RenamedIssueEventable)()
    SetReviewDismissedIssueEvent(value ReviewDismissedIssueEventable)()
    SetReviewRequestedIssueEvent(value ReviewRequestedIssueEventable)()
    SetReviewRequestRemovedIssueEvent(value ReviewRequestRemovedIssueEventable)()
    SetUnassignedIssueEvent(value UnassignedIssueEventable)()
    SetUnlabeledIssueEvent(value UnlabeledIssueEventable)()
}
