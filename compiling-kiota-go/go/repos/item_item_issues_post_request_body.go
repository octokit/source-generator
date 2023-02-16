package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemIssuesPostRequestBody 
type ItemItemIssuesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
    assignee *string
    // Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
    assignees []string
    // The contents of the issue.
    body *string
    // Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
    labels []string
    // The milestone property
    milestone ItemItemIssuesPostRequestBody_Issuesable
    // The title of the issue.
    title ItemItemIssuesPostRequestBody_Issuesable
}
// ItemItemIssuesPostRequestBody_Issues composed type wrapper for classes string, integer
type ItemItemIssuesPostRequestBody_Issues struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type integer
    integer *int32
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewItemItemIssuesPostRequestBody_Issues instantiates a new issues and sets the default values.
func NewItemItemIssuesPostRequestBody_Issues()(*ItemItemIssuesPostRequestBody_Issues) {
    m := &ItemItemIssuesPostRequestBody_Issues{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesPostRequestBody_IssuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemIssuesPostRequestBody_IssuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemItemIssuesPostRequestBody_Issues()
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
    if val, err := parseNode.GetInt32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetInteger(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemIssuesPostRequestBody_Issues) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemIssuesPostRequestBody_Issues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteger gets the integer property value. Composed type representation for type integer
func (m *ItemItemIssuesPostRequestBody_Issues) GetInteger()(*int32) {
    return m.integer
}
// GetString gets the string property value. Composed type representation for type string
func (m *ItemItemIssuesPostRequestBody_Issues) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ItemItemIssuesPostRequestBody_Issues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteger() != nil {
        err := writer.WriteInt32Value("", m.GetInteger())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *ItemItemIssuesPostRequestBody_Issues) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInteger sets the integer property value. Composed type representation for type integer
func (m *ItemItemIssuesPostRequestBody_Issues) SetInteger(value *int32)() {
    m.integer = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemItemIssuesPostRequestBody_Issues) SetString(value *string)() {
    m.string = value
}
// ItemItemIssuesPostRequestBody_Issuesable 
type ItemItemIssuesPostRequestBody_Issuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteger()(*int32)
    GetString()(*string)
    SetInteger(value *int32)()
    SetString(value *string)()
}
// NewItemItemIssuesPostRequestBody instantiates a new ItemItemIssuesPostRequestBody and sets the default values.
func NewItemItemIssuesPostRequestBody()(*ItemItemIssuesPostRequestBody) {
    m := &ItemItemIssuesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemIssuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemIssuesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignee gets the assignee property value. Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
func (m *ItemItemIssuesPostRequestBody) GetAssignee()(*string) {
    return m.assignee
}
// GetAssignees gets the assignees property value. Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) GetAssignees()([]string) {
    return m.assignees
}
// GetBody gets the body property value. The contents of the issue.
func (m *ItemItemIssuesPostRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemIssuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignee(val)
        }
        return nil
    }
    res["assignees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignees(res)
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["milestone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemIssuesPostRequestBody_IssuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestone(val.(ItemItemIssuesPostRequestBody_Issuesable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemIssuesPostRequestBody_IssuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(ItemItemIssuesPostRequestBody_Issuesable))
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) GetLabels()([]string) {
    return m.labels
}
// GetMilestone gets the milestone property value. The milestone property
func (m *ItemItemIssuesPostRequestBody) GetMilestone()(ItemItemIssuesPostRequestBody_Issuesable) {
    return m.milestone
}
// GetTitle gets the title property value. The title of the issue.
func (m *ItemItemIssuesPostRequestBody) GetTitle()(ItemItemIssuesPostRequestBody_Issuesable) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemItemIssuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignee", m.GetAssignee())
        if err != nil {
            return err
        }
    }
    if m.GetAssignees() != nil {
        err := writer.WriteCollectionOfStringValues("assignees", m.GetAssignees())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("milestone", m.GetMilestone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("title", m.GetTitle())
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
func (m *ItemItemIssuesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignee sets the assignee property value. Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
func (m *ItemItemIssuesPostRequestBody) SetAssignee(value *string)() {
    m.assignee = value
}
// SetAssignees sets the assignees property value. Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) SetAssignees(value []string)() {
    m.assignees = value
}
// SetBody sets the body property value. The contents of the issue.
func (m *ItemItemIssuesPostRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetLabels sets the labels property value. Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
func (m *ItemItemIssuesPostRequestBody) SetLabels(value []string)() {
    m.labels = value
}
// SetMilestone sets the milestone property value. The milestone property
func (m *ItemItemIssuesPostRequestBody) SetMilestone(value ItemItemIssuesPostRequestBody_Issuesable)() {
    m.milestone = value
}
// SetTitle sets the title property value. The title of the issue.
func (m *ItemItemIssuesPostRequestBody) SetTitle(value ItemItemIssuesPostRequestBody_Issuesable)() {
    m.title = value
}
