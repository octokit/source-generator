package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemMoves403Errorable 
type ColumnsCardsItemMoves403Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDocumentationUrl()(*string)
    GetErrors()([]ColumnsCardsItemMoves403Error_errorsable)
    GetMessage()(*string)
    SetDocumentationUrl(value *string)()
    SetErrors(value []ColumnsCardsItemMoves403Error_errorsable)()
    SetMessage(value *string)()
}
