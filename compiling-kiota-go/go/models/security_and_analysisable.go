package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityAndAnalysisable 
type SecurityAndAnalysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(SecurityAndAnalysis_advanced_securityable)
    GetSecretScanning()(SecurityAndAnalysis_secret_scanningable)
    GetSecretScanningPushProtection()(SecurityAndAnalysis_secret_scanning_push_protectionable)
    SetAdvancedSecurity(value SecurityAndAnalysis_advanced_securityable)()
    SetSecretScanning(value SecurityAndAnalysis_secret_scanningable)()
    SetSecretScanningPushProtection(value SecurityAndAnalysis_secret_scanning_push_protectionable)()
}
