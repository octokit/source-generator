package markdown

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MarkdownRequestBuilder builds and executes requests for operations under \markdown
type MarkdownRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MarkdownRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MarkdownRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMarkdownRequestBuilderInternal instantiates a new MarkdownRequestBuilder and sets the default values.
func NewMarkdownRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MarkdownRequestBuilder) {
    m := &MarkdownRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/markdown";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMarkdownRequestBuilder instantiates a new MarkdownRequestBuilder and sets the default values.
func NewMarkdownRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MarkdownRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMarkdownRequestBuilderInternal(urlParams, requestAdapter)
}
// Post render a Markdown document
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/markdown#render-a-markdown-document
func (m *MarkdownRequestBuilder) Post(ctx context.Context, body MarkdownPostRequestBodyable, requestConfiguration *MarkdownRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Raw the raw property
func (m *MarkdownRequestBuilder) Raw()(*RawRequestBuilder) {
    return NewRawRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
func (m *MarkdownRequestBuilder) ToPostRequestInformation(ctx context.Context, body MarkdownPostRequestBodyable, requestConfiguration *MarkdownRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
