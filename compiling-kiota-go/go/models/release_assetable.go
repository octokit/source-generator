package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReleaseAssetable 
type ReleaseAssetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrowserDownloadUrl()(*string)
    GetContentType()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDownloadCount()(*int32)
    GetId()(*int32)
    GetLabel()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetSize()(*int32)
    GetState()(*ReleaseAsset_state)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUploader()(NullableSimpleUserable)
    GetUrl()(*string)
    SetBrowserDownloadUrl(value *string)()
    SetContentType(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDownloadCount(value *int32)()
    SetId(value *int32)()
    SetLabel(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetSize(value *int32)()
    SetState(value *ReleaseAsset_state)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUploader(value NullableSimpleUserable)()
    SetUrl(value *string)()
}
