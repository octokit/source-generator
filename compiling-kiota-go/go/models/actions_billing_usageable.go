package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsBillingUsageable 
type ActionsBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedMinutes()(*int32)
    GetMinutesUsedBreakdown()(ActionsBillingUsage_minutes_used_breakdownable)
    GetTotalMinutesUsed()(*int32)
    GetTotalPaidMinutesUsed()(*int32)
    SetIncludedMinutes(value *int32)()
    SetMinutesUsedBreakdown(value ActionsBillingUsage_minutes_used_breakdownable)()
    SetTotalMinutesUsed(value *int32)()
    SetTotalPaidMinutesUsed(value *int32)()
}
