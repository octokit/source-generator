package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPullsItemCommentsPostRequestBodyable 
type ItemItemPullsItemCommentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetCommitId()(*string)
    GetInReplyTo()(*int32)
    GetLine()(*int32)
    GetPath()(*string)
    GetPosition()(*int32)
    GetStartLine()(*int32)
    SetBody(value *string)()
    SetCommitId(value *string)()
    SetInReplyTo(value *int32)()
    SetLine(value *int32)()
    SetPath(value *string)()
    SetPosition(value *int32)()
    SetStartLine(value *int32)()
}
