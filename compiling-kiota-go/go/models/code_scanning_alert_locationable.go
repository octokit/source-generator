package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAlertLocationable 
type CodeScanningAlertLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndColumn()(*int32)
    GetEndLine()(*int32)
    GetPath()(*string)
    GetStartColumn()(*int32)
    GetStartLine()(*int32)
    SetEndColumn(value *int32)()
    SetEndLine(value *int32)()
    SetPath(value *string)()
    SetStartColumn(value *int32)()
    SetStartLine(value *int32)()
}
