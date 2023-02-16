package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemContentsItemWithPathPutRequestBodyable 
type ItemItemContentsItemWithPathPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(ItemItemContentsItemWithPathPutRequestBody_authorable)
    GetBranch()(*string)
    GetCommitter()(ItemItemContentsItemWithPathPutRequestBody_committerable)
    GetContent()(*string)
    GetMessage()(*string)
    GetSha()(*string)
    SetAuthor(value ItemItemContentsItemWithPathPutRequestBody_authorable)()
    SetBranch(value *string)()
    SetCommitter(value ItemItemContentsItemWithPathPutRequestBody_committerable)()
    SetContent(value *string)()
    SetMessage(value *string)()
    SetSha(value *string)()
}
