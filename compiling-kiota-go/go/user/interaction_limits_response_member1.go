package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InteractionLimitsResponseMember1 
type InteractionLimitsResponseMember1 struct {
}
// NewInteractionLimitsResponseMember1 instantiates a new InteractionLimitsResponseMember1 and sets the default values.
func NewInteractionLimitsResponseMember1()(*InteractionLimitsResponseMember1) {
    m := &InteractionLimitsResponseMember1{
    }
    return m
}
// CreateInteractionLimitsResponseMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInteractionLimitsResponseMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInteractionLimitsResponseMember1(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InteractionLimitsResponseMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *InteractionLimitsResponseMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
