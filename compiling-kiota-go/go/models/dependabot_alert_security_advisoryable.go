package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisoryable 
type DependabotAlertSecurityAdvisoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCveId()(*string)
    GetCvss()(DependabotAlertSecurityAdvisory_cvssable)
    GetCwes()([]DependabotAlertSecurityAdvisory_cwesable)
    GetDescription()(*string)
    GetGhsaId()(*string)
    GetIdentifiers()([]DependabotAlertSecurityAdvisory_identifiersable)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferences()([]DependabotAlertSecurityAdvisory_referencesable)
    GetSeverity()(*DependabotAlertSecurityAdvisory_severity)
    GetSummary()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVulnerabilities()([]DependabotAlertSecurityVulnerabilityable)
    GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCveId(value *string)()
    SetCvss(value DependabotAlertSecurityAdvisory_cvssable)()
    SetCwes(value []DependabotAlertSecurityAdvisory_cwesable)()
    SetDescription(value *string)()
    SetGhsaId(value *string)()
    SetIdentifiers(value []DependabotAlertSecurityAdvisory_identifiersable)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferences(value []DependabotAlertSecurityAdvisory_referencesable)()
    SetSeverity(value *DependabotAlertSecurityAdvisory_severity)()
    SetSummary(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVulnerabilities(value []DependabotAlertSecurityVulnerabilityable)()
    SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
