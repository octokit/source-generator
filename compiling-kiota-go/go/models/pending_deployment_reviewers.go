package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingDeployment_reviewers 
type PendingDeployment_reviewers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reviewer property
    reviewer PendingDeployment_reviewers_Pending_deploymentsable
    // The type of reviewer.
    typeEscaped *DeploymentReviewerType
}
// PendingDeployment_reviewers_Pending_deployments composed type wrapper for classes simpleUser, team
type PendingDeployment_reviewers_Pending_deployments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type simpleUser
    simpleUser SimpleUserable
    // Composed type representation for type team
    team Teamable
}
// NewPendingDeployment_reviewers_Pending_deployments instantiates a new pending_deployments and sets the default values.
func NewPendingDeployment_reviewers_Pending_deployments()(*PendingDeployment_reviewers_Pending_deployments) {
    m := &PendingDeployment_reviewers_Pending_deployments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePendingDeployment_reviewers_Pending_deploymentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingDeployment_reviewers_Pending_deploymentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewPendingDeployment_reviewers_Pending_deployments()
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
func (m *PendingDeployment_reviewers_Pending_deployments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingDeployment_reviewers_Pending_deployments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type simpleUser
func (m *PendingDeployment_reviewers_Pending_deployments) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// GetTeam gets the team property value. Composed type representation for type team
func (m *PendingDeployment_reviewers_Pending_deployments) GetTeam()(Teamable) {
    return m.team
}
// Serialize serializes information the current object
func (m *PendingDeployment_reviewers_Pending_deployments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PendingDeployment_reviewers_Pending_deployments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type simpleUser
func (m *PendingDeployment_reviewers_Pending_deployments) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
// SetTeam sets the team property value. Composed type representation for type team
func (m *PendingDeployment_reviewers_Pending_deployments) SetTeam(value Teamable)() {
    m.team = value
}
// PendingDeployment_reviewers_Pending_deploymentsable 
type PendingDeployment_reviewers_Pending_deploymentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSimpleUser()(SimpleUserable)
    GetTeam()(Teamable)
    SetSimpleUser(value SimpleUserable)()
    SetTeam(value Teamable)()
}
// NewPendingDeployment_reviewers instantiates a new PendingDeployment_reviewers and sets the default values.
func NewPendingDeployment_reviewers()(*PendingDeployment_reviewers) {
    m := &PendingDeployment_reviewers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePendingDeployment_reviewersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingDeployment_reviewersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingDeployment_reviewers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PendingDeployment_reviewers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingDeployment_reviewers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingDeployment_reviewers_Pending_deploymentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewer(val.(PendingDeployment_reviewers_Pending_deploymentsable))
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
func (m *PendingDeployment_reviewers) GetReviewer()(PendingDeployment_reviewers_Pending_deploymentsable) {
    return m.reviewer
}
// GetType gets the type property value. The type of reviewer.
func (m *PendingDeployment_reviewers) GetType()(*DeploymentReviewerType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *PendingDeployment_reviewers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PendingDeployment_reviewers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewer sets the reviewer property value. The reviewer property
func (m *PendingDeployment_reviewers) SetReviewer(value PendingDeployment_reviewers_Pending_deploymentsable)() {
    m.reviewer = value
}
// SetType sets the type property value. The type of reviewer.
func (m *PendingDeployment_reviewers) SetType(value *DeploymentReviewerType)() {
    m.typeEscaped = value
}
