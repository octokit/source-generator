package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Feedable 
type Feedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUserActorUrl()(*string)
    GetCurrentUserOrganizationUrl()(*string)
    GetCurrentUserOrganizationUrls()([]string)
    GetCurrentUserPublicUrl()(*string)
    GetCurrentUserUrl()(*string)
    GetLinks()(Feed__linksable)
    GetRepositoryDiscussionsCategoryUrl()(*string)
    GetRepositoryDiscussionsUrl()(*string)
    GetSecurityAdvisoriesUrl()(*string)
    GetTimelineUrl()(*string)
    GetUserUrl()(*string)
    SetCurrentUserActorUrl(value *string)()
    SetCurrentUserOrganizationUrl(value *string)()
    SetCurrentUserOrganizationUrls(value []string)()
    SetCurrentUserPublicUrl(value *string)()
    SetCurrentUserUrl(value *string)()
    SetLinks(value Feed__linksable)()
    SetRepositoryDiscussionsCategoryUrl(value *string)()
    SetRepositoryDiscussionsUrl(value *string)()
    SetSecurityAdvisoriesUrl(value *string)()
    SetTimelineUrl(value *string)()
    SetUserUrl(value *string)()
}
