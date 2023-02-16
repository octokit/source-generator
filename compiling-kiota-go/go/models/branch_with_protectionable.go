package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BranchWithProtectionable 
type BranchWithProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommit()(Commitable)
    GetLinks()(BranchWithProtection__linksable)
    GetName()(*string)
    GetPattern()(*string)
    GetProtected()(*bool)
    GetProtection()(BranchProtectionable)
    GetProtectionUrl()(*string)
    GetRequiredApprovingReviewCount()(*int32)
    SetCommit(value Commitable)()
    SetLinks(value BranchWithProtection__linksable)()
    SetName(value *string)()
    SetPattern(value *string)()
    SetProtected(value *bool)()
    SetProtection(value BranchProtectionable)()
    SetProtectionUrl(value *string)()
    SetRequiredApprovingReviewCount(value *int32)()
}
