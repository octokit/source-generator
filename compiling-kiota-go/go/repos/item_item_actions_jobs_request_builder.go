package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsJobsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\jobs
type ItemItemActionsJobsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemItemActionsJobsRequestBuilderInternal instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    m := &ItemItemActionsJobsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/actions/jobs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemActionsJobsRequestBuilder instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsJobsRequestBuilderInternal(urlParams, requestAdapter)
}
