package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRequired_workflowsResponse 
type ItemItemActionsRequired_workflowsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The required_workflows property
    required_workflows []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable
    // The total_count property
    total_count *int32
}
// NewItemItemActionsRequired_workflowsResponse instantiates a new ItemItemActionsRequired_workflowsResponse and sets the default values.
func NewItemItemActionsRequired_workflowsResponse()(*ItemItemActionsRequired_workflowsResponse) {
    m := &ItemItemActionsRequired_workflowsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRequired_workflowsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRequired_workflowsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRequired_workflowsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsRequired_workflowsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsRequired_workflowsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["required_workflows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRepoRequiredWorkflowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable)
            }
            m.SetRequiredWorkflows(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetRequiredWorkflows gets the required_workflows property value. The required_workflows property
func (m *ItemItemActionsRequired_workflowsResponse) GetRequiredWorkflows()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable) {
    return m.required_workflows
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemActionsRequired_workflowsResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemActionsRequired_workflowsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRequiredWorkflows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequiredWorkflows()))
        for i, v := range m.GetRequiredWorkflows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("required_workflows", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemActionsRequired_workflowsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequiredWorkflows sets the required_workflows property value. The required_workflows property
func (m *ItemItemActionsRequired_workflowsResponse) SetRequiredWorkflows(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoRequiredWorkflowable)() {
    m.required_workflows = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsRequired_workflowsResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
