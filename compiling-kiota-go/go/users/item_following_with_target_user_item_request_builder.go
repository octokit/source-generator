package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemFollowingWithTarget_userItemRequestBuilder builds and executes requests for operations under \users\{username}\following\{target_user}
type ItemFollowingWithTarget_userItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFollowingWithTarget_userItemRequestBuilderInternal instantiates a new WithTarget_userItemRequestBuilder and sets the default values.
func NewItemFollowingWithTarget_userItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    m := &ItemFollowingWithTarget_userItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}/following/{target_user}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemFollowingWithTarget_userItemRequestBuilder instantiates a new WithTarget_userItemRequestBuilder and sets the default values.
func NewItemFollowingWithTarget_userItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowingWithTarget_userItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check if a user follows another user
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/users#check-if-a-user-follows-another-user
func (m *ItemFollowingWithTarget_userItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ItemFollowingWithTarget_userItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
