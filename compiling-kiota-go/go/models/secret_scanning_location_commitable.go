package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocationCommitable 
type SecretScanningLocationCommitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlobSha()(*string)
    GetBlobUrl()(*string)
    GetCommitSha()(*string)
    GetCommitUrl()(*string)
    GetEndColumn()(*int64)
    GetEndLine()(*int64)
    GetPath()(*string)
    GetStartColumn()(*int64)
    GetStartLine()(*int64)
    SetBlobSha(value *string)()
    SetBlobUrl(value *string)()
    SetCommitSha(value *string)()
    SetCommitUrl(value *string)()
    SetEndColumn(value *int64)()
    SetEndLine(value *int64)()
    SetPath(value *string)()
    SetStartColumn(value *int64)()
    SetStartLine(value *int64)()
}
