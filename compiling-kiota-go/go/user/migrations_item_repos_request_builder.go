package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsItemReposRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}\repos
type MigrationsItemReposRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewMigrationsItemReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewMigrationsItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposRequestBuilder) {
    m := &MigrationsItemReposRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/migrations/{migration_id}/repos";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMigrationsItemReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewMigrationsItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
