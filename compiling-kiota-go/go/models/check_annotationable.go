package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckAnnotationable 
type CheckAnnotationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnnotationLevel()(*string)
    GetBlobHref()(*string)
    GetEndColumn()(*int32)
    GetEndLine()(*int32)
    GetMessage()(*string)
    GetPath()(*string)
    GetRawDetails()(*string)
    GetStartColumn()(*int32)
    GetStartLine()(*int32)
    GetTitle()(*string)
    SetAnnotationLevel(value *string)()
    SetBlobHref(value *string)()
    SetEndColumn(value *int32)()
    SetEndLine(value *int32)()
    SetMessage(value *string)()
    SetPath(value *string)()
    SetRawDetails(value *string)()
    SetStartColumn(value *int32)()
    SetStartLine(value *int32)()
    SetTitle(value *string)()
}
