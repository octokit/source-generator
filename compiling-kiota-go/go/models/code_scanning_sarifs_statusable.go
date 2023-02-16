package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningSarifsStatusable 
type CodeScanningSarifsStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysesUrl()(*string)
    GetErrors()([]string)
    GetProcessingStatus()(*CodeScanningSarifsStatus_processing_status)
    SetAnalysesUrl(value *string)()
    SetErrors(value []string)()
    SetProcessingStatus(value *CodeScanningSarifsStatus_processing_status)()
}
