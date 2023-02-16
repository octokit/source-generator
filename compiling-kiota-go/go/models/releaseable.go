package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Releaseable 
type Releaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssets()([]ReleaseAssetable)
    GetAssetsUrl()(*string)
    GetAuthor()(SimpleUserable)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyText()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDiscussionUrl()(*string)
    GetDraft()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetMentionsCount()(*int32)
    GetName()(*string)
    GetNodeId()(*string)
    GetPrerelease()(*bool)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReactions()(ReactionRollupable)
    GetTagName()(*string)
    GetTarballUrl()(*string)
    GetTargetCommitish()(*string)
    GetUploadUrl()(*string)
    GetUrl()(*string)
    GetZipballUrl()(*string)
    SetAssets(value []ReleaseAssetable)()
    SetAssetsUrl(value *string)()
    SetAuthor(value SimpleUserable)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyText(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDiscussionUrl(value *string)()
    SetDraft(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetMentionsCount(value *int32)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPrerelease(value *bool)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReactions(value ReactionRollupable)()
    SetTagName(value *string)()
    SetTarballUrl(value *string)()
    SetTargetCommitish(value *string)()
    SetUploadUrl(value *string)()
    SetUrl(value *string)()
    SetZipballUrl(value *string)()
}
