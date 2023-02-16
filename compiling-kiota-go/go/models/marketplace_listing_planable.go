package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplaceListingPlanable 
type MarketplaceListingPlanable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountsUrl()(*string)
    GetBullets()([]string)
    GetDescription()(*string)
    GetHasFreeTrial()(*bool)
    GetId()(*int32)
    GetMonthlyPriceInCents()(*int32)
    GetName()(*string)
    GetNumber()(*int32)
    GetPriceModel()(*MarketplaceListingPlan_price_model)
    GetState()(*string)
    GetUnitName()(*string)
    GetUrl()(*string)
    GetYearlyPriceInCents()(*int32)
    SetAccountsUrl(value *string)()
    SetBullets(value []string)()
    SetDescription(value *string)()
    SetHasFreeTrial(value *bool)()
    SetId(value *int32)()
    SetMonthlyPriceInCents(value *int32)()
    SetName(value *string)()
    SetNumber(value *int32)()
    SetPriceModel(value *MarketplaceListingPlan_price_model)()
    SetState(value *string)()
    SetUnitName(value *string)()
    SetUrl(value *string)()
    SetYearlyPriceInCents(value *int32)()
}
