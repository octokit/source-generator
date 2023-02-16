package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddedToProjectIssueEvent_project_cardable 
type AddedToProjectIssueEvent_project_cardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumnName()(*string)
    GetId()(*int32)
    GetPreviousColumnName()(*string)
    GetProjectId()(*int32)
    GetProjectUrl()(*string)
    GetUrl()(*string)
    SetColumnName(value *string)()
    SetId(value *int32)()
    SetPreviousColumnName(value *string)()
    SetProjectId(value *int32)()
    SetProjectUrl(value *string)()
    SetUrl(value *string)()
}
