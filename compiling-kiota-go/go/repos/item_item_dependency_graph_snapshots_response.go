package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDependencyGraphSnapshotsResponse 
type ItemItemDependencyGraphSnapshotsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time at which the snapshot was created.
    created_at *string
    // ID of the created snapshot.
    id *int32
    // A message providing further details about the result, such as why the dependencies were not updated.
    message *string
    // Either "SUCCESS", "ACCEPTED", or "INVALID". "SUCCESS" indicates that the snapshot was successfully created and the repository's dependencies were updated. "ACCEPTED" indicates that the snapshot was successfully created, but the repository's dependencies were not updated. "INVALID" indicates that the snapshot was malformed.
    result *string
}
// NewItemItemDependencyGraphSnapshotsResponse instantiates a new ItemItemDependencyGraphSnapshotsResponse and sets the default values.
func NewItemItemDependencyGraphSnapshotsResponse()(*ItemItemDependencyGraphSnapshotsResponse) {
    m := &ItemItemDependencyGraphSnapshotsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDependencyGraphSnapshotsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDependencyGraphSnapshotsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDependencyGraphSnapshotsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDependencyGraphSnapshotsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time at which the snapshot was created.
func (m *ItemItemDependencyGraphSnapshotsResponse) GetCreatedAt()(*string) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemDependencyGraphSnapshotsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. ID of the created snapshot.
func (m *ItemItemDependencyGraphSnapshotsResponse) GetId()(*int32) {
    return m.id
}
// GetMessage gets the message property value. A message providing further details about the result, such as why the dependencies were not updated.
func (m *ItemItemDependencyGraphSnapshotsResponse) GetMessage()(*string) {
    return m.message
}
// GetResult gets the result property value. Either "SUCCESS", "ACCEPTED", or "INVALID". "SUCCESS" indicates that the snapshot was successfully created and the repository's dependencies were updated. "ACCEPTED" indicates that the snapshot was successfully created, but the repository's dependencies were not updated. "INVALID" indicates that the snapshot was malformed.
func (m *ItemItemDependencyGraphSnapshotsResponse) GetResult()(*string) {
    return m.result
}
// Serialize serializes information the current object
func (m *ItemItemDependencyGraphSnapshotsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
        err := writer.WriteStringValue("result", m.GetResult())
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
func (m *ItemItemDependencyGraphSnapshotsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time at which the snapshot was created.
func (m *ItemItemDependencyGraphSnapshotsResponse) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetId sets the id property value. ID of the created snapshot.
func (m *ItemItemDependencyGraphSnapshotsResponse) SetId(value *int32)() {
    m.id = value
}
// SetMessage sets the message property value. A message providing further details about the result, such as why the dependencies were not updated.
func (m *ItemItemDependencyGraphSnapshotsResponse) SetMessage(value *string)() {
    m.message = value
}
// SetResult sets the result property value. Either "SUCCESS", "ACCEPTED", or "INVALID". "SUCCESS" indicates that the snapshot was successfully created and the repository's dependencies were updated. "ACCEPTED" indicates that the snapshot was successfully created, but the repository's dependencies were not updated. "INVALID" indicates that the snapshot was malformed.
func (m *ItemItemDependencyGraphSnapshotsResponse) SetResult(value *string)() {
    m.result = value
}
