package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitTag_object 
type GitTag_object struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sha property
    sha *string
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewGitTag_object instantiates a new GitTag_object and sets the default values.
func NewGitTag_object()(*GitTag_object) {
    m := &GitTag_object{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitTag_objectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGitTag_objectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitTag_object(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GitTag_object) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GitTag_object) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetSha gets the sha property value. The sha property
func (m *GitTag_object) GetSha()(*string) {
    return m.sha
}
// GetType gets the type property value. The type property
func (m *GitTag_object) GetType()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
func (m *GitTag_object) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *GitTag_object) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *GitTag_object) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSha sets the sha property value. The sha property
func (m *GitTag_object) SetSha(value *string)() {
    m.sha = value
}
// SetType sets the type property value. The type property
func (m *GitTag_object) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *GitTag_object) SetUrl(value *string)() {
    m.url = value
}
