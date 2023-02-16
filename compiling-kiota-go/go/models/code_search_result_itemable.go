package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSearchResultItemable 
type CodeSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileSize()(*int32)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLanguage()(*string)
    GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLineNumbers()([]string)
    GetName()(*string)
    GetPath()(*string)
    GetRepository()(MinimalRepositoryable)
    GetScore()(*int64)
    GetSha()(*string)
    GetTextMatches()([]CodeSearchResultItem_text_matchesable)
    GetUrl()(*string)
    SetFileSize(value *int32)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLanguage(value *string)()
    SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLineNumbers(value []string)()
    SetName(value *string)()
    SetPath(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetScore(value *int64)()
    SetSha(value *string)()
    SetTextMatches(value []CodeSearchResultItem_text_matchesable)()
    SetUrl(value *string)()
}
