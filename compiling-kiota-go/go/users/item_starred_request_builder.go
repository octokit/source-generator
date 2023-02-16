package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemStarredRequestBuilder builds and executes requests for operations under \users\{username}\starred
type ItemStarredRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemStarredRequestBuilderGetQueryParameters lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
type ItemStarredRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *string
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
    // The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to.
    Sort *string
}
// ItemStarredRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemStarredRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemStarredRequestBuilderGetQueryParameters
}
// StarredResponse composed type wrapper for classes starredRepository, repository
type StarredResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type repository
    repository ItemStarredRepositoryable
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type starredRepository
    starredRepository ItemStarredRepositoryable
}
// NewStarredResponse instantiates a new starredResponse and sets the default values.
func NewStarredResponse()(*StarredResponse) {
    m := &StarredResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStarredResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStarredResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStarredResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemStarredRepositoryable); ok {
                result.SetRepository(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemStarredRepositoryFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemStarredRepositoryable); ok {
                result.SetStarredRepository(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StarredResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StarredResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetRepository gets the repository property value. Composed type representation for type repository
func (m *StarredResponse) GetRepository()(ItemStarredRepositoryable) {
    return m.repository
}
// GetStarredRepository gets the starredRepository property value. Composed type representation for type starredRepository
func (m *StarredResponse) GetStarredRepository()(ItemStarredRepositoryable) {
    return m.starredRepository
}
// Serialize serializes information the current object
func (m *StarredResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepository() != nil {
        err := writer.WriteObjectValue("", m.GetRepository())
        if err != nil {
            return err
        }
    } else if m.GetStarredRepository() != nil {
        err := writer.WriteObjectValue("", m.GetStarredRepository())
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
func (m *StarredResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepository sets the repository property value. Composed type representation for type repository
func (m *StarredResponse) SetRepository(value ItemStarredRepositoryable)() {
    m.repository = value
}
// SetStarredRepository sets the starredRepository property value. Composed type representation for type starredRepository
func (m *StarredResponse) SetStarredRepository(value ItemStarredRepositoryable)() {
    m.starredRepository = value
}
// StarredResponseable 
type StarredResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepository()(ItemStarredRepositoryable)
    GetStarredRepository()(ItemStarredRepositoryable)
    SetRepository(value ItemStarredRepositoryable)()
    SetStarredRepository(value ItemStarredRepositoryable)()
}
// NewItemStarredRequestBuilderInternal instantiates a new StarredRequestBuilder and sets the default values.
func NewItemStarredRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStarredRequestBuilder) {
    m := &ItemStarredRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}/starred{?sort*,direction*,per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemStarredRequestBuilder instantiates a new StarredRequestBuilder and sets the default values.
func NewItemStarredRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemStarredRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemStarredRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/activity#list-repositories-starred-by-a-user
func (m *ItemStarredRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemStarredRequestBuilderGetRequestConfiguration)(StarredResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateStarredResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StarredResponseable), nil
}
// ToGetRequestInformation lists repositories a user has starred.You can also find out _when_ stars were created by passing the following custom [media type](https://docs.github.com/rest/overview/media-types/) via the `Accept` header: `application/vnd.github.star+json`.
func (m *ItemStarredRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemStarredRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
