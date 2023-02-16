package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepoRequiredWorkflow a GitHub Actions required workflow
type RepoRequiredWorkflow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The badge_url property
    badge_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The path property
    path *string
    // Minimal Repository
    source_repository MinimalRepositoryable
    // The state property
    state *RepoRequiredWorkflow_state
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
}
// NewRepoRequiredWorkflow instantiates a new repoRequiredWorkflow and sets the default values.
func NewRepoRequiredWorkflow()(*RepoRequiredWorkflow) {
    m := &RepoRequiredWorkflow{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepoRequiredWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepoRequiredWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepoRequiredWorkflow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepoRequiredWorkflow) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBadgeUrl gets the badge_url property value. The badge_url property
func (m *RepoRequiredWorkflow) GetBadgeUrl()(*string) {
    return m.badge_url
}
// GetCreatedAt gets the created_at property value. The created_at property
func (m *RepoRequiredWorkflow) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepoRequiredWorkflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["badge_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBadgeUrl(val)
        }
        return nil
    }
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
    res["source_repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepoRequiredWorkflow_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*RepoRequiredWorkflow_state))
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
func (m *RepoRequiredWorkflow) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
func (m *RepoRequiredWorkflow) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
func (m *RepoRequiredWorkflow) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
func (m *RepoRequiredWorkflow) GetNodeId()(*string) {
    return m.node_id
}
// GetPath gets the path property value. The path property
func (m *RepoRequiredWorkflow) GetPath()(*string) {
    return m.path
}
// GetSourceRepository gets the source_repository property value. Minimal Repository
func (m *RepoRequiredWorkflow) GetSourceRepository()(MinimalRepositoryable) {
    return m.source_repository
}
// GetState gets the state property value. The state property
func (m *RepoRequiredWorkflow) GetState()(*RepoRequiredWorkflow_state) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
func (m *RepoRequiredWorkflow) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
func (m *RepoRequiredWorkflow) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *RepoRequiredWorkflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("badge_url", m.GetBadgeUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source_repository", m.GetSourceRepository())
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
func (m *RepoRequiredWorkflow) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBadgeUrl sets the badge_url property value. The badge_url property
func (m *RepoRequiredWorkflow) SetBadgeUrl(value *string)() {
    m.badge_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RepoRequiredWorkflow) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *RepoRequiredWorkflow) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *RepoRequiredWorkflow) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *RepoRequiredWorkflow) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *RepoRequiredWorkflow) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPath sets the path property value. The path property
func (m *RepoRequiredWorkflow) SetPath(value *string)() {
    m.path = value
}
// SetSourceRepository sets the source_repository property value. Minimal Repository
func (m *RepoRequiredWorkflow) SetSourceRepository(value MinimalRepositoryable)() {
    m.source_repository = value
}
// SetState sets the state property value. The state property
func (m *RepoRequiredWorkflow) SetState(value *RepoRequiredWorkflow_state)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *RepoRequiredWorkflow) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *RepoRequiredWorkflow) SetUrl(value *string)() {
    m.url = value
}
