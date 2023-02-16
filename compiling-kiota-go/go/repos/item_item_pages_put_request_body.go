package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesPutRequestBody 
type ItemItemPagesPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/articles/using-a-custom-domain-with-github-pages/)."
    cname *string
    // Specify whether HTTPS should be enforced for the repository.
    https_enforced *bool
    // The source property
    source ItemItemPagesPutRequestBody_Pagesable
}
// ItemItemPagesPutRequestBody_Pages composed type wrapper for classes string, pagesMember1
type ItemItemPagesPutRequestBody_Pages struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type pagesMember1
    pagesMember1 ItemItemPagesMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewItemItemPagesPutRequestBody_Pages instantiates a new pages and sets the default values.
func NewItemItemPagesPutRequestBody_Pages()(*ItemItemPagesPutRequestBody_Pages) {
    m := &ItemItemPagesPutRequestBody_Pages{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesPutRequestBody_PagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPagesPutRequestBody_PagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemPagesPutRequestBody_Pages()
    if parseNode != nil {
        if val, err := parseNode.GetStringValue(); val != nil {
            if err != nil {
                return nil, err
            }
            result.SetString(val)
        } else if val, err := parseNode.GetObjectValue(CreateItemItemPagesMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemPagesMember1able); ok {
                result.SetPagesMember1(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPagesPutRequestBody_Pages) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemPagesPutRequestBody_Pages) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetPagesMember1 gets the pagesMember1 property value. Composed type representation for type pagesMember1
func (m *ItemItemPagesPutRequestBody_Pages) GetPagesMember1()(ItemItemPagesMember1able) {
    return m.pagesMember1
}
// GetString gets the string property value. Composed type representation for type string
func (m *ItemItemPagesPutRequestBody_Pages) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemPagesPutRequestBody_Pages) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetPagesMember1() != nil {
        err := writer.WriteObjectValue("", m.GetPagesMember1())
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
func (m *ItemItemPagesPutRequestBody_Pages) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPagesMember1 sets the pagesMember1 property value. Composed type representation for type pagesMember1
func (m *ItemItemPagesPutRequestBody_Pages) SetPagesMember1(value ItemItemPagesMember1able)() {
    m.pagesMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemPagesPutRequestBody_Pages) SetString(value *string)() {
    m.string = value
}
// ItemItemPagesPutRequestBody_Pagesable 
type ItemItemPagesPutRequestBody_Pagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPagesMember1()(ItemItemPagesMember1able)
    GetString()(*string)
    SetPagesMember1(value ItemItemPagesMember1able)()
    SetString(value *string)()
}
// NewItemItemPagesPutRequestBody instantiates a new ItemItemPagesPutRequestBody and sets the default values.
func NewItemItemPagesPutRequestBody()(*ItemItemPagesPutRequestBody) {
    m := &ItemItemPagesPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPagesPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesPutRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPagesPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCname gets the cname property value. Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/articles/using-a-custom-domain-with-github-pages/)."
func (m *ItemItemPagesPutRequestBody) GetCname()(*string) {
    return m.cname
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemPagesPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCname(val)
        }
        return nil
    }
    res["https_enforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpsEnforced(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemPagesPutRequestBody_PagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(ItemItemPagesPutRequestBody_Pagesable))
        }
        return nil
    }
    return res
}
// GetHttpsEnforced gets the https_enforced property value. Specify whether HTTPS should be enforced for the repository.
func (m *ItemItemPagesPutRequestBody) GetHttpsEnforced()(*bool) {
    return m.https_enforced
}
// GetSource gets the source property value. The source property
func (m *ItemItemPagesPutRequestBody) GetSource()(ItemItemPagesPutRequestBody_Pagesable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ItemItemPagesPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cname", m.GetCname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("https_enforced", m.GetHttpsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
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
func (m *ItemItemPagesPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCname sets the cname property value. Specify a custom domain for the repository. Sending a `null` value will remove the custom domain. For more about custom domains, see "[Using a custom domain with GitHub Pages](https://docs.github.com/articles/using-a-custom-domain-with-github-pages/)."
func (m *ItemItemPagesPutRequestBody) SetCname(value *string)() {
    m.cname = value
}
// SetHttpsEnforced sets the https_enforced property value. Specify whether HTTPS should be enforced for the repository.
func (m *ItemItemPagesPutRequestBody) SetHttpsEnforced(value *bool)() {
    m.https_enforced = value
}
// SetSource sets the source property value. The source property
func (m *ItemItemPagesPutRequestBody) SetSource(value ItemItemPagesPutRequestBody_Pagesable)() {
    m.source = value
}
