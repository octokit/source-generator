package repos

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemWithRepoItemRequestBuilder builds and executes requests for operations under \repos\{org}\{repo}
type ItemWithRepoItemRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemWithRepoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemWithRepoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemWithRepoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithRepoItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Actions the actions property
func (m *ItemWithRepoItemRequestBuilder) Actions() *ItemItemActionsRequestBuilder {
	return NewItemItemActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Assignees the assignees property
func (m *ItemWithRepoItemRequestBuilder) Assignees() *ItemItemAssigneesRequestBuilder {
	return NewItemItemAssigneesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AssigneesById gets an item from the github.com/octokit/kiota/.repos.item.item.assignees.item collection
func (m *ItemWithRepoItemRequestBuilder) AssigneesById(id string) *ItemItemAssigneesWithAssigneeItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["assignee"] = id
	}
	return NewItemItemAssigneesWithAssigneeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Autolinks the autolinks property
func (m *ItemWithRepoItemRequestBuilder) Autolinks() *ItemItemAutolinksRequestBuilder {
	return NewItemItemAutolinksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AutolinksById gets an item from the github.com/octokit/kiota/.repos.item.item.autolinks.item collection
func (m *ItemWithRepoItemRequestBuilder) AutolinksById(id string) *ItemItemAutolinksWithAutolink_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["autolink_id"] = id
	}
	return NewItemItemAutolinksWithAutolink_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AutomatedSecurityFixes the automatedSecurityFixes property
func (m *ItemWithRepoItemRequestBuilder) AutomatedSecurityFixes() *ItemItemAutomatedSecurityFixesRequestBuilder {
	return NewItemItemAutomatedSecurityFixesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Branches the branches property
func (m *ItemWithRepoItemRequestBuilder) Branches() *ItemItemBranchesRequestBuilder {
	return NewItemItemBranchesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BranchesById gets an item from the github.com/octokit/kiota/.repos.item.item.branches.item collection
func (m *ItemWithRepoItemRequestBuilder) BranchesById(id string) *ItemItemBranchesWithBranchItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["branch"] = id
	}
	return NewItemItemBranchesWithBranchItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// CheckRuns the checkRuns property
func (m *ItemWithRepoItemRequestBuilder) CheckRuns() *ItemItemCheckRunsRequestBuilder {
	return NewItemItemCheckRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CheckRunsById gets an item from the github.com/octokit/kiota/.repos.item.item.checkRuns.item collection
func (m *ItemWithRepoItemRequestBuilder) CheckRunsById(id string) *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["check_run_id"] = id
	}
	return NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// CheckSuites the checkSuites property
func (m *ItemWithRepoItemRequestBuilder) CheckSuites() *ItemItemCheckSuitesRequestBuilder {
	return NewItemItemCheckSuitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CheckSuitesById gets an item from the github.com/octokit/kiota/.repos.item.item.checkSuites.item collection
func (m *ItemWithRepoItemRequestBuilder) CheckSuitesById(id string) *ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["check_suite_id"] = id
	}
	return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Codeowners the codeowners property
func (m *ItemWithRepoItemRequestBuilder) Codeowners() *ItemItemCodeownersRequestBuilder {
	return NewItemItemCodeownersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CodeScanning the codeScanning property
func (m *ItemWithRepoItemRequestBuilder) CodeScanning() *ItemItemCodeScanningRequestBuilder {
	return NewItemItemCodeScanningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Codespaces the codespaces property
func (m *ItemWithRepoItemRequestBuilder) Codespaces() *ItemItemCodespacesRequestBuilder {
	return NewItemItemCodespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Collaborators the collaborators property
func (m *ItemWithRepoItemRequestBuilder) Collaborators() *ItemItemCollaboratorsRequestBuilder {
	return NewItemItemCollaboratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CollaboratorsById gets an item from the github.com/octokit/kiota/.repos.item.item.collaborators.item collection
func (m *ItemWithRepoItemRequestBuilder) CollaboratorsById(id string) *ItemItemCollaboratorsWithUsernameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["username"] = id
	}
	return NewItemItemCollaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Comments the comments property
func (m *ItemWithRepoItemRequestBuilder) Comments() *ItemItemCommentsRequestBuilder {
	return NewItemItemCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CommentsById gets an item from the github.com/octokit/kiota/.repos.item.item.comments.item collection
func (m *ItemWithRepoItemRequestBuilder) CommentsById(id string) *ItemItemCommentsWithComment_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["comment_id"] = id
	}
	return NewItemItemCommentsWithComment_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Commits the commits property
func (m *ItemWithRepoItemRequestBuilder) Commits() *ItemItemCommitsRequestBuilder {
	return NewItemItemCommitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Community the community property
func (m *ItemWithRepoItemRequestBuilder) Community() *ItemItemCommunityRequestBuilder {
	return NewItemItemCommunityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Compare the compare property
func (m *ItemWithRepoItemRequestBuilder) Compare() *ItemItemCompareRequestBuilder {
	return NewItemItemCompareRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CompareById gets an item from the github.com/octokit/kiota/.repos.item.item.compare.item collection
func (m *ItemWithRepoItemRequestBuilder) CompareById(id string) *ItemItemCompareWithBaseheadItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["basehead"] = id
	}
	return NewItemItemCompareWithBaseheadItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// NewItemWithRepoItemRequestBuilderInternal instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemWithRepoItemRequestBuilder {
	m := &ItemWithRepoItemRequestBuilder{}
	m.urlTemplate = "{+baseurl}/repos/{org}/{repo}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemWithRepoItemRequestBuilder instantiates a new WithRepoItemRequestBuilder and sets the default values.
func NewItemWithRepoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemWithRepoItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemWithRepoItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Contents the contents property
func (m *ItemWithRepoItemRequestBuilder) Contents() *ItemItemContentsRequestBuilder {
	return NewItemItemContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ContentsById gets an item from the github.com/octokit/kiota/.repos.item.item.contents.item collection
func (m *ItemWithRepoItemRequestBuilder) ContentsById(id string) *ItemItemContentsWithPathItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["path"] = id
	}
	return NewItemItemContentsWithPathItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Contributors the contributors property
func (m *ItemWithRepoItemRequestBuilder) Contributors() *ItemItemContributorsRequestBuilder {
	return NewItemItemContributorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Delete deleting a repository requires admin access. If OAuth is used, the `delete_repo` scope is required.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/repos#delete-a-repository
func (m *ItemWithRepoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Dependabot the dependabot property
func (m *ItemWithRepoItemRequestBuilder) Dependabot() *ItemItemDependabotRequestBuilder {
	return NewItemItemDependabotRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DependencyGraph the dependencyGraph property
func (m *ItemWithRepoItemRequestBuilder) DependencyGraph() *ItemItemDependencyGraphRequestBuilder {
	return NewItemItemDependencyGraphRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Deployments the deployments property
func (m *ItemWithRepoItemRequestBuilder) Deployments() *ItemItemDeploymentsRequestBuilder {
	return NewItemItemDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DeploymentsById gets an item from the github.com/octokit/kiota/.repos.item.item.deployments.item collection
func (m *ItemWithRepoItemRequestBuilder) DeploymentsById(id string) *ItemItemDeploymentsWithDeployment_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["deployment_id"] = id
	}
	return NewItemItemDeploymentsWithDeployment_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Dispatches the dispatches property
func (m *ItemWithRepoItemRequestBuilder) Dispatches() *ItemItemDispatchesRequestBuilder {
	return NewItemItemDispatchesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Environments the environments property
func (m *ItemWithRepoItemRequestBuilder) Environments() *ItemItemEnvironmentsRequestBuilder {
	return NewItemItemEnvironmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// EnvironmentsById gets an item from the github.com/octokit/kiota/.repos.item.item.environments.item collection
func (m *ItemWithRepoItemRequestBuilder) EnvironmentsById(id string) *ItemItemEnvironmentsWithEnvironment_nameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["environment_name"] = id
	}
	return NewItemItemEnvironmentsWithEnvironment_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Events the events property
func (m *ItemWithRepoItemRequestBuilder) Events() *ItemItemEventsRequestBuilder {
	return NewItemItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Forks the forks property
func (m *ItemWithRepoItemRequestBuilder) Forks() *ItemItemForksRequestBuilder {
	return NewItemItemForksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Get the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/repos#get-a-repository
func (m *ItemWithRepoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderGetRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FullRepositoryable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateFullRepositoryFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FullRepositoryable), nil
}

// Git the git property
func (m *ItemWithRepoItemRequestBuilder) Git() *ItemItemGitRequestBuilder {
	return NewItemItemGitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Hooks the hooks property
func (m *ItemWithRepoItemRequestBuilder) Hooks() *ItemItemHooksRequestBuilder {
	return NewItemItemHooksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// HooksById gets an item from the github.com/octokit/kiota/.repos.item.item.hooks.item collection
func (m *ItemWithRepoItemRequestBuilder) HooksById(id string) *ItemItemHooksWithHook_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["hook_id"] = id
	}
	return NewItemItemHooksWithHook_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ImportEscaped the import property
func (m *ItemWithRepoItemRequestBuilder) ImportEscaped() *ItemItemImportRequestBuilder {
	return NewItemItemImportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Installation the installation property
func (m *ItemWithRepoItemRequestBuilder) Installation() *ItemItemInstallationRequestBuilder {
	return NewItemItemInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// InteractionLimits the interactionLimits property
func (m *ItemWithRepoItemRequestBuilder) InteractionLimits() *ItemItemInteractionLimitsRequestBuilder {
	return NewItemItemInteractionLimitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Invitations the invitations property
func (m *ItemWithRepoItemRequestBuilder) Invitations() *ItemItemInvitationsRequestBuilder {
	return NewItemItemInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// InvitationsById gets an item from the github.com/octokit/kiota/.repos.item.item.invitations.item collection
func (m *ItemWithRepoItemRequestBuilder) InvitationsById(id string) *ItemItemInvitationsWithInvitation_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["invitation_id"] = id
	}
	return NewItemItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Issues the issues property
func (m *ItemWithRepoItemRequestBuilder) Issues() *ItemItemIssuesRequestBuilder {
	return NewItemItemIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// IssuesById gets an item from the github.com/octokit/kiota/.repos.item.item.issues.item collection
func (m *ItemWithRepoItemRequestBuilder) IssuesById(id string) *ItemItemIssuesWithIssue_numberItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["issue_number"] = id
	}
	return NewItemItemIssuesWithIssue_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Keys the keys property
func (m *ItemWithRepoItemRequestBuilder) Keys() *ItemItemKeysRequestBuilder {
	return NewItemItemKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// KeysById gets an item from the github.com/octokit/kiota/.repos.item.item.keys.item collection
func (m *ItemWithRepoItemRequestBuilder) KeysById(id string) *ItemItemKeysWithKey_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["key_id"] = id
	}
	return NewItemItemKeysWithKey_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Labels the labels property
func (m *ItemWithRepoItemRequestBuilder) Labels() *ItemItemLabelsRequestBuilder {
	return NewItemItemLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// LabelsById gets an item from the github.com/octokit/kiota/.repos.item.item.labels.item collection
func (m *ItemWithRepoItemRequestBuilder) LabelsById(id string) *ItemItemLabelsWithNameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["name"] = id
	}
	return NewItemItemLabelsWithNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Languages the languages property
func (m *ItemWithRepoItemRequestBuilder) Languages() *ItemItemLanguagesRequestBuilder {
	return NewItemItemLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Lfs the lfs property
func (m *ItemWithRepoItemRequestBuilder) Lfs() *ItemItemLfsRequestBuilder {
	return NewItemItemLfsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// License the license property
func (m *ItemWithRepoItemRequestBuilder) License() *ItemItemLicenseRequestBuilder {
	return NewItemItemLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Merges the merges property
func (m *ItemWithRepoItemRequestBuilder) Merges() *ItemItemMergesRequestBuilder {
	return NewItemItemMergesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MergeUpstream the mergeUpstream property
func (m *ItemWithRepoItemRequestBuilder) MergeUpstream() *ItemItemMergeUpstreamRequestBuilder {
	return NewItemItemMergeUpstreamRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Milestones the milestones property
func (m *ItemWithRepoItemRequestBuilder) Milestones() *ItemItemMilestonesRequestBuilder {
	return NewItemItemMilestonesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MilestonesById gets an item from the github.com/octokit/kiota/.repos.item.item.milestones.item collection
func (m *ItemWithRepoItemRequestBuilder) MilestonesById(id string) *ItemItemMilestonesWithMilestone_numberItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["milestone_number"] = id
	}
	return NewItemItemMilestonesWithMilestone_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Notifications the notifications property
func (m *ItemWithRepoItemRequestBuilder) Notifications() *ItemItemNotificationsRequestBuilder {
	return NewItemItemNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Pages the pages property
func (m *ItemWithRepoItemRequestBuilder) Pages() *ItemItemPagesRequestBuilder {
	return NewItemItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Patch **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/reference/repos#replace-all-repository-topics) endpoint.
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/repos/repos#update-a-repository
func (m *ItemWithRepoItemRequestBuilder) Patch(ctx context.Context, body ItemItemWithRepoPatchRequestBodyable, requestConfiguration *ItemWithRepoItemRequestBuilderPatchRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FullRepositoryable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateFullRepositoryFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.FullRepositoryable), nil
}

// Projects the projects property
func (m *ItemWithRepoItemRequestBuilder) Projects() *ItemItemProjectsRequestBuilder {
	return NewItemItemProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Pulls the pulls property
func (m *ItemWithRepoItemRequestBuilder) Pulls() *ItemItemPullsRequestBuilder {
	return NewItemItemPullsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PullsById gets an item from the github.com/octokit/kiota/.repos.item.item.pulls.item collection
func (m *ItemWithRepoItemRequestBuilder) PullsById(id string) *ItemItemPullsWithPull_numberItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["pull_number"] = id
	}
	return NewItemItemPullsWithPull_numberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Readme the readme property
func (m *ItemWithRepoItemRequestBuilder) Readme() *ItemItemReadmeRequestBuilder {
	return NewItemItemReadmeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ReadmeById gets an item from the github.com/octokit/kiota/.repos.item.item.readme.item collection
func (m *ItemWithRepoItemRequestBuilder) ReadmeById(id string) *ItemItemReadmeWithDirItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["dir"] = id
	}
	return NewItemItemReadmeWithDirItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Releases the releases property
func (m *ItemWithRepoItemRequestBuilder) Releases() *ItemItemReleasesRequestBuilder {
	return NewItemItemReleasesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ReleasesById gets an item from the github.com/octokit/kiota/.repos.item.item.releases.item collection
func (m *ItemWithRepoItemRequestBuilder) ReleasesById(id string) *ItemItemReleasesWithRelease_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["release_id"] = id
	}
	return NewItemItemReleasesWithRelease_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// SecretScanning the secretScanning property
func (m *ItemWithRepoItemRequestBuilder) SecretScanning() *ItemItemSecretScanningRequestBuilder {
	return NewItemItemSecretScanningRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Stargazers the stargazers property
func (m *ItemWithRepoItemRequestBuilder) Stargazers() *ItemItemStargazersRequestBuilder {
	return NewItemItemStargazersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Stats the stats property
func (m *ItemWithRepoItemRequestBuilder) Stats() *ItemItemStatsRequestBuilder {
	return NewItemItemStatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Statuses the statuses property
func (m *ItemWithRepoItemRequestBuilder) Statuses() *ItemItemStatusesRequestBuilder {
	return NewItemItemStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// StatusesById gets an item from the github.com/octokit/kiota/.repos.item.item.statuses.item collection
func (m *ItemWithRepoItemRequestBuilder) StatusesById(id string) *ItemItemStatusesWithShaItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["sha"] = id
	}
	return NewItemItemStatusesWithShaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Subscribers the subscribers property
func (m *ItemWithRepoItemRequestBuilder) Subscribers() *ItemItemSubscribersRequestBuilder {
	return NewItemItemSubscribersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Subscription the subscription property
func (m *ItemWithRepoItemRequestBuilder) Subscription() *ItemItemSubscriptionRequestBuilder {
	return NewItemItemSubscriptionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Tags the tags property
func (m *ItemWithRepoItemRequestBuilder) Tags() *ItemItemTagsRequestBuilder {
	return NewItemItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Tarball the tarball property
func (m *ItemWithRepoItemRequestBuilder) Tarball() *ItemItemTarballRequestBuilder {
	return NewItemItemTarballRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TarballById gets an item from the github.com/octokit/kiota/.repos.item.item.tarball.item collection
func (m *ItemWithRepoItemRequestBuilder) TarballById(id string) *ItemItemTarballWithRefItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["ref"] = id
	}
	return NewItemItemTarballWithRefItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Teams the teams property
func (m *ItemWithRepoItemRequestBuilder) Teams() *ItemItemTeamsRequestBuilder {
	return NewItemItemTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ToDeleteRequestInformation deleting a repository requires admin access. If OAuth is used, the `delete_repo` scope is required.If an organization owner has configured the organization to prevent members from deleting organization-ownedrepositories, you will get a `403 Forbidden` response.
func (m *ItemWithRepoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation the `parent` and `source` objects are present when the repository is a fork. `parent` is the repository this repository was forked from, `source` is the ultimate source for the network.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
func (m *ItemWithRepoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWithRepoItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToPatchRequestInformation **Note**: To edit a repository's topics, use the [Replace all repository topics](https://docs.github.com/rest/reference/repos#replace-all-repository-topics) endpoint.
func (m *ItemWithRepoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemWithRepoPatchRequestBodyable, requestConfiguration *ItemWithRepoItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
	requestInfo.Headers.Add("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// Topics the topics property
func (m *ItemWithRepoItemRequestBuilder) Topics() *ItemItemTopicsRequestBuilder {
	return NewItemItemTopicsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Traffic the traffic property
func (m *ItemWithRepoItemRequestBuilder) Traffic() *ItemItemTrafficRequestBuilder {
	return NewItemItemTrafficRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Transfer the transfer property
func (m *ItemWithRepoItemRequestBuilder) Transfer() *ItemItemTransferRequestBuilder {
	return NewItemItemTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// VulnerabilityAlerts the vulnerabilityAlerts property
func (m *ItemWithRepoItemRequestBuilder) VulnerabilityAlerts() *ItemItemVulnerabilityAlertsRequestBuilder {
	return NewItemItemVulnerabilityAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Zipball the zipball property
func (m *ItemWithRepoItemRequestBuilder) Zipball() *ItemItemZipballRequestBuilder {
	return NewItemItemZipballRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ZipballById gets an item from the github.com/octokit/kiota/.repos.item.item.zipball.item collection
func (m *ItemWithRepoItemRequestBuilder) ZipballById(id string) *ItemItemZipballWithRefItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["ref"] = id
	}
	return NewItemItemZipballWithRefItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
