package repositories

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemEnvironmentsItemVariablesResponse 
type ItemEnvironmentsItemVariablesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_count property
    total_count *int32
    // The variables property
    variables []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable
}
// NewItemEnvironmentsItemVariablesResponse instantiates a new ItemEnvironmentsItemVariablesResponse and sets the default values.
func NewItemEnvironmentsItemVariablesResponse()(*ItemEnvironmentsItemVariablesResponse) {
    m := &ItemEnvironmentsItemVariablesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEnvironmentsItemVariablesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEnvironmentsItemVariablesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemEnvironmentsItemVariablesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["variables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateActionsVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable)
            }
            m.SetVariables(res)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemEnvironmentsItemVariablesResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// GetVariables gets the variables property value. The variables property
func (m *ItemEnvironmentsItemVariablesResponse) GetVariables()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable) {
    return m.variables
}
// Serialize serializes information the current object
func (m *ItemEnvironmentsItemVariablesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    if m.GetVariables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVariables()))
        for i, v := range m.GetVariables() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("variables", cast)
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
func (m *ItemEnvironmentsItemVariablesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemEnvironmentsItemVariablesResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// SetVariables sets the variables property value. The variables property
func (m *ItemEnvironmentsItemVariablesResponse) SetVariables(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsVariableable)() {
    m.variables = value
}
