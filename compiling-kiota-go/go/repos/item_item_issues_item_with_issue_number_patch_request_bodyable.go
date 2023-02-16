package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemIssuesItemWithIssue_numberPatchRequestBodyable 
type ItemItemIssuesItemWithIssue_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignee()(*string)
    GetAssignees()([]string)
    GetBody()(*string)
    GetLabels()([]string)
    GetMilestone()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberable)
    GetTitle()(ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberable)
    SetAssignee(value *string)()
    SetAssignees(value []string)()
    SetBody(value *string)()
    SetLabels(value []string)()
    SetMilestone(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberable)()
    SetTitle(value ItemItemIssuesItemWithIssue_numberPatchRequestBody_WithIssue_numberable)()
}
