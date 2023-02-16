package projects

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ColumnsItemCardsRequestBuilder builds and executes requests for operations under \projects\columns\{column_id}\cards
type ColumnsItemCardsRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// CardsPostRequestBody composed type wrapper for classes cardsPostRequestBodyMember1, cardsPostRequestBodyMember2
type CardsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Composed type representation for type cardsPostRequestBodyMember1
	cardsPostRequestBodyMember1 ColumnsItemCardsPostRequestBodyMember1able
	// Composed type representation for type cardsPostRequestBodyMember2
	cardsPostRequestBodyMember2 ColumnsItemCardsPostRequestBodyMember2able
	// Serialization hint for the current wrapper.
	SerializationHint *string
}

// NewCardsPostRequestBody instantiates a new cardsPostRequestBody and sets the default values.
func NewCardsPostRequestBody() *CardsPostRequestBody {
	m := &CardsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCardsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCardsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	result := NewCardsPostRequestBody()
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
func (m *CardsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCardsPostRequestBodyMember1 gets the cardsPostRequestBodyMember1 property value. Composed type representation for type cardsPostRequestBodyMember1
func (m *CardsPostRequestBody) GetCardsPostRequestBodyMember1() ColumnsItemCardsPostRequestBodyMember1able {
	return m.cardsPostRequestBodyMember1
}

// GetCardsPostRequestBodyMember2 gets the cardsPostRequestBodyMember2 property value. Composed type representation for type cardsPostRequestBodyMember2
func (m *CardsPostRequestBody) GetCardsPostRequestBodyMember2() ColumnsItemCardsPostRequestBodyMember2able {
	return m.cardsPostRequestBodyMember2
}

// GetFieldDeserializers the deserialization information for the current model
func (m *CardsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
}

// Serialize serializes information the current object
func (m *CardsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCardsPostRequestBodyMember1() != nil {
		err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyMember1())
		if err != nil {
			return err
		}
	} else if m.GetCardsPostRequestBodyMember2() != nil {
		err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyMember2())
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
func (m *CardsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCardsPostRequestBodyMember1 sets the cardsPostRequestBodyMember1 property value. Composed type representation for type cardsPostRequestBodyMember1
func (m *CardsPostRequestBody) SetCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able) {
	m.cardsPostRequestBodyMember1 = value
}

// SetCardsPostRequestBodyMember2 sets the cardsPostRequestBodyMember2 property value. Composed type representation for type cardsPostRequestBodyMember2
func (m *CardsPostRequestBody) SetCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able) {
	m.cardsPostRequestBodyMember2 = value
}

// ColumnsItemCardsRequestBuilderGetQueryParameters list project cards
type ColumnsItemCardsRequestBuilderGetQueryParameters struct {
	// Filters the project cards that are returned by the card's state.
	Archived_state *string
	// Page number of the results to fetch.
	Page *int32
	// The number of results per page (max 100).
	Per_page *int32
}

// ColumnsItemCardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsItemCardsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ColumnsItemCardsRequestBuilderGetQueryParameters
}

// ColumnsItemCardsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsItemCardsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ProjectCard422Error composed type wrapper for classes validationError, validationErrorSimple
type ProjectCard422Error struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Serialization hint for the current wrapper.
	SerializationHint *string
	// Composed type representation for type validationError
	validationError ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable
	// Composed type representation for type validationErrorSimple
	validationErrorSimple ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable
}

// NewProjectCard422Error instantiates a new ProjectCard422Error and sets the default values.
func NewProjectCard422Error() *ProjectCard422Error {
	m := &ProjectCard422Error{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProjectCard422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProjectCard422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	result := NewProjectCard422Error()
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
func (m *ProjectCard422Error) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ProjectCard422Error) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
}

// GetValidationError gets the validationError property value. Composed type representation for type validationError
func (m *ProjectCard422Error) GetValidationError() ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable {
	return m.validationError
}

// GetValidationErrorSimple gets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *ProjectCard422Error) GetValidationErrorSimple() ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable {
	return m.validationErrorSimple
}

// Serialize serializes information the current object
func (m *ProjectCard422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetValidationError() != nil {
		err := writer.WriteObjectValue("", m.GetValidationError())
		if err != nil {
			return err
		}
	} else if m.GetValidationErrorSimple() != nil {
		err := writer.WriteObjectValue("", m.GetValidationErrorSimple())
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
func (m *ProjectCard422Error) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetValidationError sets the validationError property value. Composed type representation for type validationError
func (m *ProjectCard422Error) SetValidationError(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable) {
	m.validationError = value
}

// SetValidationErrorSimple sets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *ProjectCard422Error) SetValidationErrorSimple(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable) {
	m.validationErrorSimple = value
}

// CardsPostRequestBodyable
type CardsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCardsPostRequestBodyMember1() ColumnsItemCardsPostRequestBodyMember1able
	GetCardsPostRequestBodyMember2() ColumnsItemCardsPostRequestBodyMember2able
	SetCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)
	SetCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)
}

// ProjectCard422Errorable
type ProjectCard422Errorable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValidationError() ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable
	GetValidationErrorSimple() ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable
	SetValidationError(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorable)
	SetValidationErrorSimple(value ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ValidationErrorSimpleable)
}

// NewColumnsItemCardsRequestBuilderInternal instantiates a new CardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsItemCardsRequestBuilder {
	m := &ColumnsItemCardsRequestBuilder{}
	m.urlTemplate = "{+baseurl}/projects/columns/{column_id}/cards{?archived_state*,per_page*,page*}"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewColumnsItemCardsRequestBuilder instantiates a new CardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ColumnsItemCardsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewColumnsItemCardsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list project cards
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#list-project-cards
func (m *ColumnsItemCardsRequestBuilder) Get(ctx context.Context, requestConfiguration *ColumnsItemCardsRequestBuilderGetRequestConfiguration) ([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
	}
	res, err := m.requestAdapter.SendCollection(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateProjectCardFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable, len(res))
	for i, v := range res {
		val[i] = v.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable)
	}
	return val, nil
}

// Post create a project card
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/reference/projects#create-a-project-card
func (m *ColumnsItemCardsRequestBuilder) Post(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *ColumnsItemCardsRequestBuilderPostRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.ProjectCardable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
		"503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
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
func (m *ColumnsItemCardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ColumnsItemCardsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ColumnsItemCardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *ColumnsItemCardsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
