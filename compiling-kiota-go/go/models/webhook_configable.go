package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookConfigable 
type WebhookConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(*string)
    GetInsecureSsl()(WebhookConfigInsecureSslable)
    GetSecret()(*string)
    GetUrl()(*string)
    SetContentType(value *string)()
    SetInsecureSsl(value WebhookConfigInsecureSslable)()
    SetSecret(value *string)()
    SetUrl(value *string)()
}
