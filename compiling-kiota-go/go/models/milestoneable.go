package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Milestoneable 
type Milestoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetClosedIssues()(*int32)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetDueOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLabelsUrl()(*string)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetOpenIssues()(*int32)
    GetState()(*Milestone_state)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetClosedIssues(value *int32)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetDueOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLabelsUrl(value *string)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetOpenIssues(value *int32)()
    SetState(value *Milestone_state)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
