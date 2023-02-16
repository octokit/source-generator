package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistSimple403Errorable 
type GistSimple403Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlock()(GistSimple403Error_blockable)
    GetDocumentationUrl()(*string)
    GetMessage()(*string)
    SetBlock(value GistSimple403Error_blockable)()
    SetDocumentationUrl(value *string)()
    SetMessage(value *string)()
}
