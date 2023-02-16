package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependencyGraphDiffable 
type DependencyGraphDiffable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeType()(*DependencyGraphDiff_change_type)
    GetEcosystem()(*string)
    GetLicense()(*string)
    GetManifest()(*string)
    GetName()(*string)
    GetPackageUrl()(*string)
    GetScope()(*DependencyGraphDiff_scope)
    GetSourceRepositoryUrl()(*string)
    GetVersion()(*string)
    GetVulnerabilities()([]DependencyGraphDiff_vulnerabilitiesable)
    SetChangeType(value *DependencyGraphDiff_change_type)()
    SetEcosystem(value *string)()
    SetLicense(value *string)()
    SetManifest(value *string)()
    SetName(value *string)()
    SetPackageUrl(value *string)()
    SetScope(value *DependencyGraphDiff_scope)()
    SetSourceRepositoryUrl(value *string)()
    SetVersion(value *string)()
    SetVulnerabilities(value []DependencyGraphDiff_vulnerabilitiesable)()
}
