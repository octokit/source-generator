package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Pageable 
type Pageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuildType()(*Page_build_type)
    GetCname()(*string)
    GetCustom404()(*bool)
    GetHtmlUrl()(*string)
    GetHttpsCertificate()(PagesHttpsCertificateable)
    GetHttpsEnforced()(*bool)
    GetPendingDomainUnverifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProtectedDomainState()(*Page_protected_domain_state)
    GetPublic()(*bool)
    GetSource()(PagesSourceHashable)
    GetStatus()(*Page_status)
    GetUrl()(*string)
    SetBuildType(value *Page_build_type)()
    SetCname(value *string)()
    SetCustom404(value *bool)()
    SetHtmlUrl(value *string)()
    SetHttpsCertificate(value PagesHttpsCertificateable)()
    SetHttpsEnforced(value *bool)()
    SetPendingDomainUnverifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProtectedDomainState(value *Page_protected_domain_state)()
    SetPublic(value *bool)()
    SetSource(value PagesSourceHashable)()
    SetStatus(value *Page_status)()
    SetUrl(value *string)()
}
