package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplacePurchase_marketplace_purchaseable 
type MarketplacePurchase_marketplace_purchaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillingCycle()(*string)
    GetFreeTrialEndsOn()(*string)
    GetIsInstalled()(*bool)
    GetNextBillingDate()(*string)
    GetOnFreeTrial()(*bool)
    GetPlan()(MarketplaceListingPlanable)
    GetUnitCount()(*int32)
    GetUpdatedAt()(*string)
    SetBillingCycle(value *string)()
    SetFreeTrialEndsOn(value *string)()
    SetIsInstalled(value *bool)()
    SetNextBillingDate(value *string)()
    SetOnFreeTrial(value *bool)()
    SetPlan(value MarketplaceListingPlanable)()
    SetUnitCount(value *int32)()
    SetUpdatedAt(value *string)()
}
