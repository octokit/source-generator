package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequiredWorkflow 
type RequiredWorkflow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier for a required workflow
    id *int64
    // Name present in the workflow file
    name *string
    // Path of the workflow file
    path *string
    // Ref at which the workflow file will be selected
    ref *string
    // Minimal Repository
    repository MinimalRepositoryable
    // Scope of the required workflow
    scope *RequiredWorkflow_scope
    // The selected_repositories_url property
    selected_repositories_url *string
    // State of the required workflow
    state *RequiredWorkflow_state
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewRequiredWorkflow instantiates a new requiredWorkflow and sets the default values.
func NewRequiredWorkflow()(*RequiredWorkflow) {
    m := &RequiredWorkflow{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequiredWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequiredWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequiredWorkflow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredWorkflow) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *RequiredWorkflow) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequiredWorkflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
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
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredWorkflow_scope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*RequiredWorkflow_scope))
        }
        return nil
    }
    res["selected_repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectedRepositoriesUrl(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredWorkflow_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*RequiredWorkflow_state))
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
    return res
}
// GetId gets the id property value. Unique identifier for a required workflow
func (m *RequiredWorkflow) GetId()(*int64) {
    return m.id
}
// GetName gets the name property value. Name present in the workflow file
func (m *RequiredWorkflow) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. Path of the workflow file
func (m *RequiredWorkflow) GetPath()(*string) {
    return m.path
}
// GetRef gets the ref property value. Ref at which the workflow file will be selected
func (m *RequiredWorkflow) GetRef()(*string) {
    return m.ref
}
// GetRepository gets the repository property value. Minimal Repository
func (m *RequiredWorkflow) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// GetScope gets the scope property value. Scope of the required workflow
func (m *RequiredWorkflow) GetScope()(*RequiredWorkflow_scope) {
    return m.scope
}
// GetSelectedRepositoriesUrl gets the selected_repositories_url property value. The selected_repositories_url property
func (m *RequiredWorkflow) GetSelectedRepositoriesUrl()(*string) {
    return m.selected_repositories_url
}
// GetState gets the state property value. State of the required workflow
func (m *RequiredWorkflow) GetState()(*RequiredWorkflow_state) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
func (m *RequiredWorkflow) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *RequiredWorkflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("path", m.GetPath())
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
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err := writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("selected_repositories_url", m.GetSelectedRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredWorkflow) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RequiredWorkflow) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetId sets the id property value. Unique identifier for a required workflow
func (m *RequiredWorkflow) SetId(value *int64)() {
    m.id = value
}
// SetName sets the name property value. Name present in the workflow file
func (m *RequiredWorkflow) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. Path of the workflow file
func (m *RequiredWorkflow) SetPath(value *string)() {
    m.path = value
}
// SetRef sets the ref property value. Ref at which the workflow file will be selected
func (m *RequiredWorkflow) SetRef(value *string)() {
    m.ref = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *RequiredWorkflow) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
// SetScope sets the scope property value. Scope of the required workflow
func (m *RequiredWorkflow) SetScope(value *RequiredWorkflow_scope)() {
    m.scope = value
}
// SetSelectedRepositoriesUrl sets the selected_repositories_url property value. The selected_repositories_url property
func (m *RequiredWorkflow) SetSelectedRepositoriesUrl(value *string)() {
    m.selected_repositories_url = value
}
// SetState sets the state property value. State of the required workflow
func (m *RequiredWorkflow) SetState(value *RequiredWorkflow_state)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *RequiredWorkflow) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
