package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// RepositoriesResponse 
type RepositoriesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The incomplete_results property
    incomplete_results *bool
    // The items property
    items []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoSearchResultItemable
    // The total_count property
    total_count *int32
}
// NewRepositoriesResponse instantiates a new RepositoriesResponse and sets the default values.
func NewRepositoriesResponse()(*RepositoriesResponse) {
    m := &RepositoriesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoriesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoriesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoriesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["incomplete_results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteResults(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRepoSearchResultItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoSearchResultItemable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoSearchResultItemable)
            }
            m.SetItems(res)
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
// GetIncompleteResults gets the incomplete_results property value. The incomplete_results property
func (m *RepositoriesResponse) GetIncompleteResults()(*bool) {
    return m.incomplete_results
}
// GetItems gets the items property value. The items property
func (m *RepositoriesResponse) GetItems()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoSearchResultItemable) {
    return m.items
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *RepositoriesResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *RepositoriesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("incomplete_results", m.GetIncompleteResults())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
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
func (m *RepositoriesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncompleteResults sets the incomplete_results property value. The incomplete_results property
func (m *RepositoriesResponse) SetIncompleteResults(value *bool)() {
    m.incomplete_results = value
}
// SetItems sets the items property value. The items property
func (m *RepositoriesResponse) SetItems(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.RepoSearchResultItemable)() {
    m.items = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *RepositoriesResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
