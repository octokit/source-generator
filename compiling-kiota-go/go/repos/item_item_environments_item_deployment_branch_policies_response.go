package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse 
type ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The branch_policies property
    branch_policies []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable
    // The number of deployment branch policies for the environment.
    total_count *int32
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse instantiates a new ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeploymentBranchPoliciesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemDeploymentBranchPoliciesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranchPolicies gets the branch_policies property value. The branch_policies property
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) GetBranchPolicies()([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable) {
    return m.branch_policies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch_policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDeploymentBranchPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable)
            }
            m.SetBranchPolicies(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The number of deployment branch policies for the environment.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBranchPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBranchPolicies()))
        for i, v := range m.GetBranchPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("branch_policies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranchPolicies sets the branch_policies property value. The branch_policies property
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) SetBranchPolicies(value []ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicyable)() {
    m.branch_policies = value
}
// SetTotalCount sets the total_count property value. The number of deployment branch policies for the environment.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
