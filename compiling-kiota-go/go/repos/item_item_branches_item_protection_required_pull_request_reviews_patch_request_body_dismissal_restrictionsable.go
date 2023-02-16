package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable 
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBody_dismissal_restrictionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]string)
    GetTeams()([]string)
    GetUsers()([]string)
    SetApps(value []string)()
    SetTeams(value []string)()
    SetUsers(value []string)()
}
