package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesMember1 update the source for the repository. Must include the branch name and path.
type ItemItemPagesMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository branch used to publish your site's source files.
    branch *string
}
// NewItemItemPagesMember1 instantiates a new ItemItemPagesMember1 and sets the default values.
func NewItemItemPagesMember1()(*ItemItemPagesMember1) {
    m := &ItemItemPagesMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPagesMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPagesMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the branch property value. The repository branch used to publish your site's source files.
func (m *ItemItemPagesMember1) GetBranch()(*string) {
    return m.branch
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemPagesMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemPagesMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
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
func (m *ItemItemPagesMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the branch property value. The repository branch used to publish your site's source files.
func (m *ItemItemPagesMember1) SetBranch(value *string)() {
    m.branch = value
}
