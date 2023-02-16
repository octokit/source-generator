package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequest_head_repo_ownerable 
type PullRequest_head_repo_ownerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvatarUrl()(*string)
    GetEventsUrl()(*string)
    GetFollowersUrl()(*string)
    GetFollowingUrl()(*string)
    GetGistsUrl()(*string)
    GetGravatarId()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetNodeId()(*string)
    GetOrganizationsUrl()(*string)
    GetReceivedEventsUrl()(*string)
    GetReposUrl()(*string)
    GetSiteAdmin()(*bool)
    GetStarredUrl()(*string)
    GetSubscriptionsUrl()(*string)
    GetType()(*string)
    GetUrl()(*string)
    SetAvatarUrl(value *string)()
    SetEventsUrl(value *string)()
    SetFollowersUrl(value *string)()
    SetFollowingUrl(value *string)()
    SetGistsUrl(value *string)()
    SetGravatarId(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetNodeId(value *string)()
    SetOrganizationsUrl(value *string)()
    SetReceivedEventsUrl(value *string)()
    SetReposUrl(value *string)()
    SetSiteAdmin(value *bool)()
    SetStarredUrl(value *string)()
    SetSubscriptionsUrl(value *string)()
    SetType(value *string)()
    SetUrl(value *string)()
}
