package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Runnerable 
type Runnerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusy()(*bool)
    GetId()(*int32)
    GetLabels()([]RunnerLabelable)
    GetName()(*string)
    GetOs()(*string)
    GetStatus()(*string)
    SetBusy(value *bool)()
    SetId(value *int32)()
    SetLabels(value []RunnerLabelable)()
    SetName(value *string)()
    SetOs(value *string)()
    SetStatus(value *string)()
}
