package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemActionsRunnersItemLabelsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\{runner_id}\labels
type ItemItemActionsRunnersItemLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemActionsRunnersItemLabelsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersItemLabelsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsRunnersItemLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersItemLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsRunnersItemLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersItemLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsRunnersItemLabelsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersItemLabelsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersItemLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsRequestBuilder) {
    m := &ItemItemActionsRunnersItemLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/runners/{runner_id}/labels";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsRunnersItemLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersItemLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove all custom labels from a self-hosted runner configured in arepository. Returns the remaining read-only labels from the runner.You must authenticate using an access token with the `repo` scope to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#remove-all-custom-labels-from-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderDeleteRequestConfiguration)(ItemItemActionsRunnersItemLabelsResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsResponseable), nil
}
// Get lists all labels for a self-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#list-labels-for-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderGetRequestConfiguration)(ItemItemActionsRunnersItemLabelsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsResponseable), nil
}
// Post add custom labels to a self-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#add-custom-labels-to-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) Post(ctx context.Context, body ItemItemActionsRunnersItemLabelsPostRequestBodyable, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderPostRequestConfiguration)(ItemItemActionsRunnersItemLabelsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsResponseable), nil
}
// Put remove all previous custom labels and set the new custom labels for a specificself-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/actions#set-custom-labels-for-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) Put(ctx context.Context, body ItemItemActionsRunnersItemLabelsPutRequestBodyable, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderPutRequestConfiguration)(ItemItemActionsRunnersItemLabelsResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsResponseable), nil
}
// ToDeleteRequestInformation remove all custom labels from a self-hosted runner configured in arepository. Returns the remaining read-only labels from the runner.You must authenticate using an access token with the `repo` scope to use thisendpoint.
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation lists all labels for a self-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add custom labels to a self-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemActionsRunnersItemLabelsPostRequestBodyable, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation remove all previous custom labels and set the new custom labels for a specificself-hosted runner configured in a repository.You must authenticate using an access token with the `repo` scope to use thisendpoint.
func (m *ItemItemActionsRunnersItemLabelsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemActionsRunnersItemLabelsPutRequestBodyable, requestConfiguration *ItemItemActionsRunnersItemLabelsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
