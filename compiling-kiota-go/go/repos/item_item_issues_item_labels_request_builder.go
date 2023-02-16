package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemIssuesItemLabelsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\issues\{issue_number}\labels
type ItemItemIssuesItemLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemIssuesItemLabelsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemLabelsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters list labels for an issue
type ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemIssuesItemLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters
}
// ItemItemIssuesItemLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemIssuesItemLabelsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemIssuesItemLabelsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LabelsPostRequestBody composed type wrapper for classes labelsPostRequestBodyMember1, string, labelsPostRequestBodyMember2, labelsPostRequestBodyMember3, string
type LabelsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type labelsPostRequestBodyMember1
    labelsPostRequestBodyMember1 ItemItemIssuesItemLabelsPostRequestBodyMember1able
    // Composed type representation for type labelsPostRequestBodyMember2
    labelsPostRequestBodyMember2 ItemItemIssuesItemLabelsPostRequestBodyMember2able
    // Composed type representation for type labelsPostRequestBodyMember3
    labelsPostRequestBodyMember3 ItemItemIssuesItemLabelsPostRequestBodyMember3able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewLabelsPostRequestBody instantiates a new labelsPostRequestBody and sets the default values.
func NewLabelsPostRequestBody()(*LabelsPostRequestBody) {
    m := &LabelsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLabelsPostRequestBody()
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
func (m *LabelsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetLabelsPostRequestBodyMember1 gets the labelsPostRequestBodyMember1 property value. Composed type representation for type labelsPostRequestBodyMember1
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able) {
    return m.labelsPostRequestBodyMember1
}
// GetLabelsPostRequestBodyMember2 gets the labelsPostRequestBodyMember2 property value. Composed type representation for type labelsPostRequestBodyMember2
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able) {
    return m.labelsPostRequestBodyMember2
}
// GetLabelsPostRequestBodyMember3 gets the labelsPostRequestBodyMember3 property value. Composed type representation for type labelsPostRequestBodyMember3
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyMember3()(ItemItemIssuesItemLabelsPostRequestBodyMember3able) {
    return m.labelsPostRequestBodyMember3
}
// GetString gets the string property value. Composed type representation for type string
func (m *LabelsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LabelsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabelsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyMember3() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPostRequestBodyMember3())
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
func (m *LabelsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabelsPostRequestBodyMember1 sets the labelsPostRequestBodyMember1 property value. Composed type representation for type labelsPostRequestBodyMember1
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)() {
    m.labelsPostRequestBodyMember1 = value
}
// SetLabelsPostRequestBodyMember2 sets the labelsPostRequestBodyMember2 property value. Composed type representation for type labelsPostRequestBodyMember2
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)() {
    m.labelsPostRequestBodyMember2 = value
}
// SetLabelsPostRequestBodyMember3 sets the labelsPostRequestBodyMember3 property value. Composed type representation for type labelsPostRequestBodyMember3
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyMember3(value ItemItemIssuesItemLabelsPostRequestBodyMember3able)() {
    m.labelsPostRequestBodyMember3 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LabelsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// LabelsPutRequestBody composed type wrapper for classes labelsPutRequestBodyMember1, string, labelsPutRequestBodyMember2, labelsPutRequestBodyMember3, string
type LabelsPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type labelsPutRequestBodyMember1
    labelsPutRequestBodyMember1 ItemItemIssuesItemLabelsPutRequestBodyMember1able
    // Composed type representation for type labelsPutRequestBodyMember2
    labelsPutRequestBodyMember2 ItemItemIssuesItemLabelsPutRequestBodyMember2able
    // Composed type representation for type labelsPutRequestBodyMember3
    labelsPutRequestBodyMember3 ItemItemIssuesItemLabelsPutRequestBodyMember3able
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type string
    string *string
}
// NewLabelsPutRequestBody instantiates a new labelsPutRequestBody and sets the default values.
func NewLabelsPutRequestBody()(*LabelsPutRequestBody) {
    m := &LabelsPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLabelsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLabelsPutRequestBody()
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
func (m *LabelsPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetLabelsPutRequestBodyMember1 gets the labelsPutRequestBodyMember1 property value. Composed type representation for type labelsPutRequestBodyMember1
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able) {
    return m.labelsPutRequestBodyMember1
}
// GetLabelsPutRequestBodyMember2 gets the labelsPutRequestBodyMember2 property value. Composed type representation for type labelsPutRequestBodyMember2
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able) {
    return m.labelsPutRequestBodyMember2
}
// GetLabelsPutRequestBodyMember3 gets the labelsPutRequestBodyMember3 property value. Composed type representation for type labelsPutRequestBodyMember3
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyMember3()(ItemItemIssuesItemLabelsPutRequestBodyMember3able) {
    return m.labelsPutRequestBodyMember3
}
// GetString gets the string property value. Composed type representation for type string
func (m *LabelsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LabelsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabelsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPutRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyMember3() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPutRequestBodyMember3())
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
func (m *LabelsPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabelsPutRequestBodyMember1 sets the labelsPutRequestBodyMember1 property value. Composed type representation for type labelsPutRequestBodyMember1
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)() {
    m.labelsPutRequestBodyMember1 = value
}
// SetLabelsPutRequestBodyMember2 sets the labelsPutRequestBodyMember2 property value. Composed type representation for type labelsPutRequestBodyMember2
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)() {
    m.labelsPutRequestBodyMember2 = value
}
// SetLabelsPutRequestBodyMember3 sets the labelsPutRequestBodyMember3 property value. Composed type representation for type labelsPutRequestBodyMember3
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyMember3(value ItemItemIssuesItemLabelsPutRequestBodyMember3able)() {
    m.labelsPutRequestBodyMember3 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LabelsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// LabelsPostRequestBodyable 
type LabelsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able)
    GetLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able)
    GetLabelsPostRequestBodyMember3()(ItemItemIssuesItemLabelsPostRequestBodyMember3able)
    GetString()(*string)
    SetLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)()
    SetLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)()
    SetLabelsPostRequestBodyMember3(value ItemItemIssuesItemLabelsPostRequestBodyMember3able)()
    SetString(value *string)()
}
// LabelsPutRequestBodyable 
type LabelsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able)
    GetLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able)
    GetLabelsPutRequestBodyMember3()(ItemItemIssuesItemLabelsPutRequestBodyMember3able)
    GetString()(*string)
    SetLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)()
    SetLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)()
    SetLabelsPutRequestBodyMember3(value ItemItemIssuesItemLabelsPutRequestBodyMember3able)()
    SetString(value *string)()
}
// NewItemItemIssuesItemLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemIssuesItemLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemLabelsRequestBuilder) {
    m := &ItemItemIssuesItemLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/issues/{issue_number}/labels{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemIssuesItemLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemIssuesItemLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove all labels from an issue
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#remove-all-labels-from-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list labels for an issue
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#list-labels-for-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderGetRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable)
    }
    return val, nil
}
// Post add labels to an issue
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#add-labels-to-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Post(ctx context.Context, body LabelsPostRequestBodyable, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderPostRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable)
    }
    return val, nil
}
// Put removes any previous labels and sets the new labels for an issue.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/issues#set-labels-for-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Put(ctx context.Context, body LabelsPutRequestBodyable, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderPutRequestConfiguration)([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "410": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable, len(res))
    for i, v := range res {
        val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Labelable)
    }
    return val, nil
}
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body LabelsPostRequestBodyable, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation removes any previous labels and sets the new labels for an issue.
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToPutRequestInformation(ctx context.Context, body LabelsPutRequestBodyable, requestConfiguration *ItemItemIssuesItemLabelsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
