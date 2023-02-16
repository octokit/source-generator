package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDeploymentsItemStatusesPostRequestBodyable 
type ItemItemDeploymentsItemStatusesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoInactive()(*bool)
    GetDescription()(*string)
    GetEnvironmentUrl()(*string)
    GetLogUrl()(*string)
    GetTargetUrl()(*string)
    SetAutoInactive(value *bool)()
    SetDescription(value *string)()
    SetEnvironmentUrl(value *string)()
    SetLogUrl(value *string)()
    SetTargetUrl(value *string)()
}
