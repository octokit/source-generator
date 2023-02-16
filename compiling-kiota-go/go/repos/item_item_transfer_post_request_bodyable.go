package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemTransferPostRequestBodyable 
type ItemItemTransferPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNewName()(*string)
    GetNewOwner()(*string)
    GetTeamIds()([]int32)
    SetNewName(value *string)()
    SetNewOwner(value *string)()
    SetTeamIds(value []int32)()
}
