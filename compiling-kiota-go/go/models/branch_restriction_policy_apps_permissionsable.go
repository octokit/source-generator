package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchRestrictionPolicy_apps_permissionsable 
type BranchRestrictionPolicy_apps_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContents()(*string)
    GetIssues()(*string)
    GetMetadata()(*string)
    GetSingleFile()(*string)
    SetContents(value *string)()
    SetIssues(value *string)()
    SetMetadata(value *string)()
    SetSingleFile(value *string)()
}
