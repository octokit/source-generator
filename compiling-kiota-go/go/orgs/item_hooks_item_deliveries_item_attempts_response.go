package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemHooksItemDeliveriesItemAttemptsResponse 
type ItemHooksItemDeliveriesItemAttemptsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemHooksItemDeliveriesItemAttemptsResponse instantiates a new ItemHooksItemDeliveriesItemAttemptsResponse and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsResponse()(*ItemHooksItemDeliveriesItemAttemptsResponse) {
    m := &ItemHooksItemDeliveriesItemAttemptsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemHooksItemDeliveriesItemAttemptsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemHooksItemDeliveriesItemAttemptsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemHooksItemDeliveriesItemAttemptsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemHooksItemDeliveriesItemAttemptsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemHooksItemDeliveriesItemAttemptsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
