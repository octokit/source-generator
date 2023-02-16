package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowRunUsage_billable_UBUNTUable 
type WorkflowRunUsage_billable_UBUNTUable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetJobRuns()([]WorkflowRunUsage_billable_UBUNTU_job_runsable)
    GetJobs()(*int32)
    GetTotalMs()(*int32)
    SetJobRuns(value []WorkflowRunUsage_billable_UBUNTU_job_runsable)()
    SetJobs(value *int32)()
    SetTotalMs(value *int32)()
}
