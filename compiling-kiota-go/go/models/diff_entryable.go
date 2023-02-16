package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DiffEntryable 
type DiffEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditions()(*int32)
    GetBlobUrl()(*string)
    GetChanges()(*int32)
    GetContentsUrl()(*string)
    GetDeletions()(*int32)
    GetFilename()(*string)
    GetPatch()(*string)
    GetPreviousFilename()(*string)
    GetRawUrl()(*string)
    GetSha()(*string)
    GetStatus()(*DiffEntry_status)
    SetAdditions(value *int32)()
    SetBlobUrl(value *string)()
    SetChanges(value *int32)()
    SetContentsUrl(value *string)()
    SetDeletions(value *int32)()
    SetFilename(value *string)()
    SetPatch(value *string)()
    SetPreviousFilename(value *string)()
    SetRawUrl(value *string)()
    SetSha(value *string)()
    SetStatus(value *DiffEntry_status)()
}
