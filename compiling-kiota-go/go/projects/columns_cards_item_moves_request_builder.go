package projects

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ColumnsCardsItemMovesRequestBuilder builds and executes requests for operations under \projects\columns\cards\{card_id}\moves
type ColumnsCardsItemMovesRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewColumnsCardsItemMovesRequestBuilderInternal instantiates a new MovesRequestBuilder and sets the default values.
func NewColumnsCardsItemMovesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsCardsItemMovesRequestBuilder {
	m := &ColumnsCardsItemMovesRequestBuilder{}
	m.urlTemplate = "{+baseurl}/projects/columns/cards/{card_id}/moves"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewColumnsCardsItemMovesRequestBuilder instantiates a new MovesRequestBuilder and sets the default values.
func NewColumnsCardsItemMovesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsCardsItemMovesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewColumnsCardsItemMovesRequestBuilderInternal(urlParams, requestAdapter)
}

// Post move a project card
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#move-a-project-card
func (m *ColumnsCardsItemMovesRequestBuilder) Post(ctx context.Context, body ColumnsCardsItemMovesPostRequestBodyable, requestConfiguration *ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration) (ColumnsCardsItemMovesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
		"503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, CreateColumnsCardsItemMovesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ColumnsCardsItemMovesResponseable), nil
}
func (m *ColumnsCardsItemMovesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ColumnsCardsItemMovesPostRequestBodyable, requestConfiguration *ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
