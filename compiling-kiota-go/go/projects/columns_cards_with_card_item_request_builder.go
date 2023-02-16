package projects

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ColumnsCardsWithCard_ItemRequestBuilder builds and executes requests for operations under \projects\columns\cards\{card_id}
type ColumnsCardsWithCard_ItemRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ColumnsCardsWithCard_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsCardsWithCard_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ColumnsCardsWithCard_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsCardsWithCard_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ColumnsCardsWithCard_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsCardsWithCard_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewColumnsCardsWithCard_ItemRequestBuilderInternal instantiates a new WithCard_ItemRequestBuilder and sets the default values.
func NewColumnsCardsWithCard_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsCardsWithCard_ItemRequestBuilder {
	m := &ColumnsCardsWithCard_ItemRequestBuilder{}
	m.urlTemplate = "{+baseurl}/projects/columns/cards/{card_id}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewColumnsCardsWithCard_ItemRequestBuilder instantiates a new WithCard_ItemRequestBuilder and sets the default values.
func NewColumnsCardsWithCard_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsCardsWithCard_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewColumnsCardsWithCard_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a project card
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#delete-a-project-card
func (m *ColumnsCardsWithCard_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get get a project card
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#get-a-project-card
func (m *ColumnsCardsWithCard_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderGetRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateProjectCardFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable), nil
}

// Moves the moves property
func (m *ColumnsCardsWithCard_ItemRequestBuilder) Moves() *ColumnsCardsItemMovesRequestBuilder {
	return NewColumnsCardsItemMovesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Patch update an existing project card
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#update-a-project-card
func (m *ColumnsCardsWithCard_ItemRequestBuilder) Patch(ctx context.Context, body ColumnsCardsItemWithCard_PatchRequestBodyable, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderPatchRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorSimpleFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateProjectCardFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable), nil
}
func (m *ColumnsCardsWithCard_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ColumnsCardsWithCard_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ColumnsCardsWithCard_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ColumnsCardsItemWithCard_PatchRequestBodyable, requestConfiguration *ColumnsCardsWithCard_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
