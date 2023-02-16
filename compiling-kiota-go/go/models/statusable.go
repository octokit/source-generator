package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Statusable 
type Statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetContext()(*string)
    GetCreatedAt()(*string)
    GetCreator()(NullableSimpleUserable)
    GetDescription()(*string)
    GetId()(*int32)
    GetNodeId()(*string)
    GetState()(*string)
    GetTargetUrl()(*string)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetContext(value *string)()
    SetCreatedAt(value *string)()
    SetCreator(value NullableSimpleUserable)()
    SetDescription(value *string)()
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetState(value *string)()
    SetTargetUrl(value *string)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
