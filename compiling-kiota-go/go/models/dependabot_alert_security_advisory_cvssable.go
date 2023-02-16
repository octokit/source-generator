package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory_cvssable 
type DependabotAlertSecurityAdvisory_cvssable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScore()(*int64)
    GetVectorString()(*string)
    SetScore(value *int64)()
    SetVectorString(value *string)()
}
