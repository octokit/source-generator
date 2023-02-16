package repositories

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemEnvironmentsItemVariablesRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\variables
type ItemEnvironmentsItemVariablesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters lists all environment variables. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `environments:read` repository permission to use this endpoint.
type ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 30).
    Per_page *int32
}
// ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters
}
// ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEnvironmentsItemVariablesRequestBuilderInternal instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesRequestBuilder) {
    m := &ItemEnvironmentsItemVariablesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/variables{?per_page*,page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEnvironmentsItemVariablesRequestBuilder instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemVariablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all environment variables. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `environments:read` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-environment-variables
func (m *ItemEnvironmentsItemVariablesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration)(ItemEnvironmentsItemVariablesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemVariablesResponseable), nil
}
// Post create an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#create-an-environment-variable
func (m *ItemEnvironmentsItemVariablesRequestBuilder) Post(ctx context.Context, body ItemEnvironmentsItemVariablesPostRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.EmptyObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEmptyObjectFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.EmptyObjectable), nil
}
// ToGetRequestInformation lists all environment variables. You must authenticate using an access token with the `repo` scope to use this endpoint. GitHub Apps must have the `environments:read` repository permission to use this endpoint.
func (m *ItemEnvironmentsItemVariablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.
func (m *ItemEnvironmentsItemVariablesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemEnvironmentsItemVariablesPostRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
