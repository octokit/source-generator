package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentsMember2 
type EnvironmentsMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
    reviewers []EnvironmentsMember2_reviewersable
    // The type property
    typeEscaped *string
}
// NewEnvironmentsMember2 instantiates a new environmentsMember2 and sets the default values.
func NewEnvironmentsMember2()(*EnvironmentsMember2) {
    m := &EnvironmentsMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentsMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironmentsMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironmentsMember2(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnvironmentsMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnvironmentsMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnvironmentsMember2_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnvironmentsMember2_reviewersable, len(val))
            for i, v := range val {
                res[i] = v.(EnvironmentsMember2_reviewersable)
            }
            m.SetReviewers(res)
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
    return res
}
// GetId gets the id property value. The id property
func (m *EnvironmentsMember2) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
func (m *EnvironmentsMember2) GetNodeId()(*string) {
    return m.node_id
}
// GetReviewers gets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *EnvironmentsMember2) GetReviewers()([]EnvironmentsMember2_reviewersable) {
    return m.reviewers
}
// GetType gets the type property value. The type property
func (m *EnvironmentsMember2) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *EnvironmentsMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnvironmentsMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *EnvironmentsMember2) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *EnvironmentsMember2) SetNodeId(value *string)() {
    m.node_id = value
}
// SetReviewers sets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *EnvironmentsMember2) SetReviewers(value []EnvironmentsMember2_reviewersable)() {
    m.reviewers = value
}
// SetType sets the type property value. The type property
func (m *EnvironmentsMember2) SetType(value *string)() {
    m.typeEscaped = value
}
