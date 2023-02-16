package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplacePurchase_marketplace_pending_changeable 
type MarketplacePurchase_marketplace_pending_changeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEffectiveDate()(*string)
    GetId()(*int32)
    GetIsInstalled()(*bool)
    GetPlan()(MarketplaceListingPlanable)
    GetUnitCount()(*int32)
    SetEffectiveDate(value *string)()
    SetId(value *int32)()
    SetIsInstalled(value *bool)()
    SetPlan(value MarketplaceListingPlanable)()
    SetUnitCount(value *int32)()
}
