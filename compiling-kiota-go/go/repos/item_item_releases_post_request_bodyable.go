package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemReleasesPostRequestBodyable 
type ItemItemReleasesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetDiscussionCategoryName()(*string)
    GetDraft()(*bool)
    GetGenerateReleaseNotes()(*bool)
    GetName()(*string)
    GetPrerelease()(*bool)
    GetTagName()(*string)
    GetTargetCommitish()(*string)
    SetBody(value *string)()
    SetDiscussionCategoryName(value *string)()
    SetDraft(value *bool)()
    SetGenerateReleaseNotes(value *bool)()
    SetName(value *string)()
    SetPrerelease(value *bool)()
    SetTagName(value *string)()
    SetTargetCommitish(value *string)()
}
