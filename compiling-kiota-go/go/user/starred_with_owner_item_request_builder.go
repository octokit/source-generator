package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StarredWithOwnerItemRequestBuilder builds and executes requests for operations under \user\starred\{owner}
type StarredWithOwnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewStarredWithOwnerItemRequestBuilderInternal instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewStarredWithOwnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredWithOwnerItemRequestBuilder) {
    m := &StarredWithOwnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/starred/{owner}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewStarredWithOwnerItemRequestBuilder instantiates a new WithOwnerItemRequestBuilder and sets the default values.
func NewStarredWithOwnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StarredWithOwnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStarredWithOwnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
