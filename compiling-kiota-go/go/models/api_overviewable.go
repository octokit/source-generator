package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiOverviewable 
type ApiOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    GetApi()([]string)
    GetDependabot()([]string)
    GetGit()([]string)
    GetHooks()([]string)
    GetImporter()([]string)
    GetPackages()([]string)
    GetPages()([]string)
    GetSshKeyFingerprints()(ApiOverview_ssh_key_fingerprintsable)
    GetSshKeys()([]string)
    GetVerifiablePasswordAuthentication()(*bool)
    GetWeb()([]string)
    SetActions(value []string)()
    SetApi(value []string)()
    SetDependabot(value []string)()
    SetGit(value []string)()
    SetHooks(value []string)()
    SetImporter(value []string)()
    SetPackages(value []string)()
    SetPages(value []string)()
    SetSshKeyFingerprints(value ApiOverview_ssh_key_fingerprintsable)()
    SetSshKeys(value []string)()
    SetVerifiablePasswordAuthentication(value *bool)()
    SetWeb(value []string)()
}
