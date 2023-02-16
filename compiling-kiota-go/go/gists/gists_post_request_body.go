package gists

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistsPostRequestBody 
type GistsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description of the gist
    description *string
    // Names and content for the files that make up the gist
    files GistsPostRequestBody_filesable
    // The public property
    public GistsPostRequestBody_Gistsable
}
// GistsPostRequestBody_Gists composed type wrapper for classes boolean, string
type GistsPostRequestBody_Gists struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type boolean
    boolean *bool
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewGistsPostRequestBody_Gists instantiates a new gists and sets the default values.
func NewGistsPostRequestBody_Gists()(*GistsPostRequestBody_Gists) {
    m := &GistsPostRequestBody_Gists{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistsPostRequestBody_GistsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGistsPostRequestBody_GistsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewGistsPostRequestBody_Gists()
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
    if val, err := parseNode.GetBoolValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetBoolean(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GistsPostRequestBody_Gists) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBoolean gets the boolean property value. Composed type representation for type boolean
func (m *GistsPostRequestBody_Gists) GetBoolean()(*bool) {
    return m.boolean
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GistsPostRequestBody_Gists) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *GistsPostRequestBody_Gists) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *GistsPostRequestBody_Gists) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBoolean() != nil {
        err := writer.WriteBoolValue("", m.GetBoolean())
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
func (m *GistsPostRequestBody_Gists) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBoolean sets the boolean property value. Composed type representation for type boolean
func (m *GistsPostRequestBody_Gists) SetBoolean(value *bool)() {
    m.boolean = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *GistsPostRequestBody_Gists) SetString(value *string)() {
    m.string = value
}
// GistsPostRequestBody_Gistsable 
type GistsPostRequestBody_Gistsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBoolean()(*bool)
    GetString()(*string)
    SetBoolean(value *bool)()
    SetString(value *string)()
}
// NewGistsPostRequestBody instantiates a new gistsPostRequestBody and sets the default values.
func NewGistsPostRequestBody()(*GistsPostRequestBody) {
    m := &GistsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGistsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGistsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGistsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GistsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description of the gist
func (m *GistsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GistsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistsPostRequestBody_filesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFiles(val.(GistsPostRequestBody_filesable))
        }
        return nil
    }
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGistsPostRequestBody_GistsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val.(GistsPostRequestBody_Gistsable))
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. Names and content for the files that make up the gist
func (m *GistsPostRequestBody) GetFiles()(GistsPostRequestBody_filesable) {
    return m.files
}
// GetPublic gets the public property value. The public property
func (m *GistsPostRequestBody) GetPublic()(GistsPostRequestBody_Gistsable) {
    return m.public
}
// Serialize serializes information the current object
func (m *GistsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("files", m.GetFiles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("public", m.GetPublic())
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
func (m *GistsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description of the gist
func (m *GistsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetFiles sets the files property value. Names and content for the files that make up the gist
func (m *GistsPostRequestBody) SetFiles(value GistsPostRequestBody_filesable)() {
    m.files = value
}
// SetPublic sets the public property value. The public property
func (m *GistsPostRequestBody) SetPublic(value GistsPostRequestBody_Gistsable)() {
    m.public = value
}
