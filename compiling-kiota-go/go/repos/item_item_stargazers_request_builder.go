package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemStargazersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stargazers
type ItemItemStargazersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemStargazersRequestBuilderGetQueryParameters lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
type ItemItemStargazersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemStargazersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStargazersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemStargazersRequestBuilderGetQueryParameters
}
// StargazersResponse composed type wrapper for classes simpleUser, stargazer
type StargazersResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type simpleUser
    simpleUser ItemItemStargazersSimpleUserable
    // Composed type representation for type stargazer
    stargazer ItemItemStargazersStargazerable
}
// NewStargazersResponse instantiates a new stargazersResponse and sets the default values.
func NewStargazersResponse()(*StargazersResponse) {
    m := &StargazersResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStargazersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStargazersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStargazersResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemItemStargazersSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersSimpleUserable); ok {
                result.SetSimpleUser(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemStargazersStargazerFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemStargazersStargazerable); ok {
                result.SetStargazer(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StargazersResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StargazersResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type simpleUser
func (m *StargazersResponse) GetSimpleUser()(ItemItemStargazersSimpleUserable) {
    return m.simpleUser
}
// GetStargazer gets the stargazer property value. Composed type representation for type stargazer
func (m *StargazersResponse) GetStargazer()(ItemItemStargazersStargazerable) {
    return m.stargazer
}
// Serialize serializes information the current object
func (m *StargazersResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetStargazer() != nil {
        err := writer.WriteObjectValue("", m.GetStargazer())
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
func (m *StargazersResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type simpleUser
func (m *StargazersResponse) SetSimpleUser(value ItemItemStargazersSimpleUserable)() {
    m.simpleUser = value
}
// SetStargazer sets the stargazer property value. Composed type representation for type stargazer
func (m *StargazersResponse) SetStargazer(value ItemItemStargazersStargazerable)() {
    m.stargazer = value
}
// StargazersResponseable 
type StargazersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSimpleUser()(ItemItemStargazersSimpleUserable)
    GetStargazer()(ItemItemStargazersStargazerable)
    SetSimpleUser(value ItemItemStargazersSimpleUserable)()
    SetStargazer(value ItemItemStargazersStargazerable)()
}
// NewItemItemStargazersRequestBuilderInternal instantiates a new StargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    m := &ItemItemStargazersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/stargazers{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemStargazersRequestBuilder instantiates a new StargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStargazersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/activity#list-stargazers
func (m *ItemItemStargazersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStargazersRequestBuilderGetRequestConfiguration)(StargazersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateStargazersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StargazersResponseable), nil
}
// ToGetRequestInformation lists the people that have starred the repository.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
func (m *ItemItemStargazersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStargazersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
