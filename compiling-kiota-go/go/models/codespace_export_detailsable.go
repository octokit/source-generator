package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespaceExportDetailsable 
type CodespaceExportDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExportUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetSha()(*string)
    GetState()(*string)
    SetBranch(value *string)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExportUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetSha(value *string)()
    SetState(value *string)()
}
