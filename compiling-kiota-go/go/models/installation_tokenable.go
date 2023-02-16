package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InstallationTokenable 
type InstallationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*string)
    GetHasMultipleSingleFiles()(*bool)
    GetPermissions()(AppPermissionsable)
    GetRepositories()([]Repositoryable)
    GetRepositorySelection()(*InstallationToken_repository_selection)
    GetSingleFile()(*string)
    GetSingleFilePaths()([]string)
    GetToken()(*string)
    SetExpiresAt(value *string)()
    SetHasMultipleSingleFiles(value *bool)()
    SetPermissions(value AppPermissionsable)()
    SetRepositories(value []Repositoryable)()
    SetRepositorySelection(value *InstallationToken_repository_selection)()
    SetSingleFile(value *string)()
    SetSingleFilePaths(value []string)()
    SetToken(value *string)()
}
