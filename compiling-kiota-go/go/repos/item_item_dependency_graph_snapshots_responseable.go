package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDependencyGraphSnapshotsResponseable 
type ItemItemDependencyGraphSnapshotsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetId()(*int32)
    GetMessage()(*string)
    GetResult()(*string)
    SetCreatedAt(value *string)()
    SetId(value *int32)()
    SetMessage(value *string)()
    SetResult(value *string)()
}
