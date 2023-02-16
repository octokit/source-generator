package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Codespace_git_statusable 
type Codespace_git_statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAhead()(*int32)
    GetBehind()(*int32)
    GetHasUncommittedChanges()(*bool)
    GetHasUnpushedChanges()(*bool)
    GetRef()(*string)
    SetAhead(value *int32)()
    SetBehind(value *int32)()
    SetHasUncommittedChanges(value *bool)()
    SetHasUnpushedChanges(value *bool)()
    SetRef(value *string)()
}
