package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOutside_collaboratorsItemWithUsernameResponse 
type ItemOutside_collaboratorsItemWithUsernameResponse struct {
}
// NewItemOutside_collaboratorsItemWithUsernameResponse instantiates a new ItemOutside_collaboratorsItemWithUsernameResponse and sets the default values.
func NewItemOutside_collaboratorsItemWithUsernameResponse()(*ItemOutside_collaboratorsItemWithUsernameResponse) {
    m := &ItemOutside_collaboratorsItemWithUsernameResponse{
    }
    return m
}
// CreateItemOutside_collaboratorsItemWithUsernameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOutside_collaboratorsItemWithUsernameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOutside_collaboratorsItemWithUsernameResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOutside_collaboratorsItemWithUsernameResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemOutside_collaboratorsItemWithUsernameResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
