package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamsPostRequestBodyable 
type ItemTeamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetMaintainers()([]string)
    GetName()(*string)
    GetParentTeamId()(*int32)
    GetRepoNames()([]string)
    SetDescription(value *string)()
    SetMaintainers(value []string)()
    SetName(value *string)()
    SetParentTeamId(value *int32)()
    SetRepoNames(value []string)()
}
