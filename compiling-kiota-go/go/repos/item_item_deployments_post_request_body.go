package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemDeploymentsPostRequestBody 
type ItemItemDeploymentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Attempts to automatically merge the default branch into the requested ref, if it's behind the default branch.
    auto_merge *bool
    // Short description of the deployment.
    description *string
    // Name for the target deployment environment (e.g., `production`, `staging`, `qa`).
    environment *string
    // The payload property
    payload ItemItemDeploymentsPostRequestBody_Deploymentsable
    // Specifies if the given environment is one that end-users directly interact with. Default: `true` when `environment` is `production` and `false` otherwise.
    production_environment *bool
    // The ref to deploy. This can be a branch, tag, or SHA.
    ref *string
    // The [status](https://docs.github.com/rest/commits/statuses) contexts to verify against commit status checks. If you omit this parameter, GitHub verifies all unique contexts before creating a deployment. To bypass checking entirely, pass an empty array. Defaults to all unique contexts.
    required_contexts []string
    // Specifies a task to execute (e.g., `deploy` or `deploy:migrations`).
    task *string
    // Specifies if the given environment is specific to the deployment and will no longer exist at some point in the future. Default: `false`
    transient_environment *bool
}
// ItemItemDeploymentsPostRequestBody_Deployments composed type wrapper for classes deploymentsMember1, string
type ItemItemDeploymentsPostRequestBody_Deployments struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type deploymentsMember1
    deploymentsMember1 ItemItemDeploymentsMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewItemItemDeploymentsPostRequestBody_Deployments instantiates a new deployments and sets the default values.
func NewItemItemDeploymentsPostRequestBody_Deployments()(*ItemItemDeploymentsPostRequestBody_Deployments) {
    m := &ItemItemDeploymentsPostRequestBody_Deployments{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDeploymentsPostRequestBody_DeploymentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDeploymentsPostRequestBody_DeploymentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemDeploymentsPostRequestBody_Deployments()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDeploymentsPostRequestBody_Deployments) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeploymentsMember1 gets the deploymentsMember1 property value. Composed type representation for type deploymentsMember1
func (m *ItemItemDeploymentsPostRequestBody_Deployments) GetDeploymentsMember1()(ItemItemDeploymentsMember1able) {
    return m.deploymentsMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemDeploymentsPostRequestBody_Deployments) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *ItemItemDeploymentsPostRequestBody_Deployments) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemDeploymentsPostRequestBody_Deployments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeploymentsMember1() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentsMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *ItemItemDeploymentsPostRequestBody_Deployments) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeploymentsMember1 sets the deploymentsMember1 property value. Composed type representation for type deploymentsMember1
