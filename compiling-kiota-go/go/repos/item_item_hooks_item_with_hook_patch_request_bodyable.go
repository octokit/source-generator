package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemHooksItemWithHook_PatchRequestBodyable 
type ItemItemHooksItemWithHook_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetAddEvents()([]string)
    GetConfig()(ItemItemHooksItemWithHook_PatchRequestBody_configable)
    GetEvents()([]string)
    GetRemoveEvents()([]string)
    SetActive(value *bool)()
    SetAddEvents(value []string)()
    SetConfig(value ItemItemHooksItemWithHook_PatchRequestBody_configable)()
    SetEvents(value []string)()
    SetRemoveEvents(value []string)()
}
