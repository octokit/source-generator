package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CombinedBillingUsageable 
type CombinedBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaysLeftInBillingCycle()(*int32)
    GetEstimatedPaidStorageForMonth()(*int32)
    GetEstimatedStorageForMonth()(*int32)
    SetDaysLeftInBillingCycle(value *int32)()
    SetEstimatedPaidStorageForMonth(value *int32)()
    SetEstimatedStorageForMonth(value *int32)()
}
