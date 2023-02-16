package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidationError_errors 
type ValidationError_errors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The field property
    field *string
    // The index property
    index *int32
    // The message property
    message *string
    // The resource property
    resource *string
    // The value property
    value ValidationError_errors_Gistsable
}
// ValidationError_errors_Gists composed type wrapper for classes string, integer, string
type ValidationError_errors_Gists struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type integer
    integer *int32
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewValidationError_errors_Gists instantiates a new gists and sets the default values.
func NewValidationError_errors_Gists()(*ValidationError_errors_Gists) {
    m := &ValidationError_errors_Gists{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidationError_errors_GistsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidationError_errors_GistsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewValidationError_errors_Gists()
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
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationError_errors_Gists) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidationError_errors_Gists) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *ValidationError_errors_Gists) GetInteger()(*int32) {
    return m.integer
}
// GetString gets the string property value. Composed type representation for type string
func (m *ValidationError_errors_Gists) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ValidationError_errors_Gists) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
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
func (m *ValidationError_errors_Gists) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInteger sets the integer property value. Composed type representation for type integer
func (m *ValidationError_errors_Gists) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ValidationError_errors_Gists) SetString(value *string)() {
    m.string = value
}
// ValidationError_errors_Gistsable 
type ValidationError_errors_Gistsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewValidationError_errors instantiates a new ValidationError_errors and sets the default values.
func NewValidationError_errors()(*ValidationError_errors) {
    m := &ValidationError_errors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidationError_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidationError_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationError_errors(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationError_errors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *ValidationError_errors) GetCode()(*string) {
    return m.code
}
// GetField gets the field property value. The field property
func (m *ValidationError_errors) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidationError_errors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidationError_errors_GistsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(ValidationError_errors_Gistsable))
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index property
func (m *ValidationError_errors) GetIndex()(*int32) {
    return m.index
}
// GetMessage gets the message property value. The message property
func (m *ValidationError_errors) GetMessage()(*string) {
    return m.message
}
// GetResource gets the resource property value. The resource property
func (m *ValidationError_errors) GetResource()(*string) {
    return m.resource
}
// GetValue gets the value property value. The value property
func (m *ValidationError_errors) GetValue()(ValidationError_errors_Gistsable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ValidationError_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
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
func (m *ValidationError_errors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *ValidationError_errors) SetCode(value *string)() {
    m.code = value
}
// SetField sets the field property value. The field property
func (m *ValidationError_errors) SetField(value *string)() {
    m.field = value
}
// SetIndex sets the index property value. The index property
func (m *ValidationError_errors) SetIndex(value *int32)() {
    m.index = value
}
// SetMessage sets the message property value. The message property
func (m *ValidationError_errors) SetMessage(value *string)() {
    m.message = value
}
// SetResource sets the resource property value. The resource property
func (m *ValidationError_errors) SetResource(value *string)() {
    m.resource = value
}
// SetValue sets the value property value. The value property
func (m *ValidationError_errors) SetValue(value ValidationError_errors_Gistsable)() {
    m.value = value
}
