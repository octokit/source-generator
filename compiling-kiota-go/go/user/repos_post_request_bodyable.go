package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReposPostRequestBodyable 
type ReposPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetAutoInit()(*bool)
    GetDeleteBranchOnMerge()(*bool)
    GetDescription()(*string)
    GetGitignoreTemplate()(*string)
    GetHasDiscussions()(*bool)
    GetHasDownloads()(*bool)
    GetHasIssues()(*bool)
    GetHasProjects()(*bool)
    GetHasWiki()(*bool)
    GetHomepage()(*string)
    GetIsTemplate()(*bool)
    GetLicenseTemplate()(*string)
    GetName()(*string)
    GetPrivate()(*bool)
    GetTeamId()(*int32)
    SetAllowAutoMerge(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetAutoInit(value *bool)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDescription(value *string)()
    SetGitignoreTemplate(value *string)()
    SetHasDiscussions(value *bool)()
    SetHasDownloads(value *bool)()
    SetHasIssues(value *bool)()
    SetHasProjects(value *bool)()
    SetHasWiki(value *bool)()
    SetHomepage(value *string)()
    SetIsTemplate(value *bool)()
    SetLicenseTemplate(value *string)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetTeamId(value *int32)()
}
