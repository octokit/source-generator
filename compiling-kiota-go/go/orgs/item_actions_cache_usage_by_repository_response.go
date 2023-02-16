package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemActionsCacheUsageByRepositoryResponse 
type ItemActionsCacheUsageByRepositoryResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository_cache_usages property
    repository_cache_usages []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsCacheUsageByRepositoryable
    // The total_count property
    total_count *int32
}
// NewItemActionsCacheUsageByRepositoryResponse instantiates a new ItemActionsCacheUsageByRepositoryResponse and sets the default values.
func NewItemActionsCacheUsageByRepositoryResponse()(*ItemActionsCacheUsageByRepositoryResponse) {
    m := &ItemActionsCacheUsageByRepositoryResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionsCacheUsageByRepositoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsCacheUsageByRepositoryResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionsCacheUsageByRepositoryResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionsCacheUsageByRepositoryResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repository_cache_usages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateActionsCacheUsageByRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsCacheUsageByRepositoryable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsCacheUsageByRepositoryable)
            }
            m.SetRepositoryCacheUsages(res)
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
// GetRepositoryCacheUsages gets the repository_cache_usages property value. The repository_cache_usages property
func (m *ItemActionsCacheUsageByRepositoryResponse) GetRepositoryCacheUsages()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsCacheUsageByRepositoryable) {
    return m.repository_cache_usages
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemActionsCacheUsageByRepositoryResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemActionsCacheUsageByRepositoryResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositoryCacheUsages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositoryCacheUsages()))
        for i, v := range m.GetRepositoryCacheUsages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("repository_cache_usages", cast)
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
func (m *ItemActionsCacheUsageByRepositoryResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositoryCacheUsages sets the repository_cache_usages property value. The repository_cache_usages property
func (m *ItemActionsCacheUsageByRepositoryResponse) SetRepositoryCacheUsages(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ActionsCacheUsageByRepositoryable)() {
    m.repository_cache_usages = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemActionsCacheUsageByRepositoryResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
