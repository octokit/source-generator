package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemHooksPostRequestBodyable 
type ItemItemHooksPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetConfig()(ItemItemHooksPostRequestBody_configable)
    GetEvents()([]string)
    GetName()(*string)
    SetActive(value *bool)()
    SetConfig(value ItemItemHooksPostRequestBody_configable)()
    SetEvents(value []string)()
    SetName(value *string)()
}