func (m *ItemItemDeploymentsPostRequestBody_Deployments) SetDeploymentsMember1(value ItemItemDeploymentsMember1able)() {
    m.deploymentsMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemDeploymentsPostRequestBody_Deployments) SetString(value *string)() {
    m.string = value
}
// ItemItemDeploymentsPostRequestBody_Deploymentsable 
type ItemItemDeploymentsPostRequestBody_Deploymentsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentsMember1()(ItemItemDeploymentsMember1able)
    GetString()(*string)
    SetDeploymentsMember1(value ItemItemDeploymentsMember1able)()
    SetString(value *string)()
}
// NewItemItemDeploymentsPostRequestBody instantiates a new ItemItemDeploymentsPostRequestBody and sets the default values.
func NewItemItemDeploymentsPostRequestBody()(*ItemItemDeploymentsPostRequestBody) {
    m := &ItemItemDeploymentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    environmentValue := "production"
    m.SetEnvironment(&environmentValue)
    taskValue := "deploy"
    m.SetTask(&taskValue)
    return m
}
// CreateItemItemDeploymentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDeploymentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDeploymentsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDeploymentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutoMerge gets the auto_merge property value. Attempts to automatically merge the default branch into the requested ref, if it's behind the default branch.
func (m *ItemItemDeploymentsPostRequestBody) GetAutoMerge()(*bool) {
    return m.auto_merge
}
// GetDescription gets the description property value. Short description of the deployment.
func (m *ItemItemDeploymentsPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetEnvironment gets the environment property value. Name for the target deployment environment (e.g., `production`, `staging`, `qa`).
func (m *ItemItemDeploymentsPostRequestBody) GetEnvironment()(*string) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemDeploymentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auto_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoMerge(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemDeploymentsPostRequestBody_DeploymentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(ItemItemDeploymentsPostRequestBody_Deploymentsable))
        }
        return nil
    }
    res["production_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductionEnvironment(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["required_contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRequiredContexts(res)
        }
        return nil
    }
    res["task"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTask(val)
        }
        return nil
    }
    res["transient_environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransientEnvironment(val)
        }
        return nil
    }
    return res
}
// GetPayload gets the payload property value. The payload property
func (m *ItemItemDeploymentsPostRequestBody) GetPayload()(ItemItemDeploymentsPostRequestBody_Deploymentsable) {
    return m.payload
}
// GetProductionEnvironment gets the production_environment property value. Specifies if the given environment is one that end-users directly interact with. Default: `true` when `environment` is `production` and `false` otherwise.
func (m *ItemItemDeploymentsPostRequestBody) GetProductionEnvironment()(*bool) {
    return m.production_environment
}
// GetRef gets the ref property value. The ref to deploy. This can be a branch, tag, or SHA.
func (m *ItemItemDeploymentsPostRequestBody) GetRef()(*string) {
    return m.ref
}
// GetRequiredContexts gets the required_contexts property value. The [status](https://docs.github.com/rest/commits/statuses) contexts to verify against commit status checks. If you omit this parameter, GitHub verifies all unique contexts before creating a deployment. To bypass checking entirely, pass an empty array. Defaults to all unique contexts.
func (m *ItemItemDeploymentsPostRequestBody) GetRequiredContexts()([]string) {
    return m.required_contexts
}
// GetTask gets the task property value. Specifies a task to execute (e.g., `deploy` or `deploy:migrations`).
func (m *ItemItemDeploymentsPostRequestBody) GetTask()(*string) {
    return m.task
}
// GetTransientEnvironment gets the transient_environment property value. Specifies if the given environment is specific to the deployment and will no longer exist at some point in the future. Default: `false`
func (m *ItemItemDeploymentsPostRequestBody) GetTransientEnvironment()(*bool) {
    return m.transient_environment
}
// Serialize serializes information the current object
func (m *ItemItemDeploymentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("auto_merge", m.GetAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("production_environment", m.GetProductionEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredContexts() != nil {
        err := writer.WriteCollectionOfStringValues("required_contexts", m.GetRequiredContexts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("transient_environment", m.GetTransientEnvironment())
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
func (m *ItemItemDeploymentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutoMerge sets the auto_merge property value. Attempts to automatically merge the default branch into the requested ref, if it's behind the default branch.
func (m *ItemItemDeploymentsPostRequestBody) SetAutoMerge(value *bool)() {
    m.auto_merge = value
}
// SetDescription sets the description property value. Short description of the deployment.
func (m *ItemItemDeploymentsPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetEnvironment sets the environment property value. Name for the target deployment environment (e.g., `production`, `staging`, `qa`).
func (m *ItemItemDeploymentsPostRequestBody) SetEnvironment(value *string)() {
    m.environment = value
}
// SetPayload sets the payload property value. The payload property
func (m *ItemItemDeploymentsPostRequestBody) SetPayload(value ItemItemDeploymentsPostRequestBody_Deploymentsable)() {
    m.payload = value
}
// SetProductionEnvironment sets the production_environment property value. Specifies if the given environment is one that end-users directly interact with. Default: `true` when `environment` is `production` and `false` otherwise.
func (m *ItemItemDeploymentsPostRequestBody) SetProductionEnvironment(value *bool)() {
    m.production_environment = value
}
// SetRef sets the ref property value. The ref to deploy. This can be a branch, tag, or SHA.
func (m *ItemItemDeploymentsPostRequestBody) SetRef(value *string)() {
    m.ref = value
}
// SetRequiredContexts sets the required_contexts property value. The [status](https://docs.github.com/rest/commits/statuses) contexts to verify against commit status checks. If you omit this parameter, GitHub verifies all unique contexts before creating a deployment. To bypass checking entirely, pass an empty array. Defaults to all unique contexts.
func (m *ItemItemDeploymentsPostRequestBody) SetRequiredContexts(value []string)() {
    m.required_contexts = value
}
// SetTask sets the task property value. Specifies a task to execute (e.g., `deploy` or `deploy:migrations`).
func (m *ItemItemDeploymentsPostRequestBody) SetTask(value *string)() {
    m.task = value
}
// SetTransientEnvironment sets the transient_environment property value. Specifies if the given environment is specific to the deployment and will no longer exist at some point in the future. Default: `false`
func (m *ItemItemDeploymentsPostRequestBody) SetTransientEnvironment(value *bool)() {
    m.transient_environment = value
}
