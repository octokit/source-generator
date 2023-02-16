package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Environment details of a deployment environment
type Environment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the environment was created, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
    deployment_branch_policy DeploymentBranchPolicySettingsable
    // The html_url property
    html_url *string
    // The id of the environment.
    id *int32
    // The name of the environment.
    name *string
    // The node_id property
    node_id *string
    // The protection_rules property
    protection_rules []Environment_Environmentsable
    // The time that the environment was last updated, in ISO 8601 format.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// Environment_Environments composed type wrapper for classes environmentsMember1, environmentsMember2, environmentsMember3
type Environment_Environments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type environmentsMember1
    environmentsMember1 EnvironmentsMember1able
    // Composed type representation for type environmentsMember2
    environmentsMember2 EnvironmentsMember2able
    // Composed type representation for type environmentsMember3
    environmentsMember3 EnvironmentsMember3able
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewEnvironment_Environments instantiates a new environments and sets the default values.
func NewEnvironment_Environments()(*Environment_Environments) {
    m := &Environment_Environments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironment_EnvironmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironment_EnvironmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEnvironment_Environments()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateEnvironmentsMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(EnvironmentsMember1able); ok {
                result.SetEnvironmentsMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateEnvironmentsMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(EnvironmentsMember2able); ok {
                result.SetEnvironmentsMember2(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateEnvironmentsMember3FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(EnvironmentsMember3able); ok {
                result.SetEnvironmentsMember3(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Environment_Environments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnvironmentsMember1 gets the environmentsMember1 property value. Composed type representation for type environmentsMember1
func (m *Environment_Environments) GetEnvironmentsMember1()(EnvironmentsMember1able) {
    return m.environmentsMember1
}
// GetEnvironmentsMember2 gets the environmentsMember2 property value. Composed type representation for type environmentsMember2
func (m *Environment_Environments) GetEnvironmentsMember2()(EnvironmentsMember2able) {
    return m.environmentsMember2
}
// GetEnvironmentsMember3 gets the environmentsMember3 property value. Composed type representation for type environmentsMember3
func (m *Environment_Environments) GetEnvironmentsMember3()(EnvironmentsMember3able) {
    return m.environmentsMember3
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Environment_Environments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// Serialize serializes information the current object
func (m *Environment_Environments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnvironmentsMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentsMember1())
        if err != nil {
            return err
        }
    } else if m.GetEnvironmentsMember2() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentsMember2())
        if err != nil {
            return err
        }
    } else if m.GetEnvironmentsMember3() != nil {
        err := writer.WriteObjectValue("", m.GetEnvironmentsMember3())
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
func (m *Environment_Environments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnvironmentsMember1 sets the environmentsMember1 property value. Composed type representation for type environmentsMember1
func (m *Environment_Environments) SetEnvironmentsMember1(value EnvironmentsMember1able)() {
    m.environmentsMember1 = value
}
// SetEnvironmentsMember2 sets the environmentsMember2 property value. Composed type representation for type environmentsMember2
func (m *Environment_Environments) SetEnvironmentsMember2(value EnvironmentsMember2able)() {
    m.environmentsMember2 = value
}
// SetEnvironmentsMember3 sets the environmentsMember3 property value. Composed type representation for type environmentsMember3
func (m *Environment_Environments) SetEnvironmentsMember3(value EnvironmentsMember3able)() {
    m.environmentsMember3 = value
}
// Environment_Environmentsable 
type Environment_Environmentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnvironmentsMember1()(EnvironmentsMember1able)
    GetEnvironmentsMember2()(EnvironmentsMember2able)
    GetEnvironmentsMember3()(EnvironmentsMember3able)
    SetEnvironmentsMember1(value EnvironmentsMember1able)()
    SetEnvironmentsMember2(value EnvironmentsMember2able)()
    SetEnvironmentsMember3(value EnvironmentsMember3able)()
}
// NewEnvironment instantiates a new environment and sets the default values.
func NewEnvironment()(*Environment) {
    m := &Environment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnvironmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Environment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the environment was created, in ISO 8601 format.
func (m *Environment) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDeploymentBranchPolicy gets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *Environment) GetDeploymentBranchPolicy()(DeploymentBranchPolicySettingsable) {
    return m.deployment_branch_policy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Environment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["deployment_branch_policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentBranchPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentBranchPolicy(val.(DeploymentBranchPolicySettingsable))
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
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
    res["protection_rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnvironment_EnvironmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Environment_Environmentsable, len(val))
            for i, v := range val {
                res[i] = v.(Environment_Environmentsable)
            }
            m.SetProtectionRules(res)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
// GetHtmlUrl gets the html_url property value. The html_url property
func (m *Environment) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id of the environment.
func (m *Environment) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the environment.
func (m *Environment) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
func (m *Environment) GetNodeId()(*string) {
    return m.node_id
}
// GetProtectionRules gets the protection_rules property value. The protection_rules property
func (m *Environment) GetProtectionRules()([]Environment_Environmentsable) {
    return m.protection_rules
}
// GetUpdatedAt gets the updated_at property value. The time that the environment was last updated, in ISO 8601 format.
func (m *Environment) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
func (m *Environment) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Environment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deployment_branch_policy", m.GetDeploymentBranchPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
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
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetProtectionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProtectionRules()))
        for i, v := range m.GetProtectionRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("protection_rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *Environment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the environment was created, in ISO 8601 format.
func (m *Environment) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDeploymentBranchPolicy sets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *Environment) SetDeploymentBranchPolicy(value DeploymentBranchPolicySettingsable)() {
    m.deployment_branch_policy = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Environment) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id of the environment.
func (m *Environment) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the environment.
func (m *Environment) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Environment) SetNodeId(value *string)() {
    m.node_id = value
}
// SetProtectionRules sets the protection_rules property value. The protection_rules property
func (m *Environment) SetProtectionRules(value []Environment_Environmentsable)() {
    m.protection_rules = value
}
// SetUpdatedAt sets the updated_at property value. The time that the environment was last updated, in ISO 8601 format.
func (m *Environment) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *Environment) SetUrl(value *string)() {
    m.url = value
}
