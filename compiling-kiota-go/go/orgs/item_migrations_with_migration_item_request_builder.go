package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemMigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}
type ItemMigrationsWithMigration_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
type ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // Exclude attributes from the API response to improve performance
    Exclude []string
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters
}
// Archive the archive property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Archive()(*ItemMigrationsItemArchiveRequestBuilder) {
    return NewItemMigrationsItemArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    m := &ItemMigrationsWithMigration_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/migrations/{migration_id}{?exclude*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMigrationsWithMigration_ItemRequestBuilder instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/migrations#get-an-organization-migration-status
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateMigrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Migrationable), nil
}
// Repos the repos property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repos()(*ItemMigrationsItemReposRequestBuilder) {
    return NewItemMigrationsItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReposById gets an item from the github.com/octokit/kiota/.orgs.item.migrations.item.repos.item collection
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) ReposById(id string)(*ItemMigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["repo_name"] = id
    }
    return NewItemMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Repositories the repositories property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repositories()(*ItemMigrationsItemRepositoriesRequestBuilder) {
    return NewItemMigrationsItemRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
