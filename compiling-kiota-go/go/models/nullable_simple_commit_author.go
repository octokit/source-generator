package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableSimpleCommit_author 
type NullableSimpleCommit_author struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The name property
    name *string
}
// NewNullableSimpleCommit_author instantiates a new nullableSimpleCommit_author and sets the default values.
func NewNullableSimpleCommit_author()(*NullableSimpleCommit_author) {
    m := &NullableSimpleCommit_author{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableSimpleCommit_authorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNullableSimpleCommit_authorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableSimpleCommit_author(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NullableSimpleCommit_author) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
func (m *NullableSimpleCommit_author) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NullableSimpleCommit_author) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *NullableSimpleCommit_author) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *NullableSimpleCommit_author) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *NullableSimpleCommit_author) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *NullableSimpleCommit_author) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name property
func (m *NullableSimpleCommit_author) SetName(value *string)() {
    m.name = value
}
