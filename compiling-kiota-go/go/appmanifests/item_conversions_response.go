package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemConversionsResponse 
type ItemConversionsResponse struct {
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Integration
}
// NewItemConversionsResponse instantiates a new ItemConversionsResponse and sets the default values.
func NewItemConversionsResponse()(*ItemConversionsResponse) {
    m := &ItemConversionsResponse{
        Integration: *ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.NewIntegration(),
    }
    return m
}
// CreateItemConversionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemConversionsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Integration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ItemConversionsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Integration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
