package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunityProfile_filesable 
type CommunityProfile_filesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodeOfConduct()(NullableCodeOfConductSimpleable)
    GetCodeOfConductFile()(NullableCommunityHealthFileable)
    GetContributing()(NullableCommunityHealthFileable)
    GetIssueTemplate()(NullableCommunityHealthFileable)
    GetLicense()(NullableLicenseSimpleable)
    GetPullRequestTemplate()(NullableCommunityHealthFileable)
    GetReadme()(NullableCommunityHealthFileable)
    SetCodeOfConduct(value NullableCodeOfConductSimpleable)()
    SetCodeOfConductFile(value NullableCommunityHealthFileable)()
    SetContributing(value NullableCommunityHealthFileable)()
    SetIssueTemplate(value NullableCommunityHealthFileable)()
    SetLicense(value NullableLicenseSimpleable)()
    SetPullRequestTemplate(value NullableCommunityHealthFileable)()
    SetReadme(value NullableCommunityHealthFileable)()
}
