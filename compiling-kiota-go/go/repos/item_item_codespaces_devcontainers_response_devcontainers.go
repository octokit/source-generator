package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodespacesDevcontainersResponse_devcontainers 
type ItemItemCodespacesDevcontainersResponse_devcontainers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The path property
    path *string
}
// NewItemItemCodespacesDevcontainersResponse_devcontainers instantiates a new ItemItemCodespacesDevcontainersResponse_devcontainers and sets the default values.
func NewItemItemCodespacesDevcontainersResponse_devcontainers()(*ItemItemCodespacesDevcontainersResponse_devcontainers) {
    m := &ItemItemCodespacesDevcontainersResponse_devcontainers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesDevcontainersResponse_devcontainersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesDevcontainersResponse_devcontainersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesDevcontainersResponse_devcontainers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. The path property
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) GetPath()(*string) {
    return m.path
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
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
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. The path property
func (m *ItemItemCodespacesDevcontainersResponse_devcontainers) SetPath(value *string)() {
    m.path = value
}
