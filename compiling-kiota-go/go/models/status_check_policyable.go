package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusCheckPolicyable 
type StatusCheckPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]StatusCheckPolicy_checksable)
    GetContexts()([]string)
    GetContextsUrl()(*string)
    GetStrict()(*bool)
    GetUrl()(*string)
    SetChecks(value []StatusCheckPolicy_checksable)()
    SetContexts(value []string)()
    SetContextsUrl(value *string)()
    SetStrict(value *bool)()
    SetUrl(value *string)()
}
