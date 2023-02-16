package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeownersErrors_errorsable 
type CodeownersErrors_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumn()(*int32)
    GetKind()(*string)
    GetLine()(*int32)
    GetMessage()(*string)
    GetPath()(*string)
    GetSource()(*string)
    GetSuggestion()(*string)
    SetColumn(value *int32)()
    SetKind(value *string)()
    SetLine(value *int32)()
    SetMessage(value *string)()
    SetPath(value *string)()
    SetSource(value *string)()
    SetSuggestion(value *string)()
}
