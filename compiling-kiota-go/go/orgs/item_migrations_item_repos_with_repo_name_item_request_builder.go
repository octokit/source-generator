package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsItemReposWithRepo_nameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}\repos\{repo_name}
type ItemMigrationsItemReposWithRepo_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal instantiates a new WithRepo_nameItemRequestBuilder and sets the default values.
func NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    m := &ItemMigrationsItemReposWithRepo_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/migrations/{migration_id}/repos/{repo_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMigrationsItemReposWithRepo_nameItemRequestBuilder instantiates a new WithRepo_nameItemRequestBuilder and sets the default values.
func NewItemMigrationsItemReposWithRepo_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Lock the lock property
func (m *ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) Lock()(*ItemMigrationsItemReposItemLockRequestBuilder) {
    return NewItemMigrationsItemReposItemLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
