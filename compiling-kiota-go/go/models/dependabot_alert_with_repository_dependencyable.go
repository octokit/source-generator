package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertWithRepository_dependencyable 
type DependabotAlertWithRepository_dependencyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifestPath()(*string)
    GetPackage()(DependabotAlertPackageable)
    GetScope()(*DependabotAlertWithRepository_dependency_scope)
    SetManifestPath(value *string)()
    SetPackage(value DependabotAlertPackageable)()
    SetScope(value *DependabotAlertWithRepository_dependency_scope)()
}
