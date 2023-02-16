package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesItemWithCodespace_nameResponse 
type CodespacesItemWithCodespace_nameResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewCodespacesItemWithCodespace_nameResponse instantiates a new CodespacesItemWithCodespace_nameResponse and sets the default values.
func NewCodespacesItemWithCodespace_nameResponse()(*CodespacesItemWithCodespace_nameResponse) {
    m := &CodespacesItemWithCodespace_nameResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemWithCodespace_nameResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesItemWithCodespace_nameResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesItemWithCodespace_nameResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *CodespacesItemWithCodespace_nameResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesItemWithCodespace_nameResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
