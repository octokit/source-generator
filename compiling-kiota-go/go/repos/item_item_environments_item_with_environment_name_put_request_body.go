package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody 
type ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody struct {
    // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
    deployment_branch_policy ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicySettingsable
    // The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
    reviewers []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable
    // The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
    wait_timer *int32
}
// NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody instantiates a new ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody and sets the default values.
func NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody()(*ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) {
    m := &ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody{
    }
    return m
}
// CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemWithEnvironment_namePutRequestBody(), nil
}
// GetDeploymentBranchPolicy gets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetDeploymentBranchPolicy()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicySettingsable) {
    return m.deployment_branch_policy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deployment_branch_policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateDeploymentBranchPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentBranchPolicy(val.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicySettingsable))
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable, len(val))
            for i, v := range val {
                res[i] = v.(ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["wait_timer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitTimer(val)
        }
        return nil
    }
    return res
}
// GetReviewers gets the reviewers property value. The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetReviewers()([]ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable) {
    return m.reviewers
}
// GetWaitTimer gets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) GetWaitTimer()(*int32) {
    return m.wait_timer
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deployment_branch_policy", m.GetDeploymentBranchPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wait_timer", m.GetWaitTimer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentBranchPolicy sets the deployment_branch_policy property value. The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetDeploymentBranchPolicy(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.DeploymentBranchPolicySettingsable)() {
    m.deployment_branch_policy = value
}
// SetReviewers sets the reviewers property value. The people or teams that may review jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetReviewers(value []ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody_reviewersable)() {
    m.reviewers = value
}
// SetWaitTimer sets the wait_timer property value. The amount of time to delay a job after the job is initially triggered. The time (in minutes) must be an integer between 0 and 43,200 (30 days).
func (m *ItemItemEnvironmentsItemWithEnvironment_namePutRequestBody) SetWaitTimer(value *int32)() {
    m.wait_timer = value
}
