package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// MigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}
type MigrationsWithMigration_ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/reference/migrations#download-a-user-migration-archive).
type MigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // 
    Exclude []string
}
// MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MigrationsWithMigration_ItemRequestBuilderGetQueryParameters
}
// Archive the archive property
func (m *MigrationsWithMigration_ItemRequestBuilder) Archive()(*MigrationsItemArchiveRequestBuilder) {
    return NewMigrationsItemArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    m := &MigrationsWithMigration_ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/migrations/{migration_id}{?exclude*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMigrationsWithMigration_ItemRequestBuilder instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/reference/migrations#download-a-user-migration-archive).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/migrations#get-a-user-migration-status
func (m *MigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "403": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
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
func (m *MigrationsWithMigration_ItemRequestBuilder) Repos()(*MigrationsItemReposRequestBuilder) {
    return NewMigrationsItemReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReposById gets an item from the github.com/octokit/kiota/.user.migrations.item.repos.item collection
func (m *MigrationsWithMigration_ItemRequestBuilder) ReposById(id string)(*MigrationsItemReposWithRepo_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["repo_name"] = id
    }
    return NewMigrationsItemReposWithRepo_nameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Repositories the repositories property
func (m *MigrationsWithMigration_ItemRequestBuilder) Repositories()(*MigrationsItemRepositoriesRequestBuilder) {
    return NewMigrationsItemRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/reference/migrations#download-a-user-migration-archive).
func (m *MigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
