package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRequired_workflowsPostRequestBody 
type ItemActionsRequired_workflowsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the repository that contains the workflow file.
    repository_id *string
    // A list of repository IDs where you want to enable the required workflow. You can only provide a list of repository ids when the `scope` is set to `selected`.
    selected_repository_ids []int32
    // Path of the workflow file to be configured as a required workflow.
    workflow_file_path *string
}
// NewItemActionsRequired_workflowsPostRequestBody instantiates a new ItemActionsRequired_workflowsPostRequestBody and sets the default values.
func NewItemActionsRequired_workflowsPostRequestBody()(*ItemActionsRequired_workflowsPostRequestBody) {
    m := &ItemActionsRequired_workflowsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRequired_workflowsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRequired_workflowsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRequired_workflowsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsRequired_workflowsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsRequired_workflowsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["selected_repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetSelectedRepositoryIds(res)
        }
        return nil
    }
    res["workflow_file_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowFilePath(val)
        }
        return nil
    }
    return res
}
// GetRepositoryId gets the repository_id property value. The ID of the repository that contains the workflow file.
func (m *ItemActionsRequired_workflowsPostRequestBody) GetRepositoryId()(*string) {
    return m.repository_id
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. A list of repository IDs where you want to enable the required workflow. You can only provide a list of repository ids when the `scope` is set to `selected`.
func (m *ItemActionsRequired_workflowsPostRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// GetWorkflowFilePath gets the workflow_file_path property value. Path of the workflow file to be configured as a required workflow.
func (m *ItemActionsRequired_workflowsPostRequestBody) GetWorkflowFilePath()(*string) {
    return m.workflow_file_path
}
// Serialize serializes information the current object
func (m *ItemActionsRequired_workflowsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("workflow_file_path", m.GetWorkflowFilePath())
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
func (m *ItemActionsRequired_workflowsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositoryId sets the repository_id property value. The ID of the repository that contains the workflow file.
func (m *ItemActionsRequired_workflowsPostRequestBody) SetRepositoryId(value *string)() {
    m.repository_id = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. A list of repository IDs where you want to enable the required workflow. You can only provide a list of repository ids when the `scope` is set to `selected`.
func (m *ItemActionsRequired_workflowsPostRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
// SetWorkflowFilePath sets the workflow_file_path property value. Path of the workflow file to be configured as a required workflow.
func (m *ItemActionsRequired_workflowsPostRequestBody) SetWorkflowFilePath(value *string)() {
    m.workflow_file_path = value
}
