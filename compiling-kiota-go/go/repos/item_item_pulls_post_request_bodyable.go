package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPullsPostRequestBodyable 
type ItemItemPullsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBase()(*string)
    GetBody()(*string)
    GetDraft()(*bool)
    GetHead()(*string)
    GetIssue()(*int32)
    GetMaintainerCanModify()(*bool)
    GetTitle()(*string)
    SetBase(value *string)()
    SetBody(value *string)()
    SetDraft(value *bool)()
    SetHead(value *string)()
    SetIssue(value *int32)()
    SetMaintainerCanModify(value *bool)()
    SetTitle(value *string)()
}
