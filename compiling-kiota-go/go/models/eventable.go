package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Eventable 
type Eventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActor()(Actorable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetOrg()(Actorable)
    GetPayload()(Event_payloadable)
    GetPublic()(*bool)
    GetRepo()(Event_repoable)
    GetType()(*string)
    SetActor(value Actorable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetOrg(value Actorable)()
    SetPayload(value Event_payloadable)()
    SetPublic(value *bool)()
    SetRepo(value Event_repoable)()
    SetType(value *string)()
}
