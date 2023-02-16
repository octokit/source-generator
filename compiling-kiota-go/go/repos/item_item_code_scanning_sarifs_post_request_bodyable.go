package repos

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemCodeScanningSarifsPostRequestBodyable 
type ItemItemCodeScanningSarifsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckoutUri()(*string)
    GetCommitSha()(*string)
    GetRef()(*string)
    GetSarif()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetToolName()(*string)
    GetValidate()(*bool)
    SetCheckoutUri(value *string)()
    SetCommitSha(value *string)()
    SetRef(value *string)()
    SetSarif(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetToolName(value *string)()
    SetValidate(value *bool)()
}
