package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAlertInstanceable 
type CodeScanningAlertInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisKey()(*string)
    GetCategory()(*string)
    GetClassifications()([]CodeScanningAlertClassification)
    GetCommitSha()(*string)
    GetEnvironment()(*string)
    GetHtmlUrl()(*string)
    GetLocation()(CodeScanningAlertLocationable)
    GetMessage()(CodeScanningAlertInstance_messageable)
    GetRef()(*string)
    GetState()(*CodeScanningAlertState)
    SetAnalysisKey(value *string)()
    SetCategory(value *string)()
    SetClassifications(value []CodeScanningAlertClassification)()
    SetCommitSha(value *string)()
    SetEnvironment(value *string)()
    SetHtmlUrl(value *string)()
    SetLocation(value CodeScanningAlertLocationable)()
    SetMessage(value CodeScanningAlertInstance_messageable)()
    SetRef(value *string)()
    SetState(value *CodeScanningAlertState)()
}
