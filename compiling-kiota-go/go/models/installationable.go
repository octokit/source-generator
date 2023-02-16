package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Installationable 
type Installationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessTokensUrl()(*string)
    GetAccount()(Installation_Installationsable)
    GetAppId()(*int32)
    GetAppSlug()(*string)
    GetContactEmail()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEvents()([]string)
    GetHasMultipleSingleFiles()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetPermissions()(AppPermissionsable)
    GetRepositoriesUrl()(*string)
    GetRepositorySelection()(*Installation_repository_selection)
    GetSingleFileName()(*string)
    GetSingleFilePaths()([]string)
    GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuspendedBy()(NullableSimpleUserable)
    GetTargetId()(*int32)
    GetTargetType()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessTokensUrl(value *string)()
    SetAccount(value Installation_Installationsable)()
    SetAppId(value *int32)()
    SetAppSlug(value *string)()
    SetContactEmail(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEvents(value []string)()
    SetHasMultipleSingleFiles(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetPermissions(value AppPermissionsable)()
    SetRepositoriesUrl(value *string)()
    SetRepositorySelection(value *Installation_repository_selection)()
    SetSingleFileName(value *string)()
    SetSingleFilePaths(value []string)()
    SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuspendedBy(value NullableSimpleUserable)()
    SetTargetId(value *int32)()
    SetTargetType(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
