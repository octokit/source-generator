package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCommitsItemCheckSuitesResponse 
type ItemItemCommitsItemCheckSuitesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The check_suites property
    check_suites []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable
    // The total_count property
    total_count *int32
}
// NewItemItemCommitsItemCheckSuitesResponse instantiates a new ItemItemCommitsItemCheckSuitesResponse and sets the default values.
func NewItemItemCommitsItemCheckSuitesResponse()(*ItemItemCommitsItemCheckSuitesResponse) {
    m := &ItemItemCommitsItemCheckSuitesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCommitsItemCheckSuitesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCommitsItemCheckSuitesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckSuites gets the check_suites property value. The check_suites property
func (m *ItemItemCommitsItemCheckSuitesResponse) GetCheckSuites()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable) {
    return m.check_suites
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCommitsItemCheckSuitesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["check_suites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCheckSuiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable)
            }
            m.SetCheckSuites(res)
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
func (m *ItemItemCommitsItemCheckSuitesResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCommitsItemCheckSuitesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckSuites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckSuites()))
        for i, v := range m.GetCheckSuites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("check_suites", cast)
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
func (m *ItemItemCommitsItemCheckSuitesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckSuites sets the check_suites property value. The check_suites property
func (m *ItemItemCommitsItemCheckSuitesResponse) SetCheckSuites(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckSuiteable)() {
    m.check_suites = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCommitsItemCheckSuitesResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
