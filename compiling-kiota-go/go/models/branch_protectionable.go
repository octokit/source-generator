package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchProtectionable 
type BranchProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeletions()(BranchProtection_allow_deletionsable)
    GetAllowForcePushes()(BranchProtection_allow_force_pushesable)
    GetAllowForkSyncing()(BranchProtection_allow_fork_syncingable)
    GetBlockCreations()(BranchProtection_block_creationsable)
    GetEnabled()(*bool)
    GetEnforceAdmins()(ProtectedBranchAdminEnforcedable)
    GetLockBranch()(BranchProtection_lock_branchable)
    GetName()(*string)
    GetProtectionUrl()(*string)
    GetRequiredConversationResolution()(BranchProtection_required_conversation_resolutionable)
    GetRequiredLinearHistory()(BranchProtection_required_linear_historyable)
    GetRequiredPullRequestReviews()(ProtectedBranchPullRequestReviewable)
    GetRequiredSignatures()(BranchProtection_required_signaturesable)
    GetRequiredStatusChecks()(ProtectedBranchRequiredStatusCheckable)
    GetRestrictions()(BranchRestrictionPolicyable)
    GetUrl()(*string)
    SetAllowDeletions(value BranchProtection_allow_deletionsable)()
    SetAllowForcePushes(value BranchProtection_allow_force_pushesable)()
    SetAllowForkSyncing(value BranchProtection_allow_fork_syncingable)()
    SetBlockCreations(value BranchProtection_block_creationsable)()
    SetEnabled(value *bool)()
    SetEnforceAdmins(value ProtectedBranchAdminEnforcedable)()
    SetLockBranch(value BranchProtection_lock_branchable)()
    SetName(value *string)()
    SetProtectionUrl(value *string)()
    SetRequiredConversationResolution(value BranchProtection_required_conversation_resolutionable)()
    SetRequiredLinearHistory(value BranchProtection_required_linear_historyable)()
    SetRequiredPullRequestReviews(value ProtectedBranchPullRequestReviewable)()
    SetRequiredSignatures(value BranchProtection_required_signaturesable)()
    SetRequiredStatusChecks(value ProtectedBranchRequiredStatusCheckable)()
    SetRestrictions(value BranchRestrictionPolicyable)()
    SetUrl(value *string)()
}
