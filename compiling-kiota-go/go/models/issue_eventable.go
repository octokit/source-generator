package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueEventable 
type IssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(NullableSimpleUserable)
    GetAssignee()(NullableSimpleUserable)
    GetAssigner()(NullableSimpleUserable)
    GetAuthorAssociation()(*AuthorAssociation)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedReview()(IssueEventDismissedReviewable)
    GetEvent()(*string)
    GetId()(*int32)
    GetIssue()(NullableIssueable)
    GetLabel()(IssueEventLabelable)
    GetLockReason()(*string)
    GetMilestone()(IssueEventMilestoneable)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetProjectCard()(IssueEventProjectCardable)
    GetRename()(IssueEventRenameable)
    GetRequestedReviewer()(NullableSimpleUserable)
    GetRequestedTeam()(Teamable)
    GetReviewRequester()(NullableSimpleUserable)
    GetUrl()(*string)
    SetActor(value NullableSimpleUserable)()
    SetAssignee(value NullableSimpleUserable)()
    SetAssigner(value NullableSimpleUserable)()
    SetAuthorAssociation(value *AuthorAssociation)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedReview(value IssueEventDismissedReviewable)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetIssue(value NullableIssueable)()
    SetLabel(value IssueEventLabelable)()
    SetLockReason(value *string)()
    SetMilestone(value IssueEventMilestoneable)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetProjectCard(value IssueEventProjectCardable)()
    SetRename(value IssueEventRenameable)()
    SetRequestedReviewer(value NullableSimpleUserable)()
    SetRequestedTeam(value Teamable)()
    SetReviewRequester(value NullableSimpleUserable)()
    SetUrl(value *string)()
}
