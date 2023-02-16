package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemMovesResponse 
type ColumnsCardsItemMovesResponse struct {
}
// NewColumnsCardsItemMovesResponse instantiates a new ColumnsCardsItemMovesResponse and sets the default values.
func NewColumnsCardsItemMovesResponse()(*ColumnsCardsItemMovesResponse) {
    m := &ColumnsCardsItemMovesResponse{
    }
    return m
}
// CreateColumnsCardsItemMovesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnsCardsItemMovesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemMovesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnsCardsItemMovesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ColumnsCardsItemMovesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
