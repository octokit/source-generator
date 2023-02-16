package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemWithProject_PatchRequestBodyable 
type ItemWithProject_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetName()(*string)
    GetPrivate()(*bool)
    GetState()(*string)
    SetBody(value *string)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetState(value *string)()
}
