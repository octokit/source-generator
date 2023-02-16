package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GpgKeyable 
type GpgKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanCertify()(*bool)
    GetCanEncryptComms()(*bool)
    GetCanEncryptStorage()(*bool)
    GetCanSign()(*bool)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmails()([]GpgKey_emailsable)
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int32)
    GetKeyId()(*string)
    GetName()(*string)
    GetPrimaryKeyId()(*int32)
    GetPublicKey()(*string)
    GetRawKey()(*string)
    GetRevoked()(*bool)
    GetSubkeys()([]GpgKey_subkeysable)
    SetCanCertify(value *bool)()
    SetCanEncryptComms(value *bool)()
    SetCanEncryptStorage(value *bool)()
    SetCanSign(value *bool)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmails(value []GpgKey_emailsable)()
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int32)()
    SetKeyId(value *string)()
    SetName(value *string)()
    SetPrimaryKeyId(value *int32)()
    SetPublicKey(value *string)()
    SetRawKey(value *string)()
    SetRevoked(value *bool)()
    SetSubkeys(value []GpgKey_subkeysable)()
}
