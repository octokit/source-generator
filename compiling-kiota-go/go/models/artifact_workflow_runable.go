package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Artifact_workflow_runable 
type Artifact_workflow_runable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeadBranch()(*string)
    GetHeadRepositoryId()(*int32)
    GetHeadSha()(*string)
    GetId()(*int32)
    GetRepositoryId()(*int32)
    SetHeadBranch(value *string)()
    SetHeadRepositoryId(value *int32)()
    SetHeadSha(value *string)()
    SetId(value *int32)()
    SetRepositoryId(value *int32)()
}
