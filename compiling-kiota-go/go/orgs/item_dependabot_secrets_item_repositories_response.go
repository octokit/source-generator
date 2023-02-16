package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemDependabotSecretsItemRepositoriesResponse 
type ItemDependabotSecretsItemRepositoriesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repositories property
    repositories []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MinimalRepositoryable
    // The total_count property
    total_count *int32
}
// NewItemDependabotSecretsItemRepositoriesResponse instantiates a new ItemDependabotSecretsItemRepositoriesResponse and sets the default values.
func NewItemDependabotSecretsItemRepositoriesResponse()(*ItemDependabotSecretsItemRepositoriesResponse) {
    m := &ItemDependabotSecretsItemRepositoriesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemDependabotSecretsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDependabotSecretsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDependabotSecretsItemRepositoriesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemDependabotSecretsItemRepositoriesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemDependabotSecretsItemRepositoriesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MinimalRepositoryable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MinimalRepositoryable)
            }
            m.SetRepositories(res)
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
// GetRepositories gets the repositories property value. The repositories property
func (m *ItemDependabotSecretsItemRepositoriesResponse) GetRepositories()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MinimalRepositoryable) {
    return m.repositories
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemDependabotSecretsItemRepositoriesResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemDependabotSecretsItemRepositoriesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
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
func (m *ItemDependabotSecretsItemRepositoriesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. The repositories property
func (m *ItemDependabotSecretsItemRepositoriesResponse) SetRepositories(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.MinimalRepositoryable)() {
    m.repositories = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemDependabotSecretsItemRepositoriesResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
