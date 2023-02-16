package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable 
type ProtectedBranchPullRequestReview_bypass_pull_request_allowancesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]Integrationable)
    GetTeams()([]Teamable)
    GetUsers()([]SimpleUserable)
    SetApps(value []Integrationable)()
    SetTeams(value []Teamable)()
    SetUsers(value []SimpleUserable)()
}
