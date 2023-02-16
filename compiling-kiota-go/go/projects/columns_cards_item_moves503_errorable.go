package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemMoves503Errorable 
type ColumnsCardsItemMoves503Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDocumentationUrl()(*string)
    GetErrors()([]ColumnsCardsItemMoves503Error_errorsable)
    GetMessage()(*string)
    SetCode(value *string)()
    SetDocumentationUrl(value *string)()
    SetErrors(value []ColumnsCardsItemMoves503Error_errorsable)()
    SetMessage(value *string)()
}
