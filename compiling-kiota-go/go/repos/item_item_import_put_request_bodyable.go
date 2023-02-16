package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemImportPutRequestBodyable 
type ItemItemImportPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTfvcProject()(*string)
    GetVcsPassword()(*string)
    GetVcsUrl()(*string)
    GetVcsUsername()(*string)
    SetTfvcProject(value *string)()
    SetVcsPassword(value *string)()
    SetVcsUrl(value *string)()
    SetVcsUsername(value *string)()
}
