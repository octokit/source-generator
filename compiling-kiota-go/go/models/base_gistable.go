package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseGistable 
type BaseGistable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetFiles()(BaseGist_filesable)
    GetForksUrl()(*string)
    GetGitPullUrl()(*string)
    GetGitPushUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetNodeId()(*string)
    GetOwner()(SimpleUserable)
    GetPublic()(*bool)
    GetTruncated()(*bool)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetUser()(NullableSimpleUserable)
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetFiles(value BaseGist_filesable)()
    SetForksUrl(value *string)()
    SetGitPullUrl(value *string)()
    SetGitPushUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetNodeId(value *string)()
    SetOwner(value SimpleUserable)()
    SetPublic(value *bool)()
    SetTruncated(value *bool)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetUser(value NullableSimpleUserable)()
}
