package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GpgKey_subkeysable 
type GpgKey_subkeysable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanCertify()(*bool)
    GetCanEncryptComms()(*bool)
    GetCanEncryptStorage()(*bool)
    GetCanSign()(*bool)
    GetCreatedAt()(*string)
    GetEmails()([]GpgKey_subkeys_emailsable)
    GetExpiresAt()(*string)
    GetId()(*int32)
    GetKeyId()(*string)
    GetPrimaryKeyId()(*int32)
    GetPublicKey()(*string)
    GetRawKey()(*string)
    GetRevoked()(*bool)
    SetCanCertify(value *bool)()
    SetCanEncryptComms(value *bool)()
    SetCanEncryptStorage(value *bool)()
    SetCanSign(value *bool)()
    SetCreatedAt(value *string)()
    SetEmails(value []GpgKey_subkeys_emailsable)()
    SetExpiresAt(value *string)()
    SetId(value *int32)()
    SetKeyId(value *string)()
    SetPrimaryKeyId(value *int32)()
    SetPublicKey(value *string)()
    SetRawKey(value *string)()
    SetRevoked(value *bool)()
}
