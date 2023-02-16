package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicSearchResultItemable 
type TopicSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAliases()([]TopicSearchResultItem_aliasesable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(*string)
    GetCurated()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFeatured()(*bool)
    GetLogoUrl()(*string)
    GetName()(*string)
    GetRelated()([]TopicSearchResultItem_relatedable)
    GetReleased()(*string)
    GetRepositoryCount()(*int32)
    GetScore()(*int64)
    GetShortDescription()(*string)
    GetTextMatches()([]TopicSearchResultItem_text_matchesable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAliases(value []TopicSearchResultItem_aliasesable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value *string)()
    SetCurated(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFeatured(value *bool)()
    SetLogoUrl(value *string)()
    SetName(value *string)()
    SetRelated(value []TopicSearchResultItem_relatedable)()
    SetReleased(value *string)()
    SetRepositoryCount(value *int32)()
    SetScore(value *int64)()
    SetShortDescription(value *string)()
    SetTextMatches(value []TopicSearchResultItem_text_matchesable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
