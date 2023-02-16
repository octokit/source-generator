package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueSearchResultItem_text_matchesable 
type IssueSearchResultItem_text_matchesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFragment()(*string)
    GetMatches()([]IssueSearchResultItem_text_matches_matchesable)
    GetObjectType()(*string)
    GetObjectUrl()(*string)
    GetProperty()(*string)
    SetFragment(value *string)()
    SetMatches(value []IssueSearchResultItem_text_matches_matchesable)()
    SetObjectType(value *string)()
    SetObjectUrl(value *string)()
    SetProperty(value *string)()
}
