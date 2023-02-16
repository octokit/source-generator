package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocation 
type SecretScanningLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details property
    details SecretScanningLocation_Locationsable
    // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
    typeEscaped *SecretScanningLocation_type
}
// SecretScanningLocation_Locations composed type wrapper for classes secretScanningLocationCommit, secretScanningLocationIssueTitle, secretScanningLocationIssueBody, secretScanningLocationIssueComment
type SecretScanningLocation_Locations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type secretScanningLocationCommit
    secretScanningLocationCommit SecretScanningLocationCommitable
    // Composed type representation for type secretScanningLocationIssueBody
    secretScanningLocationIssueBody SecretScanningLocationIssueBodyable
    // Composed type representation for type secretScanningLocationIssueComment
    secretScanningLocationIssueComment SecretScanningLocationIssueCommentable
    // Composed type representation for type secretScanningLocationIssueTitle
    secretScanningLocationIssueTitle SecretScanningLocationIssueTitleable
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewSecretScanningLocation_Locations instantiates a new locations and sets the default values.
func NewSecretScanningLocation_Locations()(*SecretScanningLocation_Locations) {
    m := &SecretScanningLocation_Locations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocation_LocationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecretScanningLocation_LocationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecretScanningLocation_Locations()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecretScanningLocation_Locations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecretScanningLocation_Locations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetSecretScanningLocationCommit gets the secretScanningLocationCommit property value. Composed type representation for type secretScanningLocationCommit
func (m *SecretScanningLocation_Locations) GetSecretScanningLocationCommit()(SecretScanningLocationCommitable) {
    return m.secretScanningLocationCommit
}
// GetSecretScanningLocationIssueBody gets the secretScanningLocationIssueBody property value. Composed type representation for type secretScanningLocationIssueBody
func (m *SecretScanningLocation_Locations) GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable) {
    return m.secretScanningLocationIssueBody
}
// GetSecretScanningLocationIssueComment gets the secretScanningLocationIssueComment property value. Composed type representation for type secretScanningLocationIssueComment
func (m *SecretScanningLocation_Locations) GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable) {
    return m.secretScanningLocationIssueComment
}
// GetSecretScanningLocationIssueTitle gets the secretScanningLocationIssueTitle property value. Composed type representation for type secretScanningLocationIssueTitle
func (m *SecretScanningLocation_Locations) GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable) {
    return m.secretScanningLocationIssueTitle
}
// Serialize serializes information the current object
func (m *SecretScanningLocation_Locations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSecretScanningLocationCommit() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationCommit())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueBody() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueBody())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueTitle() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueTitle())
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
func (m *SecretScanningLocation_Locations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecretScanningLocationCommit sets the secretScanningLocationCommit property value. Composed type representation for type secretScanningLocationCommit
func (m *SecretScanningLocation_Locations) SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)() {
    m.secretScanningLocationCommit = value
}
// SetSecretScanningLocationIssueBody sets the secretScanningLocationIssueBody property value. Composed type representation for type secretScanningLocationIssueBody
func (m *SecretScanningLocation_Locations) SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)() {
    m.secretScanningLocationIssueBody = value
}
// SetSecretScanningLocationIssueComment sets the secretScanningLocationIssueComment property value. Composed type representation for type secretScanningLocationIssueComment
func (m *SecretScanningLocation_Locations) SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)() {
    m.secretScanningLocationIssueComment = value
}
// SetSecretScanningLocationIssueTitle sets the secretScanningLocationIssueTitle property value. Composed type representation for type secretScanningLocationIssueTitle
func (m *SecretScanningLocation_Locations) SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)() {
    m.secretScanningLocationIssueTitle = value
}
// SecretScanningLocation_Locationsable 
type SecretScanningLocation_Locationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecretScanningLocationCommit()(SecretScanningLocationCommitable)
    GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable)
    GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable)
    GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable)
    SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)()
    SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)()
    SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)()
    SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)()
}
// NewSecretScanningLocation instantiates a new SecretScanningLocation and sets the default values.
func NewSecretScanningLocation()(*SecretScanningLocation) {
    m := &SecretScanningLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecretScanningLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecretScanningLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. The details property
func (m *SecretScanningLocation) GetDetails()(SecretScanningLocation_Locationsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecretScanningLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecretScanningLocation_LocationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(SecretScanningLocation_Locationsable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningLocation_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*SecretScanningLocation_type))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
func (m *SecretScanningLocation) GetType()(*SecretScanningLocation_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SecretScanningLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *SecretScanningLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. The details property
func (m *SecretScanningLocation) SetDetails(value SecretScanningLocation_Locationsable)() {
    m.details = value
}
// SetType sets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
func (m *SecretScanningLocation) SetType(value *SecretScanningLocation_type)() {
    m.typeEscaped = value
}
