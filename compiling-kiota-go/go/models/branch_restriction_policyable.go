package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchRestrictionPolicyable 
type BranchRestrictionPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]BranchRestrictionPolicy_appsable)
    GetAppsUrl()(*string)
    GetTeams()([]BranchRestrictionPolicy_teamsable)
    GetTeamsUrl()(*string)
    GetUrl()(*string)
    GetUsers()([]BranchRestrictionPolicy_usersable)
    GetUsersUrl()(*string)
    SetApps(value []BranchRestrictionPolicy_appsable)()
    SetAppsUrl(value *string)()
    SetTeams(value []BranchRestrictionPolicy_teamsable)()
    SetTeamsUrl(value *string)()
    SetUrl(value *string)()
    SetUsers(value []BranchRestrictionPolicy_usersable)()
    SetUsersUrl(value *string)()
}
