package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationTokenable 
type AuthenticationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPermissions()(AuthenticationToken_permissionsable)
    GetRepositories()([]Repositoryable)
    GetRepositorySelection()(*AuthenticationToken_repository_selection)
    GetSingleFile()(*string)
    GetToken()(*string)
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPermissions(value AuthenticationToken_permissionsable)()
    SetRepositories(value []Repositoryable)()
    SetRepositorySelection(value *AuthenticationToken_repository_selection)()
    SetSingleFile(value *string)()
    SetToken(value *string)()
}
