package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPatchRequestBodyable 
type UserPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBio()(*string)
    GetBlog()(*string)
    GetCompany()(*string)
    GetEmail()(*string)
    GetHireable()(*bool)
    GetLocation()(*string)
    GetName()(*string)
    GetTwitterUsername()(*string)
    SetBio(value *string)()
    SetBlog(value *string)()
    SetCompany(value *string)()
    SetEmail(value *string)()
    SetHireable(value *bool)()
    SetLocation(value *string)()
    SetName(value *string)()
    SetTwitterUsername(value *string)()
}
