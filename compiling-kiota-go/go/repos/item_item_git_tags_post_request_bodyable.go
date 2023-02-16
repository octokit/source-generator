package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemGitTagsPostRequestBodyable 
type ItemItemGitTagsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    GetObject()(*string)
    GetTag()(*string)
    GetTagger()(ItemItemGitTagsPostRequestBody_taggerable)
    SetMessage(value *string)()
    SetObject(value *string)()
    SetTag(value *string)()
    SetTagger(value ItemItemGitTagsPostRequestBody_taggerable)()
}
