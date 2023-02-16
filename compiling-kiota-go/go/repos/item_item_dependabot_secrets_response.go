package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemDependabotSecretsResponse 
type ItemItemDependabotSecretsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The secrets property
    secrets []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotSecretable
    // The total_count property
    total_count *int32
}
// NewItemItemDependabotSecretsResponse instantiates a new ItemItemDependabotSecretsResponse and sets the default values.
func NewItemItemDependabotSecretsResponse()(*ItemItemDependabotSecretsResponse) {
    m := &ItemItemDependabotSecretsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemDependabotSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemDependabotSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDependabotSecretsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemDependabotSecretsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemDependabotSecretsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDependabotSecretFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotSecretable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotSecretable)
            }
            m.SetSecrets(res)
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
// GetSecrets gets the secrets property value. The secrets property
func (m *ItemItemDependabotSecretsResponse) GetSecrets()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotSecretable) {
    return m.secrets
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemDependabotSecretsResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemDependabotSecretsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSecrets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecrets()))
        for i, v := range m.GetSecrets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("secrets", cast)
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
func (m *ItemItemDependabotSecretsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecrets sets the secrets property value. The secrets property
func (m *ItemItemDependabotSecretsResponse) SetSecrets(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DependabotSecretable)() {
    m.secrets = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemDependabotSecretsResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
