package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckSuiteable 
type CheckSuiteable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAfter()(*string)
    GetApp()(NullableIntegrationable)
    GetBefore()(*string)
    GetCheckRunsUrl()(*string)
    GetConclusion()(*CheckSuite_conclusion)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHeadBranch()(*string)
    GetHeadCommit()(SimpleCommitable)
    GetHeadSha()(*string)
    GetId()(*int32)
    GetLatestCheckRunsCount()(*int32)
    GetNodeId()(*string)
    GetPullRequests()([]PullRequestMinimalable)
    GetRepository()(MinimalRepositoryable)
    GetRerequestable()(*bool)
    GetRunsRerequestable()(*bool)
    GetStatus()(*CheckSuite_status)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAfter(value *string)()
    SetApp(value NullableIntegrationable)()
    SetBefore(value *string)()
    SetCheckRunsUrl(value *string)()
    SetConclusion(value *CheckSuite_conclusion)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHeadBranch(value *string)()
    SetHeadCommit(value SimpleCommitable)()
    SetHeadSha(value *string)()
    SetId(value *int32)()
    SetLatestCheckRunsCount(value *int32)()
    SetNodeId(value *string)()
    SetPullRequests(value []PullRequestMinimalable)()
    SetRepository(value MinimalRepositoryable)()
    SetRerequestable(value *bool)()
    SetRunsRerequestable(value *bool)()
    SetStatus(value *CheckSuite_status)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
