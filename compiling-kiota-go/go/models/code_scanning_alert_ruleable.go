package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAlertRuleable 
type CodeScanningAlertRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFullDescription()(*string)
    GetHelp()(*string)
    GetHelpUri()(*string)
    GetId()(*string)
    GetName()(*string)
    GetSecuritySeverityLevel()(*CodeScanningAlertRule_security_severity_level)
    GetSeverity()(*CodeScanningAlertRule_severity)
    GetTags()([]string)
    SetDescription(value *string)()
    SetFullDescription(value *string)()
    SetHelp(value *string)()
    SetHelpUri(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetSecuritySeverityLevel(value *CodeScanningAlertRule_security_severity_level)()
    SetSeverity(value *CodeScanningAlertRule_severity)()
    SetTags(value []string)()
}
