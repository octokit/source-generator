package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable 
type ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)
    GetContexts()([]string)
    GetStrict()(*bool)
    SetChecks(value []ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)()
    SetContexts(value []string)()
    SetStrict(value *bool)()
}
