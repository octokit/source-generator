package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertable 
type DependabotAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDependency()(DependabotAlert_dependencyable)
    GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedBy()(NullableSimpleUserable)
    GetDismissedComment()(*string)
    GetDismissedReason()(*DependabotAlert_dismissed_reason)
    GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetNumber()(*int32)
    GetSecurityAdvisory()(DependabotAlertSecurityAdvisoryable)
    GetSecurityVulnerability()(DependabotAlertSecurityVulnerabilityable)
    GetState()(*DependabotAlert_state)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDependency(value DependabotAlert_dependencyable)()
    SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedBy(value NullableSimpleUserable)()
    SetDismissedComment(value *string)()
    SetDismissedReason(value *DependabotAlert_dismissed_reason)()
    SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetNumber(value *int32)()
    SetSecurityAdvisory(value DependabotAlertSecurityAdvisoryable)()
    SetSecurityVulnerability(value DependabotAlertSecurityVulnerabilityable)()
    SetState(value *DependabotAlert_state)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
