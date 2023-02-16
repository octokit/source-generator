package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesDevcontainersResponseable 
type ItemItemCodespacesDevcontainersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevcontainers()([]ItemItemCodespacesDevcontainersResponse_devcontainersable)
    GetTotalCount()(*int32)
    SetDevcontainers(value []ItemItemCodespacesDevcontainersResponse_devcontainersable)()
    SetTotalCount(value *int32)()
}
