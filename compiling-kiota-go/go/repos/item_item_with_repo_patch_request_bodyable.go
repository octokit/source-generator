package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemWithRepoPatchRequestBodyable 
type ItemItemWithRepoPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowForking()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetAllowUpdateBranch()(*bool)
    GetArchived()(*bool)
    GetDefaultBranch()(*string)
    GetDeleteBranchOnMerge()(*bool)
    GetDescription()(*string)
    GetHasIssues()(*bool)
    GetHasProjects()(*bool)
    GetHasWiki()(*bool)
    GetHomepage()(*string)
    GetIsTemplate()(*bool)
    GetName()(*string)
    GetPrivate()(*bool)
    GetSecurityAndAnalysis()(ItemItemWithRepoPatchRequestBody_security_and_analysisable)
    GetUseSquashPrTitleAsDefault()(*bool)
    GetWebCommitSignoffRequired()(*bool)
    SetAllowAutoMerge(value *bool)()
    SetAllowForking(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetAllowUpdateBranch(value *bool)()
    SetArchived(value *bool)()
    SetDefaultBranch(value *string)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDescription(value *string)()
    SetHasIssues(value *bool)()
    SetHasProjects(value *bool)()
    SetHasWiki(value *bool)()
    SetHomepage(value *string)()
    SetIsTemplate(value *bool)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetSecurityAndAnalysis(value ItemItemWithRepoPatchRequestBody_security_and_analysisable)()
    SetUseSquashPrTitleAsDefault(value *bool)()
    SetWebCommitSignoffRequired(value *bool)()
}
