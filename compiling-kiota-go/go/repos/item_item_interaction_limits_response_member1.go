package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemInteractionLimitsResponseMember1 
type ItemItemInteractionLimitsResponseMember1 struct {
}
// NewItemItemInteractionLimitsResponseMember1 instantiates a new ItemItemInteractionLimitsResponseMember1 and sets the default values.
func NewItemItemInteractionLimitsResponseMember1()(*ItemItemInteractionLimitsResponseMember1) {
    m := &ItemItemInteractionLimitsResponseMember1{
    }
    return m
}
// CreateItemItemInteractionLimitsResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemInteractionLimitsResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemInteractionLimitsResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemInteractionLimitsResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemInteractionLimitsResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
