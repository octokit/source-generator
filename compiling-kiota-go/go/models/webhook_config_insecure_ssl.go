package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookConfigInsecureSsl composed type wrapper for classes string, int64
type WebhookConfigInsecureSsl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type int64
    int64 *int64
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewWebhookConfigInsecureSsl instantiates a new WebhookConfigInsecureSsl and sets the default values.
func NewWebhookConfigInsecureSsl()(*WebhookConfigInsecureSsl) {
    m := &WebhookConfigInsecureSsl{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebhookConfigInsecureSslFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebhookConfigInsecureSslFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWebhookConfigInsecureSsl()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetInt64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInt64(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebhookConfigInsecureSsl) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebhookConfigInsecureSsl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInt64 gets the int64 property value. Composed type representation for type int64
func (m *WebhookConfigInsecureSsl) GetInt64()(*int64) {
    return m.int64
}
// GetString gets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *WebhookConfigInsecureSsl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInt64() != nil {
        err := writer.WriteInt64Value("", m.GetInt64())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebhookConfigInsecureSsl) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInt64 sets the int64 property value. Composed type representation for type int64
func (m *WebhookConfigInsecureSsl) SetInt64(value *int64)() {
    m.int64 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetString(value *string)() {
    m.string = value
}
