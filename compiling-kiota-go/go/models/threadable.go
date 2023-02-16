package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Threadable 
type Threadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetLastReadAt()(*string)
    GetReason()(*string)
    GetRepository()(MinimalRepositoryable)
    GetSubject()(Thread_subjectable)
    GetSubscriptionUrl()(*string)
    GetUnread()(*bool)
    GetUpdatedAt()(*string)
    GetUrl()(*string)
    SetId(value *string)()
    SetLastReadAt(value *string)()
    SetReason(value *string)()
    SetRepository(value MinimalRepositoryable)()
    SetSubject(value Thread_subjectable)()
    SetSubscriptionUrl(value *string)()
    SetUnread(value *bool)()
    SetUpdatedAt(value *string)()
    SetUrl(value *string)()
}
