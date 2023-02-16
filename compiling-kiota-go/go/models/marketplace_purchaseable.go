package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplacePurchaseable 
type MarketplacePurchaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetMarketplacePendingChange()(MarketplacePurchase_marketplace_pending_changeable)
    GetMarketplacePurchase()(MarketplacePurchase_marketplace_purchaseable)
    GetOrganizationBillingEmail()(*string)
    GetType()(*string)
    GetUrl()(*string)
    SetEmail(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetMarketplacePendingChange(value MarketplacePurchase_marketplace_pending_changeable)()
    SetMarketplacePurchase(value MarketplacePurchase_marketplace_purchaseable)()
    SetOrganizationBillingEmail(value *string)()
    SetType(value *string)()
    SetUrl(value *string)()
}
