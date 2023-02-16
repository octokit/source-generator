package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemCompareWithBaseheadItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\compare\{basehead}
type ItemItemCompareWithBaseheadItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemCompareWithBaseheadItemRequestBuilderGetQueryParameters compares two commits against one another. You can compare branches in the same repository, or you can compare branches that exist in different repositories within the same repository network, including fork branches. For more information about how to view a repository's network, see "[Understanding connections between repositories](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository/understanding-connections-between-repositories)."This endpoint is equivalent to running the `git log BASE...HEAD` command, but it returns commits in a different order. The `git log BASE...HEAD` command returns commits in reverse chronological order, whereas the API returns commits in chronological order. You can pass the appropriate [media type](https://docs.github.com/rest/overview/media-types/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.The API response includes details about the files that were changed between the two commits. This includes the status of the change (if a file was added, removed, modified, or renamed), and details of the change itself. For example, files with a `renamed` status have a `previous_filename` field showing the previous filename of the file, and files with a `modified` status have a `patch` field showing the changes made to the file.When calling this endpoint without any paging parameter (`per_page` or `page`), the returned list is limited to 250 commits, and the last commit in the list is the most recent of the entire comparison.**Working with large comparisons**To process a response with a large number of commits, use a query parameter (`per_page` or `page`) to paginate the results. When using pagination:- The list of changed files is only shown on the first page of results, but it includes all changed files for the entire comparison.- The results are returned in chronological order, but the last commit in the returned list may not be the most recent one in the entire set if there are more pages of results.For more information on working with pagination, see "[Using pagination in the REST API](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api)."**Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The `verification` object includes the following fields:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
type ItemItemCompareWithBaseheadItemRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32
    // The number of results per page (max 100).
    Per_page *int32
}
// ItemItemCompareWithBaseheadItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCompareWithBaseheadItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCompareWithBaseheadItemRequestBuilderGetQueryParameters
}
// NewItemItemCompareWithBaseheadItemRequestBuilderInternal instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemCompareWithBaseheadItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareWithBaseheadItemRequestBuilder) {
    m := &ItemItemCompareWithBaseheadItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/compare/{basehead}{?page*,per_page*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemCompareWithBaseheadItemRequestBuilder instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemCompareWithBaseheadItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCompareWithBaseheadItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCompareWithBaseheadItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get compares two commits against one another. You can compare branches in the same repository, or you can compare branches that exist in different repositories within the same repository network, including fork branches. For more information about how to view a repository's network, see "[Understanding connections between repositories](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository/understanding-connections-between-repositories)."This endpoint is equivalent to running the `git log BASE...HEAD` command, but it returns commits in a different order. The `git log BASE...HEAD` command returns commits in reverse chronological order, whereas the API returns commits in chronological order. You can pass the appropriate [media type](https://docs.github.com/rest/overview/media-types/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.The API response includes details about the files that were changed between the two commits. This includes the status of the change (if a file was added, removed, modified, or renamed), and details of the change itself. For example, files with a `renamed` status have a `previous_filename` field showing the previous filename of the file, and files with a `modified` status have a `patch` field showing the changes made to the file.When calling this endpoint without any paging parameter (`per_page` or `page`), the returned list is limited to 250 commits, and the last commit in the list is the most recent of the entire comparison.**Working with large comparisons**To process a response with a large number of commits, use a query parameter (`per_page` or `page`) to paginate the results. When using pagination:- The list of changed files is only shown on the first page of results, but it includes all changed files for the entire comparison.- The results are returned in chronological order, but the last commit in the returned list may not be the most recent one in the entire set if there are more pages of results.For more information on working with pagination, see "[Using pagination in the REST API](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api)."**Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The `verification` object includes the following fields:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/commits/commits#compare-two-commits
func (m *ItemItemCompareWithBaseheadItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CommitComparisonable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "500": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "503": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCommitComparison503ErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCommitComparisonFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CommitComparisonable), nil
}
// ToGetRequestInformation compares two commits against one another. You can compare branches in the same repository, or you can compare branches that exist in different repositories within the same repository network, including fork branches. For more information about how to view a repository's network, see "[Understanding connections between repositories](https://docs.github.com/repositories/viewing-activity-and-data-for-your-repository/understanding-connections-between-repositories)."This endpoint is equivalent to running the `git log BASE...HEAD` command, but it returns commits in a different order. The `git log BASE...HEAD` command returns commits in reverse chronological order, whereas the API returns commits in chronological order. You can pass the appropriate [media type](https://docs.github.com/rest/overview/media-types/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.The API response includes details about the files that were changed between the two commits. This includes the status of the change (if a file was added, removed, modified, or renamed), and details of the change itself. For example, files with a `renamed` status have a `previous_filename` field showing the previous filename of the file, and files with a `modified` status have a `patch` field showing the changes made to the file.When calling this endpoint without any paging parameter (`per_page` or `page`), the returned list is limited to 250 commits, and the last commit in the list is the most recent of the entire comparison.**Working with large comparisons**To process a response with a large number of commits, use a query parameter (`per_page` or `page`) to paginate the results. When using pagination:- The list of changed files is only shown on the first page of results, but it includes all changed files for the entire comparison.- The results are returned in chronological order, but the last commit in the returned list may not be the most recent one in the entire set if there are more pages of results.For more information on working with pagination, see "[Using pagination in the REST API](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api)."**Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The `verification` object includes the following fields:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
func (m *ItemItemCompareWithBaseheadItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
