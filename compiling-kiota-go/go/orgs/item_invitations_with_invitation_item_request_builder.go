package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemInvitationsWithInvitation_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\invitations\{invitation_id}
type ItemInvitationsWithInvitation_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInvitationsWithInvitation_ItemRequestBuilderInternal instantiates a new WithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    m := &ItemInvitationsWithInvitation_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/invitations/{invitation_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemInvitationsWithInvitation_ItemRequestBuilder instantiates a new WithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#cancel-an-organization-invitation
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Teams the teams property
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Teams()(*ItemInvitationsItemTeamsRequestBuilder) {
    return NewItemInvitationsItemTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications).
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
