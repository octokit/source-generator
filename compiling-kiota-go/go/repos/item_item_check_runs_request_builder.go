package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-runs
type ItemItemCheckRunsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CheckRunsPostRequestBody composed type wrapper for classes checkRunsPostRequestBodyMember1, checkRunsPostRequestBodyMember2
type CheckRunsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type checkRunsPostRequestBodyMember1
    checkRunsPostRequestBodyMember1 ItemItemCheckRunsPostRequestBodyMember1able
    // Composed type representation for type checkRunsPostRequestBodyMember2
    checkRunsPostRequestBodyMember2 ItemItemCheckRunsPostRequestBodyMember2able
    // Serialization hint for the current wrapper.
    SerializationHint *string
}
// NewCheckRunsPostRequestBody instantiates a new checkRunsPostRequestBody and sets the default values.
func NewCheckRunsPostRequestBody()(*CheckRunsPostRequestBody) {
    m := &CheckRunsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckRunsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCheckRunsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCheckRunsPostRequestBody()
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
func (m *CheckRunsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckRunsPostRequestBodyMember1 gets the checkRunsPostRequestBodyMember1 property value. Composed type representation for type checkRunsPostRequestBodyMember1
func (m *CheckRunsPostRequestBody) GetCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able) {
    return m.checkRunsPostRequestBodyMember1
}
// GetCheckRunsPostRequestBodyMember2 gets the checkRunsPostRequestBodyMember2 property value. Composed type representation for type checkRunsPostRequestBodyMember2
func (m *CheckRunsPostRequestBody) GetCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able) {
    return m.checkRunsPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CheckRunsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// Serialize serializes information the current object
func (m *CheckRunsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckRunsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCheckRunsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCheckRunsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCheckRunsPostRequestBodyMember2())
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
func (m *CheckRunsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckRunsPostRequestBodyMember1 sets the checkRunsPostRequestBodyMember1 property value. Composed type representation for type checkRunsPostRequestBodyMember1
func (m *CheckRunsPostRequestBody) SetCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)() {
    m.checkRunsPostRequestBodyMember1 = value
}
// SetCheckRunsPostRequestBodyMember2 sets the checkRunsPostRequestBodyMember2 property value. Composed type representation for type checkRunsPostRequestBodyMember2
func (m *CheckRunsPostRequestBody) SetCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)() {
    m.checkRunsPostRequestBodyMember2 = value
}
// ItemItemCheckRunsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckRunsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckRunsPostRequestBodyable 
type CheckRunsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able)
    GetCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able)
    SetCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)()
    SetCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)()
}
// NewItemItemCheckRunsRequestBuilderInternal instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsRequestBuilder) {
    m := &ItemItemCheckRunsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/check-runs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCheckRunsRequestBuilder instantiates a new CheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Creates a new check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to create check runs.In a check suite, GitHub limits the number of check runs with the same name to 1000. Once these check runs exceed 1000, GitHub will start to automatically delete older check runs.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#create-a-check-run
func (m *ItemItemCheckRunsRequestBuilder) Post(ctx context.Context, body CheckRunsPostRequestBodyable, requestConfiguration *ItemItemCheckRunsRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCheckRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable), nil
}
// ToPostRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Creates a new check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to create check runs.In a check suite, GitHub limits the number of check runs with the same name to 1000. Once these check runs exceed 1000, GitHub will start to automatically delete older check runs.
func (m *ItemItemCheckRunsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CheckRunsPostRequestBodyable, requestConfiguration *ItemItemCheckRunsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
