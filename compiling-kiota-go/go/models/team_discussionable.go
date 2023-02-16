package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamDiscussionable 
type TeamDiscussionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetBody()(*string)
    GetBodyHtml()(*string)
    GetBodyVersion()(*string)
    GetCommentsCount()(*int32)
    GetCommentsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetLastEditedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNodeId()(*string)
    GetNumber()(*int32)
    GetPinned()(*bool)
    GetPrivate()(*bool)
    GetReactions()(ReactionRollupable)
    GetTeamUrl()(*string)
    GetTitle()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAuthor(value NullableSimpleUserable)()
    SetBody(value *string)()
    SetBodyHtml(value *string)()
    SetBodyVersion(value *string)()
    SetCommentsCount(value *int32)()
    SetCommentsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetLastEditedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNodeId(value *string)()
    SetNumber(value *int32)()
    SetPinned(value *bool)()
    SetPrivate(value *bool)()
    SetReactions(value ReactionRollupable)()
    SetTeamUrl(value *string)()
    SetTitle(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
