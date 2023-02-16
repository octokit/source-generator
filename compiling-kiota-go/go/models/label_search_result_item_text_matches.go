package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelSearchResultItem_text_matches 
type LabelSearchResultItem_text_matches struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fragment property
    fragment *string
    // The matches property
    matches []LabelSearchResultItem_text_matches_matchesable
    // The object_type property
    object_type *string
    // The object_url property
    object_url *string
    // The property property
    property *string
}
// NewLabelSearchResultItem_text_matches instantiates a new labelSearchResultItem_text_matches and sets the default values.
func NewLabelSearchResultItem_text_matches()(*LabelSearchResultItem_text_matches) {
    m := &LabelSearchResultItem_text_matches{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelSearchResultItem_text_matchesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelSearchResultItem_text_matchesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelSearchResultItem_text_matches(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelSearchResultItem_text_matches) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelSearchResultItem_text_matches) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fragment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFragment(val)
        }
        return nil
    }
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelSearchResultItem_text_matches_matchesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelSearchResultItem_text_matches_matchesable, len(val))
            for i, v := range val {
                res[i] = v.(LabelSearchResultItem_text_matches_matchesable)
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["object_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectType(val)
        }
        return nil
    }
    res["object_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectUrl(val)
        }
        return nil
    }
    res["property"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperty(val)
        }
        return nil
    }
    return res
}
// GetFragment gets the fragment property value. The fragment property
func (m *LabelSearchResultItem_text_matches) GetFragment()(*string) {
    return m.fragment
}
// GetMatches gets the matches property value. The matches property
func (m *LabelSearchResultItem_text_matches) GetMatches()([]LabelSearchResultItem_text_matches_matchesable) {
    return m.matches
}
// GetObjectType gets the object_type property value. The object_type property
func (m *LabelSearchResultItem_text_matches) GetObjectType()(*string) {
    return m.object_type
}
// GetObjectUrl gets the object_url property value. The object_url property
func (m *LabelSearchResultItem_text_matches) GetObjectUrl()(*string) {
    return m.object_url
}
// GetProperty gets the property property value. The property property
func (m *LabelSearchResultItem_text_matches) GetProperty()(*string) {
    return m.property
}
// Serialize serializes information the current object
func (m *LabelSearchResultItem_text_matches) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fragment", m.GetFragment())
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("matches", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("object_type", m.GetObjectType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("object_url", m.GetObjectUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("property", m.GetProperty())
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
func (m *LabelSearchResultItem_text_matches) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFragment sets the fragment property value. The fragment property
func (m *LabelSearchResultItem_text_matches) SetFragment(value *string)() {
    m.fragment = value
}
// SetMatches sets the matches property value. The matches property
func (m *LabelSearchResultItem_text_matches) SetMatches(value []LabelSearchResultItem_text_matches_matchesable)() {
    m.matches = value
}
// SetObjectType sets the object_type property value. The object_type property
func (m *LabelSearchResultItem_text_matches) SetObjectType(value *string)() {
    m.object_type = value
}
// SetObjectUrl sets the object_url property value. The object_url property
func (m *LabelSearchResultItem_text_matches) SetObjectUrl(value *string)() {
    m.object_url = value
}
// SetProperty sets the property property value. The property property
func (m *LabelSearchResultItem_text_matches) SetProperty(value *string)() {
    m.property = value
}
