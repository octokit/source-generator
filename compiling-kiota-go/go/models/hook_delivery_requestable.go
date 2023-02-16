package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HookDelivery_requestable 
type HookDelivery_requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(HookDelivery_request_headersable)
    GetPayload()(HookDelivery_request_payloadable)
    SetHeaders(value HookDelivery_request_headersable)()
    SetPayload(value HookDelivery_request_payloadable)()
}
