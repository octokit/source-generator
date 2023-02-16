package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContributorActivityable 
type ContributorActivityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetTotal()(*int32)
    GetWeeks()([]ContributorActivity_weeksable)
    SetAuthor(value NullableSimpleUserable)()
    SetTotal(value *int32)()
    SetWeeks(value []ContributorActivity_weeksable)()
}
