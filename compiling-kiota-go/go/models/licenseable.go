package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Licenseable 
type Licenseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetConditions()([]string)
    GetDescription()(*string)
    GetFeatured()(*bool)
    GetHtmlUrl()(*string)
    GetImplementation()(*string)
    GetKey()(*string)
    GetLimitations()([]string)
    GetName()(*string)
    GetNodeId()(*string)
    GetPermissions()([]string)
    GetSpdxId()(*string)
    GetUrl()(*string)
    SetBody(value *string)()
    SetConditions(value []string)()
    SetDescription(value *string)()
    SetFeatured(value *bool)()
    SetHtmlUrl(value *string)()
    SetImplementation(value *string)()
    SetKey(value *string)()
    SetLimitations(value []string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetPermissions(value []string)()
    SetSpdxId(value *string)()
    SetUrl(value *string)()
}
