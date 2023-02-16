package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// ItemItemGitCommitsWithCommit_shaItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\git\commits\{commit_sha}
type ItemItemGitCommitsWithCommit_shaItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemGitCommitsWithCommit_shaItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemGitCommitsWithCommit_shaItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemGitCommitsWithCommit_shaItemRequestBuilderInternal instantiates a new WithCommit_shaItemRequestBuilder and sets the default values.
func NewItemItemGitCommitsWithCommit_shaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitCommitsWithCommit_shaItemRequestBuilder) {
    m := &ItemItemGitCommitsWithCommit_shaItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/repos/{owner}/{repo}/git/commits/{commit_sha}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemGitCommitsWithCommit_shaItemRequestBuilder instantiates a new WithCommit_shaItemRequestBuilder and sets the default values.
func NewItemItemGitCommitsWithCommit_shaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitCommitsWithCommit_shaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitCommitsWithCommit_shaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a Git [commit object](https://git-scm.com/book/en/v1/Git-Internals-Git-Objects#Commit-Objects).**Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in the table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/git#get-a-commit
func (m *ItemItemGitCommitsWithCommit_shaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemGitCommitsWithCommit_shaItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.GitCommitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateGitCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.GitCommitable), nil
}
// ToGetRequestInformation gets a Git [commit object](https://git-scm.com/book/en/v1/Git-Internals-Git-Objects#Commit-Objects).**Signature verification object**The response will include a `verification` object that describes the result of verifying the commit's signature. The following fields are included in the `verification` object:| Name | Type | Description || ---- | ---- | ----------- || `verified` | `boolean` | Indicates whether GitHub considers the signature in this commit to be verified. || `reason` | `string` | The reason for verified value. Possible values and their meanings are enumerated in the table below. || `signature` | `string` | The signature that was extracted from the commit. || `payload` | `string` | The value that was signed. |These are the possible values for `reason` in the `verification` object:| Value | Description || ----- | ----------- || `expired_key` | The key that made the signature is expired. || `not_signing_key` | The "signing" flag is not among the usage flags in the GPG key that made the signature. || `gpgverify_error` | There was an error communicating with the signature verification service. || `gpgverify_unavailable` | The signature verification service is currently unavailable. || `unsigned` | The object does not include a signature. || `unknown_signature_type` | A non-PGP signature was found in the commit. || `no_user` | No user was associated with the `committer` email address in the commit. || `unverified_email` | The `committer` email address in the commit was associated with a user, but the email address is not verified on her/his account. || `bad_email` | The `committer` email address in the commit is not included in the identities of the PGP key that made the signature. || `unknown_key` | The key that made the signature has not been registered with any user's account. || `malformed_signature` | There was an error parsing the signature. || `invalid` | The signature could not be cryptographically verified using the key whose key-id was found in the signature. || `valid` | None of the above errors applied, so the signature is considered to be verified. |
func (m *ItemItemGitCommitsWithCommit_shaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemGitCommitsWithCommit_shaItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
