package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemWithEnablementItemRequestBuilder builds and executes requests for operations under \orgs\{org}\{security_product}\{enablement}
type ItemItemWithEnablementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemWithEnablementItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemWithEnablementItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemWithEnablementItemRequestBuilderInternal instantiates a new WithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    m := &ItemItemWithEnablementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/orgs/{org}/{security_product}/{enablement}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemWithEnablementItemRequestBuilder instantiates a new WithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemWithEnablementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables or disables the specified security feature for all repositories in an organization.To use this endpoint, you must be an organization owner or be member of a team with the security manager role.A token with the 'write:org' scope is also required.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reference/orgs#enable-or-disable-security-product-on-all-org-repos
func (m *ItemItemWithEnablementItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemWithEnablementItemRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation enables or disables the specified security feature for all repositories in an organization.To use this endpoint, you must be an organization owner or be member of a team with the security manager role.A token with the 'write:org' scope is also required.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
func (m *ItemItemWithEnablementItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemWithEnablementItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
