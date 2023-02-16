package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailRequestBuilder builds and executes requests for operations under \user\email
type EmailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEmailRequestBuilderInternal instantiates a new EmailRequestBuilder and sets the default values.
func NewEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailRequestBuilder) {
    m := &EmailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/email";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEmailRequestBuilder instantiates a new EmailRequestBuilder and sets the default values.
func NewEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Visibility the visibility property
func (m *EmailRequestBuilder) Visibility()(*EmailVisibilityRequestBuilder) {
    return NewEmailVisibilityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
