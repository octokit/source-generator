package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PackagesBillingUsageable 
type PackagesBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedGigabytesBandwidth()(*int32)
    GetTotalGigabytesBandwidthUsed()(*int32)
    GetTotalPaidGigabytesBandwidthUsed()(*int32)
    SetIncludedGigabytesBandwidth(value *int32)()
    SetTotalGigabytesBandwidthUsed(value *int32)()
    SetTotalPaidGigabytesBandwidthUsed(value *int32)()
}
