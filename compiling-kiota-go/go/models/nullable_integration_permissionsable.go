package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableIntegration_permissionsable 
type NullableIntegration_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()(*string)
    GetContents()(*string)
    GetDeployments()(*string)
    GetIssues()(*string)
    GetMetadata()(*string)
    SetChecks(value *string)()
    SetContents(value *string)()
    SetDeployments(value *string)()
    SetIssues(value *string)()
    SetMetadata(value *string)()
}
