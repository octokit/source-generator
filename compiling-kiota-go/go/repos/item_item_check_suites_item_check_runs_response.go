package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCheckSuitesItemCheckRunsResponse 
type ItemItemCheckSuitesItemCheckRunsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_runs property
    check_runs []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable
    // The total_count property
    total_count *int32
}
// NewItemItemCheckSuitesItemCheckRunsResponse instantiates a new ItemItemCheckSuitesItemCheckRunsResponse and sets the default values.
func NewItemItemCheckSuitesItemCheckRunsResponse()(*ItemItemCheckSuitesItemCheckRunsResponse) {
    m := &ItemItemCheckSuitesItemCheckRunsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCheckSuitesItemCheckRunsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCheckSuitesItemCheckRunsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCheckSuitesItemCheckRunsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckRuns gets the check_runs property value. The check_runs property
func (m *ItemItemCheckSuitesItemCheckRunsResponse) GetCheckRuns()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable) {
    return m.check_runs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCheckSuitesItemCheckRunsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCheckRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable)
            }
            m.SetCheckRuns(res)
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
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemCheckSuitesItemCheckRunsResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCheckSuitesItemCheckRunsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckRuns()))
        for i, v := range m.GetCheckRuns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("check_runs", cast)
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
func (m *ItemItemCheckSuitesItemCheckRunsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckRuns sets the check_runs property value. The check_runs property
func (m *ItemItemCheckSuitesItemCheckRunsResponse) SetCheckRuns(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable)() {
    m.check_runs = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCheckSuitesItemCheckRunsResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
