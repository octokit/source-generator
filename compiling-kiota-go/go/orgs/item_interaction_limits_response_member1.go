package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInteractionLimitsResponseMember1 
type ItemInteractionLimitsResponseMember1 struct {
}
// NewItemInteractionLimitsResponseMember1 instantiates a new ItemInteractionLimitsResponseMember1 and sets the default values.
func NewItemInteractionLimitsResponseMember1()(*ItemInteractionLimitsResponseMember1) {
    m := &ItemInteractionLimitsResponseMember1{
    }
    return m
}
// CreateItemInteractionLimitsResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInteractionLimitsResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInteractionLimitsResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInteractionLimitsResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemInteractionLimitsResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
