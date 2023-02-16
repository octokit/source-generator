package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionsRequired_workflowsItemRepositoriesPutRequestBody 
type ItemActionsRequired_workflowsItemRepositoriesPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The IDs of the repositories for which the workflow should be required.
    selected_repository_ids []int32
}
// NewItemActionsRequired_workflowsItemRepositoriesPutRequestBody instantiates a new ItemActionsRequired_workflowsItemRepositoriesPutRequestBody and sets the default values.
func NewItemActionsRequired_workflowsItemRepositoriesPutRequestBody()(*ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) {
    m := &ItemActionsRequired_workflowsItemRepositoriesPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRequired_workflowsItemRepositoriesPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsRequired_workflowsItemRepositoriesPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRequired_workflowsItemRepositoriesPutRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. The IDs of the repositories for which the workflow should be required.
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// Serialize serializes information the current object
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
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
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. The IDs of the repositories for which the workflow should be required.
func (m *ItemActionsRequired_workflowsItemRepositoriesPutRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
