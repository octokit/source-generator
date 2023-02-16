package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsItemReposRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}\repos
type ItemMigrationsItemReposRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemMigrationsItemReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewItemMigrationsItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposRequestBuilder) {
    m := &ItemMigrationsItemReposRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/migrations/{migration_id}/repos";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMigrationsItemReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewItemMigrationsItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
