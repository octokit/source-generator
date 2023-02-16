package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowUsage_billableable 
type WorkflowUsage_billableable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMACOS()(WorkflowUsage_billable_MACOSable)
    GetUBUNTU()(WorkflowUsage_billable_UBUNTUable)
    GetWINDOWS()(WorkflowUsage_billable_WINDOWSable)
    SetMACOS(value WorkflowUsage_billable_MACOSable)()
    SetUBUNTU(value WorkflowUsage_billable_UBUNTUable)()
    SetWINDOWS(value WorkflowUsage_billable_WINDOWSable)()
}
