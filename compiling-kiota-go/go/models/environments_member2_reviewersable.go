package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentsMember2_reviewersable 
type EnvironmentsMember2_reviewersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewer()(EnvironmentsMember2_reviewers_Environmentsable)
    GetType()(*DeploymentReviewerType)
    SetReviewer(value EnvironmentsMember2_reviewers_Environmentsable)()
    SetType(value *DeploymentReviewerType)()
}
