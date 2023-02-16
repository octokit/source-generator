package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchable 
type ProtectedBranchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(ProtectedBranch_allow_deletionsable)
    GetAllowForcePushes()(ProtectedBranch_allow_force_pushesable)
    GetAllowForkSyncing()(ProtectedBranch_allow_fork_syncingable)
    GetBlockCreations()(ProtectedBranch_block_creationsable)
    GetEnforceAdmins()(ProtectedBranch_enforce_adminsable)
    GetLockBranch()(ProtectedBranch_lock_branchable)
    GetRequiredConversationResolution()(ProtectedBranch_required_conversation_resolutionable)
    GetRequiredLinearHistory()(ProtectedBranch_required_linear_historyable)
    GetRequiredPullRequestReviews()(ProtectedBranch_required_pull_request_reviewsable)
    GetRequiredSignatures()(ProtectedBranch_required_signaturesable)
    GetRequiredStatusChecks()(StatusCheckPolicyable)
    GetRestrictions()(BranchRestrictionPolicyable)
    GetUrl()(*string)
    SetAllowDeletions(value ProtectedBranch_allow_deletionsable)()
    SetAllowForcePushes(value ProtectedBranch_allow_force_pushesable)()
    SetAllowForkSyncing(value ProtectedBranch_allow_fork_syncingable)()
    SetBlockCreations(value ProtectedBranch_block_creationsable)()
    SetEnforceAdmins(value ProtectedBranch_enforce_adminsable)()
    SetLockBranch(value ProtectedBranch_lock_branchable)()
    SetRequiredConversationResolution(value ProtectedBranch_required_conversation_resolutionable)()
    SetRequiredLinearHistory(value ProtectedBranch_required_linear_historyable)()
    SetRequiredPullRequestReviews(value ProtectedBranch_required_pull_request_reviewsable)()
    SetRequiredSignatures(value ProtectedBranch_required_signaturesable)()
    SetRequiredStatusChecks(value StatusCheckPolicyable)()
    SetRestrictions(value BranchRestrictionPolicyable)()
    SetUrl(value *string)()
}
