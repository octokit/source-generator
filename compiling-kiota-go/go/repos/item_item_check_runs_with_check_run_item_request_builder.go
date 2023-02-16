package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCheckRunsWithCheck_run_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-runs\{check_run_id}
type ItemItemCheckRunsWithCheck_run_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCheckRunsWithCheck_run_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckRunsWithCheck_run_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemCheckRunsWithCheck_run_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckRunsWithCheck_run_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithCheck_run_PatchRequestBody composed type wrapper for classes WithCheck_run_PatchRequestBodyMember1, WithCheck_run_PatchRequestBodyMember2
type WithCheck_run_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Serialization hint for the current wrapper.
    SerializationHint *string
    // Composed type representation for type WithCheck_run_PatchRequestBodyMember1
    withCheck_run_PatchRequestBodyMember1 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
    // Composed type representation for type WithCheck_run_PatchRequestBodyMember2
    withCheck_run_PatchRequestBodyMember2 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
}
// NewWithCheck_run_PatchRequestBody instantiates a new WithCheck_run_PatchRequestBody and sets the default values.
func NewWithCheck_run_PatchRequestBody()(*WithCheck_run_PatchRequestBody) {
    m := &WithCheck_run_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWithCheck_run_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithCheck_run_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithCheck_run_PatchRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able); ok {
                result.SetWithCheckRunPatchRequestBodyMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able); ok {
                result.SetWithCheckRunPatchRequestBodyMember2(cast)
            }
        }
    }
    return result, nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WithCheck_run_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WithCheck_run_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetWithCheckRunPatchRequestBodyMember1 gets the withCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type WithCheck_run_PatchRequestBodyMember1
func (m *WithCheck_run_PatchRequestBody) GetWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able) {
    return m.withCheck_run_PatchRequestBodyMember1
}
// GetWithCheckRunPatchRequestBodyMember2 gets the withCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type WithCheck_run_PatchRequestBodyMember2
func (m *WithCheck_run_PatchRequestBody) GetWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able) {
    return m.withCheck_run_PatchRequestBodyMember2
}
// Serialize serializes information the current object
func (m *WithCheck_run_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWithCheckRunPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetWithCheckRunPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetWithCheckRunPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetWithCheckRunPatchRequestBodyMember2())
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
func (m *WithCheck_run_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWithCheckRunPatchRequestBodyMember1 sets the withCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type WithCheck_run_PatchRequestBodyMember1
func (m *WithCheck_run_PatchRequestBody) SetWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)() {
    m.withCheck_run_PatchRequestBodyMember1 = value
}
// SetWithCheckRunPatchRequestBodyMember2 sets the withCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type WithCheck_run_PatchRequestBodyMember2
func (m *WithCheck_run_PatchRequestBody) SetWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)() {
    m.withCheck_run_PatchRequestBodyMember2 = value
}
// WithCheck_run_PatchRequestBodyable 
type WithCheck_run_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)
    GetWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)
    SetWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)()
    SetWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)()
}
// Annotations the annotations property
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Annotations()(*ItemItemCheckRunsItemAnnotationsRequestBuilder) {
    return NewItemItemCheckRunsItemAnnotationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal instantiates a new WithCheck_run_ItemRequestBuilder and sets the default values.
func NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    m := &ItemItemCheckRunsWithCheck_run_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/check-runs/{check_run_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCheckRunsWithCheck_run_ItemRequestBuilder instantiates a new WithCheck_run_ItemRequestBuilder and sets the default values.
func NewItemItemCheckRunsWithCheck_run_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Gets a single check run using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#get-a-check-run
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCheckRunsWithCheck_run_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Updates a check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to edit check runs.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/checks#update-a-check-run
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Patch(ctx context.Context, body WithCheck_run_PatchRequestBodyable, requestConfiguration *ItemItemCheckRunsWithCheck_run_ItemRequestBuilderPatchRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CheckRunable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Rerequest the rerequest property
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Rerequest()(*ItemItemCheckRunsItemRerequestRequestBuilder) {
    return NewItemItemCheckRunsItemRerequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Gets a single check run using its `id`. GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to get check runs. OAuth Apps and authenticated users must have the `repo` scope to get check runs in a private repository.
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCheckRunsWithCheck_run_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.Updates a check run for a specific commit in a repository. Your GitHub App must have the `checks:write` permission to edit check runs.
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body WithCheck_run_PatchRequestBodyable, requestConfiguration *ItemItemCheckRunsWithCheck_run_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
