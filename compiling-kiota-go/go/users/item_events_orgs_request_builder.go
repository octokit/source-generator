package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEventsOrgsRequestBuilder builds and executes requests for operations under \users\{username}\events\orgs
type ItemEventsOrgsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemEventsOrgsRequestBuilderInternal instantiates a new OrgsRequestBuilder and sets the default values.
func NewItemEventsOrgsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsRequestBuilder) {
    m := &ItemEventsOrgsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{username}/events/orgs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemEventsOrgsRequestBuilder instantiates a new OrgsRequestBuilder and sets the default values.
func NewItemEventsOrgsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsOrgsRequestBuilderInternal(urlParams, requestAdapter)
}
