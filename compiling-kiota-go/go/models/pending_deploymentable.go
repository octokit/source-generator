package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingDeploymentable 
type PendingDeploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUserCanApprove()(*bool)
    GetEnvironment()(PendingDeployment_environmentable)
    GetReviewers()([]PendingDeployment_reviewersable)
    GetWaitTimer()(*int32)
    GetWaitTimerStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCurrentUserCanApprove(value *bool)()
    SetEnvironment(value PendingDeployment_environmentable)()
    SetReviewers(value []PendingDeployment_reviewersable)()
    SetWaitTimer(value *int32)()
    SetWaitTimerStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
