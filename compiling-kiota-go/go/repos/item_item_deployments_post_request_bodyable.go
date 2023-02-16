package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDeploymentsPostRequestBodyable 
type ItemItemDeploymentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoMerge()(*bool)
    GetDescription()(*string)
    GetEnvironment()(*string)
    GetPayload()(ItemItemDeploymentsPostRequestBody_Deploymentsable)
    GetProductionEnvironment()(*bool)
    GetRef()(*string)
    GetRequiredContexts()([]string)
    GetTask()(*string)
    GetTransientEnvironment()(*bool)
    SetAutoMerge(value *bool)()
    SetDescription(value *string)()
    SetEnvironment(value *string)()
    SetPayload(value ItemItemDeploymentsPostRequestBody_Deploymentsable)()
    SetProductionEnvironment(value *bool)()
    SetRef(value *string)()
    SetRequiredContexts(value []string)()
    SetTask(value *string)()
    SetTransientEnvironment(value *bool)()
}
