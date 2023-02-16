package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPullsItemUpdateBranchResponse 
type ItemItemPullsItemUpdateBranchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The message property
    message *string
    // The url property
    url *string
}
// NewItemItemPullsItemUpdateBranchResponse instantiates a new ItemItemPullsItemUpdateBranchResponse and sets the default values.
func NewItemItemPullsItemUpdateBranchResponse()(*ItemItemPullsItemUpdateBranchResponse) {
    m := &ItemItemPullsItemUpdateBranchResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPullsItemUpdateBranchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPullsItemUpdateBranchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPullsItemUpdateBranchResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPullsItemUpdateBranchResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemPullsItemUpdateBranchResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetMessage gets the message property value. The message property
func (m *ItemItemPullsItemUpdateBranchResponse) GetMessage()(*string) {
    return m.message
}
// GetUrl gets the url property value. The url property
func (m *ItemItemPullsItemUpdateBranchResponse) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ItemItemPullsItemUpdateBranchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *ItemItemPullsItemUpdateBranchResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The message property
func (m *ItemItemPullsItemUpdateBranchResponse) SetMessage(value *string)() {
    m.message = value
}
// SetUrl sets the url property value. The url property
func (m *ItemItemPullsItemUpdateBranchResponse) SetUrl(value *string)() {
    m.url = value
}
