package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCodespacesNewResponse 
type ItemItemCodespacesNewResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    billable_owner ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable
    // The defaults property
    defaults ItemItemCodespacesNewResponse_defaultsable
}
// NewItemItemCodespacesNewResponse instantiates a new ItemItemCodespacesNewResponse and sets the default values.
func NewItemItemCodespacesNewResponse()(*ItemItemCodespacesNewResponse) {
    m := &ItemItemCodespacesNewResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesNewResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesNewResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesNewResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCodespacesNewResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillableOwner gets the billable_owner property value. A GitHub user.
func (m *ItemItemCodespacesNewResponse) GetBillableOwner()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable) {
    return m.billable_owner
}
// GetDefaults gets the defaults property value. The defaults property
func (m *ItemItemCodespacesNewResponse) GetDefaults()(ItemItemCodespacesNewResponse_defaultsable) {
    return m.defaults
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCodespacesNewResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billable_owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillableOwner(val.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable))
        }
        return nil
    }
    res["defaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemCodespacesNewResponse_defaultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaults(val.(ItemItemCodespacesNewResponse_defaultsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesNewResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("billable_owner", m.GetBillableOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaults", m.GetDefaults())
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
func (m *ItemItemCodespacesNewResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillableOwner sets the billable_owner property value. A GitHub user.
func (m *ItemItemCodespacesNewResponse) SetBillableOwner(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.SimpleUserable)() {
    m.billable_owner = value
}
// SetDefaults sets the defaults property value. The defaults property
func (m *ItemItemCodespacesNewResponse) SetDefaults(value ItemItemCodespacesNewResponse_defaultsable)() {
    m.defaults = value
}
