package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MigrationsPostRequestBodyable 
type MigrationsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeAttachments()(*bool)
    GetExcludeGitData()(*bool)
    GetExcludeMetadata()(*bool)
    GetExcludeOwnerProjects()(*bool)
    GetExcludeReleases()(*bool)
    GetLockRepositories()(*bool)
    GetOrgMetadataOnly()(*bool)
    GetRepositories()([]string)
    SetExcludeAttachments(value *bool)()
    SetExcludeGitData(value *bool)()
    SetExcludeMetadata(value *bool)()
    SetExcludeOwnerProjects(value *bool)()
    SetExcludeReleases(value *bool)()
    SetLockRepositories(value *bool)()
    SetOrgMetadataOnly(value *bool)()
    SetRepositories(value []string)()
}
