package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBodyable 
type ItemItemBranchesItemProtectionPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(*bool)
    GetAllowForcePushes()(*bool)
    GetAllowForkSyncing()(*bool)
    GetBlockCreations()(*bool)
    GetEnforceAdmins()(*bool)
    GetLockBranch()(*bool)
    GetRequiredConversationResolution()(*bool)
    GetRequiredLinearHistory()(*bool)
    GetRequiredPullRequestReviews()(ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable)
    GetRequiredStatusChecks()(ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable)
    GetRestrictions()(ItemItemBranchesItemProtectionPutRequestBody_restrictionsable)
    SetAllowDeletions(value *bool)()
    SetAllowForcePushes(value *bool)()
    SetAllowForkSyncing(value *bool)()
    SetBlockCreations(value *bool)()
    SetEnforceAdmins(value *bool)()
    SetLockBranch(value *bool)()
    SetRequiredConversationResolution(value *bool)()
    SetRequiredLinearHistory(value *bool)()
    SetRequiredPullRequestReviews(value ItemItemBranchesItemProtectionPutRequestBody_required_pull_request_reviewsable)()
    SetRequiredStatusChecks(value ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable)()
    SetRestrictions(value ItemItemBranchesItemProtectionPutRequestBody_restrictionsable)()
}
