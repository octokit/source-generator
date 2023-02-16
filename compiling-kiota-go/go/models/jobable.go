package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Jobable 
type Jobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRunUrl()(*string)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConclusion()(*Job_conclusion)
    GetHeadBranch()(*string)
    GetHeadSha()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLabels()([]string)
    GetName()(*string)
    GetNodeId()(*string)
    GetRunAttempt()(*int32)
    GetRunId()(*int32)
    GetRunnerGroupId()(*int32)
    GetRunnerGroupName()(*string)
    GetRunnerId()(*int32)
    GetRunnerName()(*string)
    GetRunUrl()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*Job_status)
    GetSteps()([]Job_stepsable)
    GetUrl()(*string)
    GetWorkflowName()(*string)
    SetCheckRunUrl(value *string)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConclusion(value *Job_conclusion)()
    SetHeadBranch(value *string)()
    SetHeadSha(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLabels(value []string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetRunAttempt(value *int32)()
    SetRunId(value *int32)()
    SetRunnerGroupId(value *int32)()
    SetRunnerGroupName(value *string)()
    SetRunnerId(value *int32)()
    SetRunnerName(value *string)()
    SetRunUrl(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *Job_status)()
    SetSteps(value []Job_stepsable)()
    SetUrl(value *string)()
    SetWorkflowName(value *string)()
}
