package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchRequiredStatusCheckable 
type ProtectedBranchRequiredStatusCheckable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]ProtectedBranchRequiredStatusCheck_checksable)
    GetContexts()([]string)
    GetContextsUrl()(*string)
    GetEnforcementLevel()(*string)
    GetStrict()(*bool)
    GetUrl()(*string)
    SetChecks(value []ProtectedBranchRequiredStatusCheck_checksable)()
    SetContexts(value []string)()
    SetContextsUrl(value *string)()
    SetEnforcementLevel(value *string)()
    SetStrict(value *bool)()
    SetUrl(value *string)()
}
