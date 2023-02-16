package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentsMember2_reviewers 
type EnvironmentsMember2_reviewers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reviewer property
    reviewer EnvironmentsMember2_reviewers_Environmentsable
    // The type of reviewer.
    typeEscaped *DeploymentReviewerType
}
// EnvironmentsMember2_reviewers_Environments composed type wrapper for classes simpleUser, team
type EnvironmentsMember2_reviewers_Environments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type simpleUser
    simpleUser SimpleUserable
    // Composed type representation for type team
    team Teamable
}
// NewEnvironmentsMember2_reviewers_Environments instantiates a new environments and sets the default values.
func NewEnvironmentsMember2_reviewers_Environments()(*EnvironmentsMember2_reviewers_Environments) {
    m := &EnvironmentsMember2_reviewers_Environments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentsMember2_reviewers_EnvironmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironmentsMember2_reviewers_EnvironmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEnvironmentsMember2_reviewers_Environments()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(SimpleUserable); ok {
                result.SetSimpleUser(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTeamFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Teamable); ok {
                result.SetTeam(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnvironmentsMember2_reviewers_Environments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnvironmentsMember2_reviewers_Environments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type simpleUser
func (m *EnvironmentsMember2_reviewers_Environments) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// GetTeam gets the team property value. Composed type representation for type team
func (m *EnvironmentsMember2_reviewers_Environments) GetTeam()(Teamable) {
    return m.team
}
// Serialize serializes information the current object
func (m *EnvironmentsMember2_reviewers_Environments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetTeam() != nil {
        err := writer.WriteObjectValue("", m.GetTeam())
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
func (m *EnvironmentsMember2_reviewers_Environments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type simpleUser
func (m *EnvironmentsMember2_reviewers_Environments) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
// SetTeam sets the team property value. Composed type representation for type team
func (m *EnvironmentsMember2_reviewers_Environments) SetTeam(value Teamable)() {
    m.team = value
}
// EnvironmentsMember2_reviewers_Environmentsable 
type EnvironmentsMember2_reviewers_Environmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSimpleUser()(SimpleUserable)
    GetTeam()(Teamable)
    SetSimpleUser(value SimpleUserable)()
    SetTeam(value Teamable)()
}
// NewEnvironmentsMember2_reviewers instantiates a new environmentsMember2_reviewers and sets the default values.
func NewEnvironmentsMember2_reviewers()(*EnvironmentsMember2_reviewers) {
    m := &EnvironmentsMember2_reviewers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentsMember2_reviewersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironmentsMember2_reviewersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironmentsMember2_reviewers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnvironmentsMember2_reviewers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnvironmentsMember2_reviewers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnvironmentsMember2_reviewers_EnvironmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewer(val.(EnvironmentsMember2_reviewers_Environmentsable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentReviewerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*DeploymentReviewerType))
        }
        return nil
    }
    return res
}
// GetReviewer gets the reviewer property value. The reviewer property
func (m *EnvironmentsMember2_reviewers) GetReviewer()(EnvironmentsMember2_reviewers_Environmentsable) {
    return m.reviewer
}
// GetType gets the type property value. The type of reviewer.
func (m *EnvironmentsMember2_reviewers) GetType()(*DeploymentReviewerType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *EnvironmentsMember2_reviewers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("reviewer", m.GetReviewer())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *EnvironmentsMember2_reviewers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewer sets the reviewer property value. The reviewer property
func (m *EnvironmentsMember2_reviewers) SetReviewer(value EnvironmentsMember2_reviewers_Environmentsable)() {
    m.reviewer = value
}
// SetType sets the type property value. The type of reviewer.
func (m *EnvironmentsMember2_reviewers) SetType(value *DeploymentReviewerType)() {
    m.typeEscaped = value
}
