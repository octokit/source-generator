package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentStatusable 
type DeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDeploymentUrl()(*string)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetEnvironmentUrl()(*string)
    GetId()(*int32)
    GetLogUrl()(*string)
    GetNodeId()(*string)
    GetPerformedViaGithubApp()(NullableIntegrationable)
    GetRepositoryUrl()(*string)
    GetState()(*DeploymentStatus_state)
    GetTargetUrl()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDeploymentUrl(value *string)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetEnvironmentUrl(value *string)()
    SetId(value *int32)()
    SetLogUrl(value *string)()
    SetNodeId(value *string)()
    SetPerformedViaGithubApp(value NullableIntegrationable)()
    SetRepositoryUrl(value *string)()
    SetState(value *DeploymentStatus_state)()
    SetTargetUrl(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
