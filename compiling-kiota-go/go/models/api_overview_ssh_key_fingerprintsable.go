package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiOverview_ssh_key_fingerprintsable 
type ApiOverview_ssh_key_fingerprintsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSHA256DSA()(*string)
    GetSHA256ECDSA()(*string)
    GetSHA256ED25519()(*string)
    GetSHA256RSA()(*string)
    SetSHA256DSA(value *string)()
    SetSHA256ECDSA(value *string)()
    SetSHA256ED25519(value *string)()
    SetSHA256RSA(value *string)()
}
