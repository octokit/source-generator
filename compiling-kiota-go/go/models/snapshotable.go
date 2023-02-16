package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Snapshotable 
type Snapshotable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetector()(Snapshot_detectorable)
    GetJob()(Snapshot_jobable)
    GetManifests()(Snapshot_manifestsable)
    GetMetadata()(Metadataable)
    GetRef()(*string)
    GetScanned()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSha()(*string)
    GetVersion()(*int32)
    SetDetector(value Snapshot_detectorable)()
    SetJob(value Snapshot_jobable)()
    SetManifests(value Snapshot_manifestsable)()
    SetMetadata(value Metadataable)()
    SetRef(value *string)()
    SetScanned(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSha(value *string)()
    SetVersion(value *int32)()
}
