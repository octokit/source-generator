package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsRequired_workflowsResponseable 
type ItemActionsRequired_workflowsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequiredWorkflows()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable)
    GetTotalCount()(*int32)
    SetRequiredWorkflows(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RequiredWorkflowable)()
    SetTotalCount(value *int32)()
}
