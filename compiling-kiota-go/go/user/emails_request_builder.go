package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// EmailsRequestBuilder builds and executes requests for operations under \user\emails
type EmailsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EmailsDeleteRequestBody composed type wrapper for classes emailsDeleteRequestBodyMember1, string, string
type EmailsDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type emailsDeleteRequestBodyMember1
    emailsDeleteRequestBodyMember1 EmailsDeleteRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewEmailsDeleteRequestBody instantiates a new emailsDeleteRequestBody and sets the default values.
func NewEmailsDeleteRequestBody()(*EmailsDeleteRequestBody) {
    m := &EmailsDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEmailsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEmailsDeleteRequestBody()
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
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailsDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmailsDeleteRequestBodyMember1 gets the emailsDeleteRequestBodyMember1 property value. Composed type representation for type emailsDeleteRequestBodyMember1
func (m *EmailsDeleteRequestBody) GetEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able) {
    return m.emailsDeleteRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *EmailsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *EmailsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *EmailsDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmailsDeleteRequestBodyMember1 sets the emailsDeleteRequestBodyMember1 property value. Composed type representation for type emailsDeleteRequestBodyMember1
func (m *EmailsDeleteRequestBody) SetEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)() {
    m.emailsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *EmailsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// EmailsPostRequestBody composed type wrapper for classes emailsPostRequestBodyMember1, string, string
type EmailsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type emailsPostRequestBodyMember1
    emailsPostRequestBodyMember1 EmailsPostRequestBodyMember1able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewEmailsPostRequestBody instantiates a new emailsPostRequestBody and sets the default values.
func NewEmailsPostRequestBody()(*EmailsPostRequestBody) {
    m := &EmailsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEmailsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEmailsPostRequestBody()
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
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmailsPostRequestBodyMember1 gets the emailsPostRequestBodyMember1 property value. Composed type representation for type emailsPostRequestBodyMember1
func (m *EmailsPostRequestBody) GetEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able) {
    return m.emailsPostRequestBodyMember1
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetString gets the string property value. Composed type representation for type string
func (m *EmailsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *EmailsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
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
func (m *EmailsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmailsPostRequestBodyMember1 sets the emailsPostRequestBodyMember1 property value. Composed type representation for type emailsPostRequestBodyMember1
func (m *EmailsPostRequestBody) SetEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)() {
    m.emailsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *EmailsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// EmailsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmailsRequestBuilderGetQueryParameters lists all of your email addresses, and specifies which one is visible to the public. This endpoint is accessible with the `user:email` scope.
type EmailsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// EmailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmailsRequestBuilderGetQueryParameters
}
// EmailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmailsDeleteRequestBodyable 
type EmailsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
// EmailsPostRequestBodyable 
type EmailsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able)
    GetString()(*string)
    SetEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)()
    SetString(value *string)()
}
// NewEmailsRequestBuilderInternal instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailsRequestBuilder) {
    m := &EmailsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/emails{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEmailsRequestBuilder instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete this endpoint is accessible with the `user` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#delete-an-email-address-for-the-authenticated-user
func (m *EmailsRequestBuilder) Delete(ctx context.Context, body EmailsDeleteRequestBodyable, requestConfiguration *EmailsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get lists all of your email addresses, and specifies which one is visible to the public. This endpoint is accessible with the `user:email` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#list-email-addresses-for-the-authenticated-user
func (m *EmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *EmailsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEmailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable)
    }
    return val, nil
}
// Post this endpoint is accessible with the `user` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#add-an-email-address-for-the-authenticated-user
func (m *EmailsRequestBuilder) Post(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *EmailsRequestBuilderPostRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEmailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Emailable)
    }
    return val, nil
}
// ToDeleteRequestInformation this endpoint is accessible with the `user` scope.
func (m *EmailsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body EmailsDeleteRequestBodyable, requestConfiguration *EmailsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// ToGetRequestInformation lists all of your email addresses, and specifies which one is visible to the public. This endpoint is accessible with the `user:email` scope.
func (m *EmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation this endpoint is accessible with the `user` scope.
func (m *EmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *EmailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
