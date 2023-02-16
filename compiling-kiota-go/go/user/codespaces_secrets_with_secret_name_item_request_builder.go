package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
)

// CodespacesSecretsWithSecret_nameItemRequestBuilder builds and executes requests for operations under \user\codespaces\secrets\{secret_name}
type CodespacesSecretsWithSecret_nameItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CodespacesSecretsWithSecret_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsWithSecret_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesSecretsWithSecret_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsWithSecret_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesSecretsWithSecret_nameItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsWithSecret_nameItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesSecretsWithSecret_nameItemRequestBuilderInternal instantiates a new WithSecret_nameItemRequestBuilder and sets the default values.
func NewCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsWithSecret_nameItemRequestBuilder) {
    m := &CodespacesSecretsWithSecret_nameItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/codespaces/secrets/{secret_name}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCodespacesSecretsWithSecret_nameItemRequestBuilder instantiates a new WithSecret_nameItemRequestBuilder and sets the default values.
func NewCodespacesSecretsWithSecret_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a secret from a user's codespaces using the secret name. Deleting the secret will remove access from all codespaces that were allowed to access the secret.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have write access to the `codespaces_user_secrets` user permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#delete-a-secret-for-the-authenticated-user
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get gets a secret available to a user's codespaces without revealing its encrypted value.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#get-a-secret-for-the-authenticated-user
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderGetRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodespacesSecretable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateCodespacesSecretFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CodespacesSecretable), nil
}
// Put creates or updates a secret for a user's codespace with an encrypted value. Encrypt your secret using[LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages).You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must also have Codespaces access to use this endpoint.GitHub Apps must have write access to the `codespaces_user_secrets` user permission and `codespaces_secrets` repository permission on all referenced repositories to use this endpoint.#### Example encrypting a secret using Node.jsEncrypt your secret using the [libsodium-wrappers](https://www.npmjs.com/package/libsodium-wrappers) library.```const sodium = require('libsodium-wrappers')const secret = 'plain-text-secret' // replace with the secret you want to encryptconst key = 'base64-encoded-public-key' // replace with the Base64 encoded public key//Check if libsodium is ready and then proceed.sodium.ready.then(() => {  // Convert Secret & Base64 key to Uint8Array.  let binkey = sodium.from_base64(key, sodium.base64_variants.ORIGINAL)  let binsec = sodium.from_string(secret)  //Encrypt the secret using LibSodium  let encBytes = sodium.crypto_box_seal(binsec, binkey)  // Convert encrypted Uint8Array to Base64  let output = sodium.to_base64(encBytes, sodium.base64_variants.ORIGINAL)  console.log(output)});```#### Example encrypting a secret using PythonEncrypt your secret using [pynacl](https://pynacl.readthedocs.io/en/latest/public/#nacl-public-sealedbox) with Python 3.```from base64 import b64encodefrom nacl import encoding, publicdef encrypt(public_key: str, secret_value: str) -> str:  """Encrypt a Unicode string using the public key."""  public_key = public.PublicKey(public_key.encode("utf-8"), encoding.Base64Encoder())  sealed_box = public.SealedBox(public_key)  encrypted = sealed_box.encrypt(secret_value.encode("utf-8"))  return b64encode(encrypted).decode("utf-8")```#### Example encrypting a secret using C#Encrypt your secret using the [Sodium.Core](https://www.nuget.org/packages/Sodium.Core/) package.```var secretValue = System.Text.Encoding.UTF8.GetBytes("mySecret");var publicKey = Convert.FromBase64String("2Sg8iYjAxxmI2LvUXpJjkYrMxURPc8r+dB7TJyvvcCU=");var sealedPublicKeyBox = Sodium.SealedPublicKeyBox.Create(secretValue, publicKey);Console.WriteLine(Convert.ToBase64String(sealedPublicKeyBox));```#### Example encrypting a secret using RubyEncrypt your secret using the [rbnacl](https://github.com/RubyCrypto/rbnacl) gem.```rubyrequire "rbnacl"require "base64"key = Base64.decode64("+ZYvJDZMHUfBkJdyq5Zm9SKqeuBQ4sj+6sfjlH4CgG0=")public_key = RbNaCl::PublicKey.new(key)box = RbNaCl::Boxes::Sealed.from_public_key(public_key)encrypted_secret = box.encrypt("my_secret")# Print the base64 encoded secretputs Base64.strict_encode64(encrypted_secret)```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/codespaces#create-or-update-a-secret-for-the-authenticated-user
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) Put(ctx context.Context, body CodespacesSecretsItemWithSecret_namePutRequestBodyable, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderPutRequestConfiguration)(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.EmptyObjectable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateBasicErrorFromDiscriminatorValue,
        "422": ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateEmptyObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.EmptyObjectable), nil
}
// Repositories the repositories property
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) Repositories()(*CodespacesSecretsItemRepositoriesRequestBuilder) {
    return NewCodespacesSecretsItemRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RepositoriesById gets an item from the github.com/octokit/kiota/.user.codespaces.secrets.item.repositories.item collection
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) RepositoriesById(id string)(*CodespacesSecretsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["repository_id"] = id
    }
    return NewCodespacesSecretsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation deletes a secret from a user's codespaces using the secret name. Deleting the secret will remove access from all codespaces that were allowed to access the secret.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have write access to the `codespaces_user_secrets` user permission to use this endpoint.
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets a secret available to a user's codespaces without revealing its encrypted value.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation creates or updates a secret for a user's codespace with an encrypted value. Encrypt your secret using[LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages).You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must also have Codespaces access to use this endpoint.GitHub Apps must have write access to the `codespaces_user_secrets` user permission and `codespaces_secrets` repository permission on all referenced repositories to use this endpoint.#### Example encrypting a secret using Node.jsEncrypt your secret using the [libsodium-wrappers](https://www.npmjs.com/package/libsodium-wrappers) library.```const sodium = require('libsodium-wrappers')const secret = 'plain-text-secret' // replace with the secret you want to encryptconst key = 'base64-encoded-public-key' // replace with the Base64 encoded public key//Check if libsodium is ready and then proceed.sodium.ready.then(() => {  // Convert Secret & Base64 key to Uint8Array.  let binkey = sodium.from_base64(key, sodium.base64_variants.ORIGINAL)  let binsec = sodium.from_string(secret)  //Encrypt the secret using LibSodium  let encBytes = sodium.crypto_box_seal(binsec, binkey)  // Convert encrypted Uint8Array to Base64  let output = sodium.to_base64(encBytes, sodium.base64_variants.ORIGINAL)  console.log(output)});```#### Example encrypting a secret using PythonEncrypt your secret using [pynacl](https://pynacl.readthedocs.io/en/latest/public/#nacl-public-sealedbox) with Python 3.```from base64 import b64encodefrom nacl import encoding, publicdef encrypt(public_key: str, secret_value: str) -> str:  """Encrypt a Unicode string using the public key."""  public_key = public.PublicKey(public_key.encode("utf-8"), encoding.Base64Encoder())  sealed_box = public.SealedBox(public_key)  encrypted = sealed_box.encrypt(secret_value.encode("utf-8"))  return b64encode(encrypted).decode("utf-8")```#### Example encrypting a secret using C#Encrypt your secret using the [Sodium.Core](https://www.nuget.org/packages/Sodium.Core/) package.```var secretValue = System.Text.Encoding.UTF8.GetBytes("mySecret");var publicKey = Convert.FromBase64String("2Sg8iYjAxxmI2LvUXpJjkYrMxURPc8r+dB7TJyvvcCU=");var sealedPublicKeyBox = Sodium.SealedPublicKeyBox.Create(secretValue, publicKey);Console.WriteLine(Convert.ToBase64String(sealedPublicKeyBox));```#### Example encrypting a secret using RubyEncrypt your secret using the [rbnacl](https://github.com/RubyCrypto/rbnacl) gem.```rubyrequire "rbnacl"require "base64"key = Base64.decode64("+ZYvJDZMHUfBkJdyq5Zm9SKqeuBQ4sj+6sfjlH4CgG0=")public_key = RbNaCl::PublicKey.new(key)box = RbNaCl::Boxes::Sealed.from_public_key(public_key)encrypted_secret = box.encrypt("my_secret")# Print the base64 encoded secretputs Base64.strict_encode64(encrypted_secret)```
func (m *CodespacesSecretsWithSecret_nameItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body CodespacesSecretsItemWithSecret_namePutRequestBodyable, requestConfiguration *CodespacesSecretsWithSecret_nameItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
