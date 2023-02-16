package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineAssignedIssueEventable 
type TimelineAssignedIssueEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(SimpleUserable)
    GetAssignee()(SimpleUserable)
    GetCommitId()(*string)
    GetCommitUrl()(*string)
    GetCreatedAt()(*string)
    GetEvent()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetUrl()(*string)
    SetActor(value SimpleUserable)()
    SetAssignee(value SimpleUserable)()
    SetCommitId(value *string)()
    SetCommitUrl(value *string)()
    SetCreatedAt(value *string)()
    SetEvent(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetUrl(value *string)()
}
