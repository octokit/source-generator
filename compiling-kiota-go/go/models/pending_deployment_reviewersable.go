package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingDeployment_reviewersable 
type PendingDeployment_reviewersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewer()(PendingDeployment_reviewers_Pending_deploymentsable)
    GetType()(*DeploymentReviewerType)
    SetReviewer(value PendingDeployment_reviewers_Pending_deploymentsable)()
    SetType(value *DeploymentReviewerType)()
}
