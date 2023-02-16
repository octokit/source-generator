package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentSimpleable 
type DeploymentSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetOriginalEnvironment()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetProductionEnvironment()(*bool)
    GetRepositoryUrl()(*string)
    GetStatusesUrl()(*string)
    GetTask()(*string)
    GetTransientEnvironment()(*bool)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetOriginalEnvironment(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetProductionEnvironment(value *bool)()
    SetRepositoryUrl(value *string)()
    SetStatusesUrl(value *string)()
    SetTask(value *string)()
    SetTransientEnvironment(value *bool)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
