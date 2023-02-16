package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// WithUsernameItemRequestBuilder builds and executes requests for operations under \users\{username}
type WithUsernameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WithUsernameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithUsernameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithUsernameResponse composed type wrapper for classes privateUser, publicUser
type WithUsernameResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type privateUser
    privateUser ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable
    // Composed type representation for type publicUser
    publicUser ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewWithUsernameResponse instantiates a new WithUsernameResponse and sets the default values.
func NewWithUsernameResponse()(*WithUsernameResponse) {
    m := &WithUsernameResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWithUsernameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithUsernameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithUsernameResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WithUsernameResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WithUsernameResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *WithUsernameResponse) GetPrivateUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameResponse) GetPublicUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *WithUsernameResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPrivateUser() != nil {
        err := writer.WriteObjectValue("", m.GetPrivateUser())
        if err != nil {
            return err
        }
    } else if m.GetPublicUser() != nil {
        err := writer.WriteObjectValue("", m.GetPublicUser())
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
func (m *WithUsernameResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPrivateUser sets the privateUser property value. Composed type representation for type privateUser
func (m *WithUsernameResponse) SetPrivateUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameResponse) SetPublicUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)() {
    m.publicUser = value
}
// WithUsernameResponseable 
type WithUsernameResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)
    GetPublicUser()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)
    SetPrivateUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PrivateUserable)()
    SetPublicUser(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.PublicUserable)()
}
// NewWithUsernameItemRequestBuilderInternal instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUsernameItemRequestBuilder) {
    m := &WithUsernameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWithUsernameItemRequestBuilder instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Events the events property
func (m *WithUsernameItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Followers the followers property
func (m *WithUsernameItemRequestBuilder) Followers()(*ItemFollowersRequestBuilder) {
    return NewItemFollowersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Following the following property
func (m *WithUsernameItemRequestBuilder) Following()(*ItemFollowingRequestBuilder) {
    return NewItemFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FollowingById gets an item from the github.com/octokit/kiota/.users.item.following.item collection
func (m *WithUsernameItemRequestBuilder) FollowingById(id string)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["target_user"] = id
    }
    return NewItemFollowingWithTarget_userItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get provides publicly available information about someone with a GitHub account.GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://docs.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see 'Response with GitHub plan information' below"The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/overview/resources-in-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/reference/users#emails)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#get-a-user
func (m *WithUsernameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithUsernameItemRequestBuilderGetRequestConfiguration)(WithUsernameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateWithUsernameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithUsernameResponseable), nil
}
// Gists the gists property
func (m *WithUsernameItemRequestBuilder) Gists()(*ItemGistsRequestBuilder) {
    return NewItemGistsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Gpg_keys the gpg_keys property
func (m *WithUsernameItemRequestBuilder) Gpg_keys()(*ItemGpg_keysRequestBuilder) {
    return NewItemGpg_keysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Hovercard the hovercard property
func (m *WithUsernameItemRequestBuilder) Hovercard()(*ItemHovercardRequestBuilder) {
    return NewItemHovercardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Installation the installation property
func (m *WithUsernameItemRequestBuilder) Installation()(*ItemInstallationRequestBuilder) {
    return NewItemInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Keys the keys property
func (m *WithUsernameItemRequestBuilder) Keys()(*ItemKeysRequestBuilder) {
    return NewItemKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Orgs the orgs property
func (m *WithUsernameItemRequestBuilder) Orgs()(*ItemOrgsRequestBuilder) {
    return NewItemOrgsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Packages the packages property
func (m *WithUsernameItemRequestBuilder) Packages()(*ItemPackagesRequestBuilder) {
    return NewItemPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PackagesById gets an item from the github.com/octokit/kiota/.users.item.packages.item collection
func (m *WithUsernameItemRequestBuilder) PackagesById(id string)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["package_type"] = id
    }
    return NewItemPackagesWithPackage_typeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Projects the projects property
func (m *WithUsernameItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Received_events the received_events property
func (m *WithUsernameItemRequestBuilder) Received_events()(*ItemReceived_eventsRequestBuilder) {
    return NewItemReceived_eventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Repos the repos property
func (m *WithUsernameItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Settings the settings property
func (m *WithUsernameItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ssh_signing_keys the ssh_signing_keys property
func (m *WithUsernameItemRequestBuilder) Ssh_signing_keys()(*ItemSsh_signing_keysRequestBuilder) {
    return NewItemSsh_signing_keysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Starred the starred property
func (m *WithUsernameItemRequestBuilder) Starred()(*ItemStarredRequestBuilder) {
    return NewItemStarredRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Subscriptions the subscriptions property
func (m *WithUsernameItemRequestBuilder) Subscriptions()(*ItemSubscriptionsRequestBuilder) {
    return NewItemSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation provides publicly available information about someone with a GitHub account.GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://docs.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see 'Response with GitHub plan information' below"The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/overview/resources-in-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/reference/users#emails)".
func (m *WithUsernameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithUsernameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
