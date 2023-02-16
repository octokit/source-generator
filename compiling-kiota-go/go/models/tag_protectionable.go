package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TagProtectionable 
type TagProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetEnabled()(*bool)
    GetId()(*int32)
    GetPattern()(*string)
    GetUpdatedAt()(*string)
    SetCreatedAt(value *string)()
    SetEnabled(value *bool)()
    SetId(value *int32)()
    SetPattern(value *string)()
    SetUpdatedAt(value *string)()
}
