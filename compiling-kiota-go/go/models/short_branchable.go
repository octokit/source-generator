package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ShortBranchable 
type ShortBranchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(ShortBranch_commitable)
    GetName()(*string)
    GetProtected()(*bool)
    GetProtection()(BranchProtectionable)
    GetProtectionUrl()(*string)
    SetCommit(value ShortBranch_commitable)()
    SetName(value *string)()
    SetProtected(value *bool)()
    SetProtection(value BranchProtectionable)()
    SetProtectionUrl(value *string)()
}
