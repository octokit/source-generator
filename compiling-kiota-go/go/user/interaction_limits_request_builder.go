package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// InteractionLimitsRequestBuilder builds and executes requests for operations under \user\interaction-limits
type InteractionLimitsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InteractionLimitsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsResponse composed type wrapper for classes interactionLimitResponse, interactionLimitsResponseMember1
type InteractionLimitsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable
    // Composed type representation for type interactionLimitsResponseMember1
    interactionLimitsResponseMember1 InteractionLimitsResponseMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewInteractionLimitsResponse instantiates a new interactionLimitsResponse and sets the default values.
func NewInteractionLimitsResponse()(*InteractionLimitsResponse) {
    m := &InteractionLimitsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInteractionLimitsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInteractionLimitsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewInteractionLimitsResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateInteractionLimitResponseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable); ok {
                result.SetInteractionLimitResponse(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateInteractionLimitsResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(InteractionLimitsResponseMember1able); ok {
                result.SetInteractionLimitsResponseMember1(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InteractionLimitsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InteractionLimitsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteractionLimitResponse gets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsResponse) GetInteractionLimitResponse()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable) {
    return m.interactionLimitResponse
}
// GetInteractionLimitsResponseMember1 gets the interactionLimitsResponseMember1 property value. Composed type representation for type interactionLimitsResponseMember1
func (m *InteractionLimitsResponse) GetInteractionLimitsResponseMember1()(InteractionLimitsResponseMember1able) {
    return m.interactionLimitsResponseMember1
}
// Serialize serializes information the current object
func (m *InteractionLimitsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetInteractionLimitsResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitsResponseMember1())
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
func (m *InteractionLimitsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInteractionLimitResponse sets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsResponse) SetInteractionLimitResponse(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable)() {
    m.interactionLimitResponse = value
}
// SetInteractionLimitsResponseMember1 sets the interactionLimitsResponseMember1 property value. Composed type representation for type interactionLimitsResponseMember1
func (m *InteractionLimitsResponse) SetInteractionLimitsResponseMember1(value InteractionLimitsResponseMember1able)() {
    m.interactionLimitsResponseMember1 = value
}
// InteractionLimitsResponseable 
type InteractionLimitsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable)
    GetInteractionLimitsResponseMember1()(InteractionLimitsResponseMember1able)
    SetInteractionLimitResponse(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable)()
    SetInteractionLimitsResponseMember1(value InteractionLimitsResponseMember1able)()
}
// NewInteractionLimitsRequestBuilderInternal instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewInteractionLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InteractionLimitsRequestBuilder) {
    m := &InteractionLimitsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/interaction-limits";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInteractionLimitsRequestBuilder instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewInteractionLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InteractionLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInteractionLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes any interaction restrictions from your public repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/interactions#remove-interaction-restrictions-from-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Delete(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get shows which type of GitHub user can interact with your public repositories and when the restriction expires.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/interactions#get-interaction-restrictions-for-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Get(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateInteractionLimitsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InteractionLimitsResponseable), nil
}
// Put temporarily restricts which type of GitHub user can interact with your public repositories. Setting the interaction limit at the user level will overwrite any interaction limits that are set for individual repositories owned by the user.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/interactions#set-interaction-restrictions-for-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Put(ctx context.Context, body ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitable, requestConfiguration *InteractionLimitsRequestBuilderPutRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateInteractionLimitResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitResponseable), nil
}
// ToDeleteRequestInformation removes any interaction restrictions from your public repositories.
func (m *InteractionLimitsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation shows which type of GitHub user can interact with your public repositories and when the restriction expires.
func (m *InteractionLimitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation temporarily restricts which type of GitHub user can interact with your public repositories. Setting the interaction limit at the user level will overwrite any interaction limits that are set for individual repositories owned by the user.
func (m *InteractionLimitsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.InteractionLimitable, requestConfiguration *InteractionLimitsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
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
