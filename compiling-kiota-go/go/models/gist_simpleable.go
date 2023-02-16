package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistSimpleable 
type GistSimpleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComments()(*int32)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCreatedAt()(*string)
    GetDescription()(*string)
    GetFiles()(GistSimple_filesable)
    GetForkOf()(Gistable)
    GetForks()([]GistSimple_forksable)
    GetForksUrl()(*string)
    GetGitPullUrl()(*string)
    GetGitPushUrl()(*string)
    GetHistory()([]GistHistoryable)
    GetHtmlUrl()(*string)
    GetId()(*string)
    GetNodeId()(*string)
    GetOwner()(SimpleUserable)
    GetPublic()(*bool)
    GetTruncated()(*bool)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    GetUser()(*string)
    SetComments(value *int32)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCreatedAt(value *string)()
    SetDescription(value *string)()
    SetFiles(value GistSimple_filesable)()
    SetForkOf(value Gistable)()
    SetForks(value []GistSimple_forksable)()
    SetForksUrl(value *string)()
    SetGitPullUrl(value *string)()
    SetGitPushUrl(value *string)()
    SetHistory(value []GistHistoryable)()
    SetHtmlUrl(value *string)()
    SetId(value *string)()
    SetNodeId(value *string)()
    SetOwner(value SimpleUserable)()
    SetPublic(value *bool)()
    SetTruncated(value *bool)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
    SetUser(value *string)()
}
