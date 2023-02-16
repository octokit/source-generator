package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckRunable 
type CheckRunable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(NullableIntegrationable)
    GetCheckSuite()(CheckRun_check_suiteable)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConclusion()(*CheckRun_conclusion)
    GetDeployment()(DeploymentSimpleable)
    GetDetailsUrl()(*string)
    GetExternalId()(*string)
    GetHeadSha()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetOutput()(CheckRun_outputable)
    GetPullRequests()([]PullRequestMinimalable)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*CheckRun_status)
    GetUrl()(*string)
    SetApp(value NullableIntegrationable)()
    SetCheckSuite(value CheckRun_check_suiteable)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConclusion(value *CheckRun_conclusion)()
    SetDeployment(value DeploymentSimpleable)()
    SetDetailsUrl(value *string)()
    SetExternalId(value *string)()
    SetHeadSha(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetOutput(value CheckRun_outputable)()
    SetPullRequests(value []PullRequestMinimalable)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *CheckRun_status)()
    SetUrl(value *string)()
}
