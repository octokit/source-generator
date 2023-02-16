package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseContentable 
type LicenseContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetDownloadUrl()(*string)
    GetEncoding()(*string)
    GetGitUrl()(*string)
    GetHtmlUrl()(*string)
    GetLicense()(NullableLicenseSimpleable)
    GetLinks()(LicenseContent__linksable)
    GetName()(*string)
    GetPath()(*string)
    GetSha()(*string)
    GetSize()(*int32)
    GetType()(*string)
    GetUrl()(*string)
    SetContent(value *string)()
    SetDownloadUrl(value *string)()
    SetEncoding(value *string)()
    SetGitUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetLicense(value NullableLicenseSimpleable)()
    SetLinks(value LicenseContent__linksable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetSha(value *string)()
    SetSize(value *int32)()
    SetType(value *string)()
    SetUrl(value *string)()
}
