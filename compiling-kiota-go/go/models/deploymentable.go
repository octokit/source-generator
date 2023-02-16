package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deploymentable 
type Deploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetOriginalEnvironment()(*string)
    GetPayload()(Deployment_Pending_deploymentsable)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetProductionEnvironment()(*bool)
    GetRef()(*string)
    GetRepositoryUrl()(*string)
    GetSha()(*string)
    GetStatusesUrl()(*string)
    GetTask()(*string)
    GetTransientEnvironment()(*bool)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetOriginalEnvironment(value *string)()
    SetPayload(value Deployment_Pending_deploymentsable)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetProductionEnvironment(value *bool)()
    SetRef(value *string)()
    SetRepositoryUrl(value *string)()
    SetSha(value *string)()
    SetStatusesUrl(value *string)()
    SetTask(value *string)()
    SetTransientEnvironment(value *bool)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
