package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentSubmoduleable 
type ContentSubmoduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadUrl()(*string)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLinks()(ContentSubmodule__linksable)
    GetName()(*string)
    GetPath()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetSubmoduleGitUrl()(*string)
    GetType()(*ContentSubmodule_type)
    GetUrl()(*string)
    SetDownloadUrl(value *string)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLinks(value ContentSubmodule__linksable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetSubmoduleGitUrl(value *string)()
    SetType(value *ContentSubmodule_type)()
    SetUrl(value *string)()
}
