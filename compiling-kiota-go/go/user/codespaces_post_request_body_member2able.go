package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesPostRequestBodyMember2able 
type CodespacesPostRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevcontainerPath()(*string)
    GetIdleTimeoutMinutes()(*int32)
    GetLocation()(*string)
    GetMachine()(*string)
    GetPullRequest()(CodespacesPostRequestBodyMember2_pull_requestable)
    GetWorkingDirectory()(*string)
    SetDevcontainerPath(value *string)()
    SetIdleTimeoutMinutes(value *int32)()
    SetLocation(value *string)()
    SetMachine(value *string)()
    SetPullRequest(value CodespacesPostRequestBodyMember2_pull_requestable)()
    SetWorkingDirectory(value *string)()
}
