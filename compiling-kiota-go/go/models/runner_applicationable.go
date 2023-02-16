package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunnerApplicationable 
type RunnerApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchitecture()(*string)
    GetDownloadUrl()(*string)
    GetFilename()(*string)
    GetOs()(*string)
    GetSha256Checksum()(*string)
    GetTempDownloadToken()(*string)
    SetArchitecture(value *string)()
    SetDownloadUrl(value *string)()
    SetFilename(value *string)()
    SetOs(value *string)()
    SetSha256Checksum(value *string)()
    SetTempDownloadToken(value *string)()
}
