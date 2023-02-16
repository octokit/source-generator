package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidationError_errorsable 
type ValidationError_errorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetField()(*string)
    GetIndex()(*int32)
    GetMessage()(*string)
    GetResource()(*string)
    GetValue()(ValidationError_errors_Gistsable)
    SetCode(value *string)()
    SetField(value *string)()
    SetIndex(value *int32)()
    SetMessage(value *string)()
    SetResource(value *string)()
    SetValue(value ValidationError_errors_Gistsable)()
}
