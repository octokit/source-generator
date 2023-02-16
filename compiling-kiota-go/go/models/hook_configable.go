package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hook_configable 
type Hook_configable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetDigest()(*string)
    GetEmail()(*string)
    GetInsecureSsl()(WebhookConfigInsecureSslable)
    GetPassword()(*string)
    GetRoom()(*string)
    GetSecret()(*string)
    GetSubdomain()(*string)
    GetToken()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetDigest(value *string)()
    SetEmail(value *string)()
    SetInsecureSsl(value WebhookConfigInsecureSslable)()
    SetPassword(value *string)()
    SetRoom(value *string)()
    SetSecret(value *string)()
    SetSubdomain(value *string)()
    SetToken(value *string)()
    SetUrl(value *string)()
}
