package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAnalysisable 
type CodeScanningAnalysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisKey()(*string)
    GetCategory()(*string)
    GetCommitSha()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletable()(*bool)
    GetEnvironment()(*string)
    GetError()(*string)
    GetId()(*int32)
    GetRef()(*string)
    GetResultsCount()(*int32)
    GetRulesCount()(*int32)
    GetSarifId()(*string)
    GetTool()(CodeScanningAnalysisToolable)
    GetUrl()(*string)
    GetWarning()(*string)
    SetAnalysisKey(value *string)()
    SetCategory(value *string)()
    SetCommitSha(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletable(value *bool)()
    SetEnvironment(value *string)()
    SetError(value *string)()
    SetId(value *int32)()
    SetRef(value *string)()
    SetResultsCount(value *int32)()
    SetRulesCount(value *int32)()
    SetSarifId(value *string)()
    SetTool(value CodeScanningAnalysisToolable)()
    SetUrl(value *string)()
    SetWarning(value *string)()
}
