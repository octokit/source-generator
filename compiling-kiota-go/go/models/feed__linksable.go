package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Feed__linksable 
type Feed__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUser()(LinkWithTypeable)
    GetCurrentUserActor()(LinkWithTypeable)
    GetCurrentUserOrganization()(LinkWithTypeable)
    GetCurrentUserOrganizations()([]LinkWithTypeable)
    GetCurrentUserPublic()(LinkWithTypeable)
    GetRepositoryDiscussions()(LinkWithTypeable)
    GetRepositoryDiscussionsCategory()(LinkWithTypeable)
    GetSecurityAdvisories()(LinkWithTypeable)
    GetTimeline()(LinkWithTypeable)
    GetUser()(LinkWithTypeable)
    SetCurrentUser(value LinkWithTypeable)()
    SetCurrentUserActor(value LinkWithTypeable)()
    SetCurrentUserOrganization(value LinkWithTypeable)()
    SetCurrentUserOrganizations(value []LinkWithTypeable)()
    SetCurrentUserPublic(value LinkWithTypeable)()
    SetRepositoryDiscussions(value LinkWithTypeable)()
    SetRepositoryDiscussionsCategory(value LinkWithTypeable)()
    SetSecurityAdvisories(value LinkWithTypeable)()
    SetTimeline(value LinkWithTypeable)()
    SetUser(value LinkWithTypeable)()
}
