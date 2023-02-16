package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesSecretsItemWithSecret_namePutRequestBodyable 
type CodespacesSecretsItemWithSecret_namePutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptedValue()(*string)
    GetKeyId()(*string)
    GetSelectedRepositoryIds()([]int32)
    SetEncryptedValue(value *string)()
    SetKeyId(value *string)()
    SetSelectedRepositoryIds(value []int32)()
}
