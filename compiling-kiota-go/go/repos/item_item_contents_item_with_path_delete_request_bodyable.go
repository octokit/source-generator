package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemContentsItemWithPathDeleteRequestBodyable 
type ItemItemContentsItemWithPathDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(ItemItemContentsItemWithPathDeleteRequestBody_authorable)
    GetBranch()(*string)
    GetCommitter()(ItemItemContentsItemWithPathDeleteRequestBody_committerable)
    GetMessage()(*string)
    GetSha()(*string)
    SetAuthor(value ItemItemContentsItemWithPathDeleteRequestBody_authorable)()
    SetBranch(value *string)()
    SetCommitter(value ItemItemContentsItemWithPathDeleteRequestBody_committerable)()
    SetMessage(value *string)()
    SetSha(value *string)()
}
