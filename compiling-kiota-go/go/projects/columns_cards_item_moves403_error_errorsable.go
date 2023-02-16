package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemMoves403Error_errorsable 
type ColumnsCardsItemMoves403Error_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetField()(*string)
    GetMessage()(*string)
    GetResource()(*string)
    SetCode(value *string)()
    SetField(value *string)()
    SetMessage(value *string)()
    SetResource(value *string)()
}
