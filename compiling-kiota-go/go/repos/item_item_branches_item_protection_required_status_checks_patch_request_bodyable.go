package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRequired_status_checksPatchRequestBodyable 
type ItemItemBranchesItemProtectionRequired_status_checksPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]ItemItemBranchesItemProtectionRequired_status_checksPatchRequestBody_checksable)
    GetContexts()([]string)
    GetStrict()(*bool)
    SetChecks(value []ItemItemBranchesItemProtectionRequired_status_checksPatchRequestBody_checksable)()
    SetContexts(value []string)()
    SetStrict(value *bool)()
}
