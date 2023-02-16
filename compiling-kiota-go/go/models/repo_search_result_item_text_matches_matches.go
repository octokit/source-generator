package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepoSearchResultItem_text_matches_matches 
type RepoSearchResultItem_text_matches_matches struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The indices property
    indices []int32
    // The text property
    text *string
}
// NewRepoSearchResultItem_text_matches_matches instantiates a new repoSearchResultItem_text_matches_matches and sets the default values.
func NewRepoSearchResultItem_text_matches_matches()(*RepoSearchResultItem_text_matches_matches) {
    m := &RepoSearchResultItem_text_matches_matches{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepoSearchResultItem_text_matches_matchesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepoSearchResultItem_text_matches_matchesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepoSearchResultItem_text_matches_matches(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepoSearchResultItem_text_matches_matches) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepoSearchResultItem_text_matches_matches) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["indices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetIndices(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetIndices gets the indices property value. The indices property
func (m *RepoSearchResultItem_text_matches_matches) GetIndices()([]int32) {
    return m.indices
}
// GetText gets the text property value. The text property
func (m *RepoSearchResultItem_text_matches_matches) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *RepoSearchResultItem_text_matches_matches) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIndices() != nil {
        err := writer.WriteCollectionOfInt32Values("indices", m.GetIndices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *RepoSearchResultItem_text_matches_matches) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndices sets the indices property value. The indices property
func (m *RepoSearchResultItem_text_matches_matches) SetIndices(value []int32)() {
    m.indices = value
}
// SetText sets the text property value. The text property
func (m *RepoSearchResultItem_text_matches_matches) SetText(value *string)() {
    m.text = value
}
