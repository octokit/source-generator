package gists

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GistsPostRequestBodyable 
type GistsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFiles()(GistsPostRequestBody_filesable)
    GetPublic()(GistsPostRequestBody_Gistsable)
    SetDescription(value *string)()
    SetFiles(value GistsPostRequestBody_filesable)()
    SetPublic(value GistsPostRequestBody_Gistsable)()
}
