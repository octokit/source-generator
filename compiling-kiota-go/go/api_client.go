package kiota

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
	i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
	i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
	ifc4403625ac6d65b037d51dac0b8450f83f84fa95b011a32a4849688ebba1847 "github.com/octokit/kiota/app"
	i2699aa35e37f92b7330ae1cb15ac48dc1b4acdba81a00fabef41c5ee04453f80 "github.com/octokit/kiota/applications"
	i8cbdd523d1c8011d919ededcefb726cf630133ef73727eb02c192154e7c30bb5 "github.com/octokit/kiota/appmanifests"
	i4b17bcc282c8f00bfe90e6ff0ae5cb9f1eeb3d4e5e709c8524666c2b329424b9 "github.com/octokit/kiota/apps"
	ief9e243f7e9fbe7559cbefe40021aeab548090fc2e7840776f5be6482725de66 "github.com/octokit/kiota/codes_of_conduct"
	icc58f28852f33dd1da7e5fc0888426a7731874fc43ab04dd7f80e71ad75b9898 "github.com/octokit/kiota/emojis"
	ic95d4c9533090124aeefb820210c3a6ace61f1343ad8a48701700a71b3db91fa "github.com/octokit/kiota/enterprises"
	i0cd4a2573531d16179febd158356149a840f5d9e36df409a11e0bb1701e94747 "github.com/octokit/kiota/events"
	i0920a03be85467934e7a725767beaee47c97d6a75a7784d8f268ecde3d674649 "github.com/octokit/kiota/feeds"
	ieb4dc2aa96ee49a9447136681151b34303d62a118bbefd1253f6a8e33ff83edd "github.com/octokit/kiota/gists"
	ic44494a4083245e5f55895de144524b0d56ad575f33c29a8126deb725fe5482d "github.com/octokit/kiota/gitignore"
	ia8ea8260b731a8ca93961c5626e612351c4806b47ca7bc69805a20816af983f0 "github.com/octokit/kiota/installation"
	i894e3977fd5044a6acf34294e75df471cf88b9180f5f2613d84ae80443a2d724 "github.com/octokit/kiota/issues"
	i46fb96c085345cc2c06802ef555a6db8cfe8813b2bd17279e9ea253c000c9c69 "github.com/octokit/kiota/licenses"
	i4eb5ed635f9174d505702c48759f0ad2b89f1ff464abea9eb08459ef5952ec39 "github.com/octokit/kiota/markdown"
	i75bb1278ffc7bc959bb48a59206386b6507e2f34032cac0402c0032b059fe38f "github.com/octokit/kiota/marketplace_listing"
	ic93d3f945177a95a27cd547a45ebe1af89266e16b168bfb4aafbdd44d091592e "github.com/octokit/kiota/meta"
	ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614 "github.com/octokit/kiota/models"
	i9c2b7b9ecd37d6f9b17a83c5375293b4a47201599946743b7a926c1662716e53 "github.com/octokit/kiota/networks"
	iafaf7e38ec55d5234b841bc37ca7ab382b1201d69b893229de6b4c28bbce5997 "github.com/octokit/kiota/notifications"
	i693758f02a1fe111bc9d7e9808b2af65875dc48a60c4966d97bd2f337ca209f6 "github.com/octokit/kiota/octocat"
	ia06ee45b2593060af9c83898c26eef6da34275abca09940b5b41189cfbcba9f6 "github.com/octokit/kiota/organizations"
	if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151 "github.com/octokit/kiota/orgs"
	ia8baae919ed4e66c225e11c3e720c6da496812cdc045572d9849a1dbb5e59968 "github.com/octokit/kiota/projects"
	i59bd8821966f2115a63e1175391dc02fcc43d18d08459578e7eabd1cefe27a71 "github.com/octokit/kiota/rate_limit"
	i8eafed0d53cfe43f24065757459c8f79db39bb35e89a337dca4b137b91e50ffb "github.com/octokit/kiota/repos"
	ic6eef621790395e913c75097cf29d29c556901f507b531936f0a20d7108a7761 "github.com/octokit/kiota/repositories"
	i18c145efc2c69ba291165324ac9e680d45c2d9505b15d2f89944a9e0fb024e4b "github.com/octokit/kiota/search"
	ib5d19adce871fea292bd0c3769fd56df3e2690b21bb4b2fdf75554348d7f08e1 "github.com/octokit/kiota/teams"
	i3bab0a19bddffbc7922c6730f917bf1fb5a86e8fd5284cdb6242868938480ada "github.com/octokit/kiota/user"
	ib030f3063fc29ad31e8f3e9bc376f7bc95fbefeb1b6650bfcda0c013b5100f4c "github.com/octokit/kiota/users"
	i2cddf8cacd082fe848a95a943df1426a17f1924e5aa0e9abe2b0fc5e665fe614 "github.com/octokit/kiota/versions"
	ibcbb38b39827be5357edfe958c78d1a47e767f1c31635d5b9abe43243cff0d42 "github.com/octokit/kiota/zen"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ApiClientApiClientApiClientGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApiClientApiClientApiClientGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// App the app property
func (m *ApiClient) App() *ifc4403625ac6d65b037d51dac0b8450f83f84fa95b011a32a4849688ebba1847.AppRequestBuilder {
	return ifc4403625ac6d65b037d51dac0b8450f83f84fa95b011a32a4849688ebba1847.NewAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Applications the applications property
func (m *ApiClient) Applications() *i2699aa35e37f92b7330ae1cb15ac48dc1b4acdba81a00fabef41c5ee04453f80.ApplicationsRequestBuilder {
	return i2699aa35e37f92b7330ae1cb15ac48dc1b4acdba81a00fabef41c5ee04453f80.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ApplicationsById gets an item from the github.com/octokit/kiota/.applications.item collection
func (m *ApiClient) ApplicationsById(id string) *i2699aa35e37f92b7330ae1cb15ac48dc1b4acdba81a00fabef41c5ee04453f80.WithClient_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["client_id"] = id
	}
	return i2699aa35e37f92b7330ae1cb15ac48dc1b4acdba81a00fabef41c5ee04453f80.NewWithClient_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AppManifests the appManifests property
func (m *ApiClient) AppManifests() *i8cbdd523d1c8011d919ededcefb726cf630133ef73727eb02c192154e7c30bb5.AppManifestsRequestBuilder {
	return i8cbdd523d1c8011d919ededcefb726cf630133ef73727eb02c192154e7c30bb5.NewAppManifestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AppManifestsById gets an item from the github.com/octokit/kiota/.appManifests.item collection
func (m *ApiClient) AppManifestsById(id string) *i8cbdd523d1c8011d919ededcefb726cf630133ef73727eb02c192154e7c30bb5.WithCodeItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["code"] = id
	}
	return i8cbdd523d1c8011d919ededcefb726cf630133ef73727eb02c192154e7c30bb5.NewWithCodeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Apps the apps property
func (m *ApiClient) Apps() *i4b17bcc282c8f00bfe90e6ff0ae5cb9f1eeb3d4e5e709c8524666c2b329424b9.AppsRequestBuilder {
	return i4b17bcc282c8f00bfe90e6ff0ae5cb9f1eeb3d4e5e709c8524666c2b329424b9.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AppsById gets an item from the github.com/octokit/kiota/.apps.item collection
func (m *ApiClient) AppsById(id string) *i4b17bcc282c8f00bfe90e6ff0ae5cb9f1eeb3d4e5e709c8524666c2b329424b9.WithApp_slugItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["app_slug"] = id
	}
	return i4b17bcc282c8f00bfe90e6ff0ae5cb9f1eeb3d4e5e709c8524666c2b329424b9.NewWithApp_slugItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Codes_of_conduct the codes_of_conduct property
func (m *ApiClient) Codes_of_conduct() *ief9e243f7e9fbe7559cbefe40021aeab548090fc2e7840776f5be6482725de66.Codes_of_conductRequestBuilder {
	return ief9e243f7e9fbe7559cbefe40021aeab548090fc2e7840776f5be6482725de66.NewCodes_of_conductRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Codes_of_conductById gets an item from the github.com/octokit/kiota/.codes_of_conduct.item collection
func (m *ApiClient) Codes_of_conductById(id string) *ief9e243f7e9fbe7559cbefe40021aeab548090fc2e7840776f5be6482725de66.WithKeyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["key"] = id
	}
	return ief9e243f7e9fbe7559cbefe40021aeab548090fc2e7840776f5be6482725de66.NewWithKeyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ApiClient {
	m := &ApiClient{}
	m.pathParameters = make(map[string]string)
	m.urlTemplate = "{+baseurl}"
	m.requestAdapter = requestAdapter
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory()
	})
	if m.requestAdapter.GetBaseUrl() == "" {
		m.requestAdapter.SetBaseUrl("https://api.github.com")
	}
	m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
	return m
}

// Emojis the emojis property
func (m *ApiClient) Emojis() *icc58f28852f33dd1da7e5fc0888426a7731874fc43ab04dd7f80e71ad75b9898.EmojisRequestBuilder {
	return icc58f28852f33dd1da7e5fc0888426a7731874fc43ab04dd7f80e71ad75b9898.NewEmojisRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Enterprises the enterprises property
func (m *ApiClient) Enterprises() *ic95d4c9533090124aeefb820210c3a6ace61f1343ad8a48701700a71b3db91fa.EnterprisesRequestBuilder {
	return ic95d4c9533090124aeefb820210c3a6ace61f1343ad8a48701700a71b3db91fa.NewEnterprisesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// EnterprisesById gets an item from the github.com/octokit/kiota/.enterprises.item collection
func (m *ApiClient) EnterprisesById(id string) *ic95d4c9533090124aeefb820210c3a6ace61f1343ad8a48701700a71b3db91fa.WithEnterpriseItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["enterprise"] = id
	}
	return ic95d4c9533090124aeefb820210c3a6ace61f1343ad8a48701700a71b3db91fa.NewWithEnterpriseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Events the events property
func (m *ApiClient) Events() *i0cd4a2573531d16179febd158356149a840f5d9e36df409a11e0bb1701e94747.EventsRequestBuilder {
	return i0cd4a2573531d16179febd158356149a840f5d9e36df409a11e0bb1701e94747.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Feeds the feeds property
func (m *ApiClient) Feeds() *i0920a03be85467934e7a725767beaee47c97d6a75a7784d8f268ecde3d674649.FeedsRequestBuilder {
	return i0920a03be85467934e7a725767beaee47c97d6a75a7784d8f268ecde3d674649.NewFeedsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Get get Hypermedia links to resources accessible in GitHub's REST API
// [API method documentation]
//
// [API method documentation]: https://docs.github.com/rest/overview/resources-in-the-rest-api#root-endpoint
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration) (ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Rootable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.requestAdapter.Send(ctx, requestInfo, ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.CreateRootFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ib19f1474616f80bb588b8a50bea93468536b3ba5d3dffa1c19a10387bd034614.Rootable), nil
}

// Gists the gists property
func (m *ApiClient) Gists() *ieb4dc2aa96ee49a9447136681151b34303d62a118bbefd1253f6a8e33ff83edd.GistsRequestBuilder {
	return ieb4dc2aa96ee49a9447136681151b34303d62a118bbefd1253f6a8e33ff83edd.NewGistsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GistsById gets an item from the github.com/octokit/kiota/.gists.item collection
func (m *ApiClient) GistsById(id string) *ieb4dc2aa96ee49a9447136681151b34303d62a118bbefd1253f6a8e33ff83edd.WithGist_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["gist_id"] = id
	}
	return ieb4dc2aa96ee49a9447136681151b34303d62a118bbefd1253f6a8e33ff83edd.NewWithGist_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Gitignore the gitignore property
func (m *ApiClient) Gitignore() *ic44494a4083245e5f55895de144524b0d56ad575f33c29a8126deb725fe5482d.GitignoreRequestBuilder {
	return ic44494a4083245e5f55895de144524b0d56ad575f33c29a8126deb725fe5482d.NewGitignoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Installation the installation property
func (m *ApiClient) Installation() *ia8ea8260b731a8ca93961c5626e612351c4806b47ca7bc69805a20816af983f0.InstallationRequestBuilder {
	return ia8ea8260b731a8ca93961c5626e612351c4806b47ca7bc69805a20816af983f0.NewInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Issues the issues property
func (m *ApiClient) Issues() *i894e3977fd5044a6acf34294e75df471cf88b9180f5f2613d84ae80443a2d724.IssuesRequestBuilder {
	return i894e3977fd5044a6acf34294e75df471cf88b9180f5f2613d84ae80443a2d724.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Licenses the licenses property
func (m *ApiClient) Licenses() *i46fb96c085345cc2c06802ef555a6db8cfe8813b2bd17279e9ea253c000c9c69.LicensesRequestBuilder {
	return i46fb96c085345cc2c06802ef555a6db8cfe8813b2bd17279e9ea253c000c9c69.NewLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// LicensesById gets an item from the github.com/octokit/kiota/.licenses.item collection
func (m *ApiClient) LicensesById(id string) *i46fb96c085345cc2c06802ef555a6db8cfe8813b2bd17279e9ea253c000c9c69.WithLicenseItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["license"] = id
	}
	return i46fb96c085345cc2c06802ef555a6db8cfe8813b2bd17279e9ea253c000c9c69.NewWithLicenseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Markdown the markdown property
func (m *ApiClient) Markdown() *i4eb5ed635f9174d505702c48759f0ad2b89f1ff464abea9eb08459ef5952ec39.MarkdownRequestBuilder {
	return i4eb5ed635f9174d505702c48759f0ad2b89f1ff464abea9eb08459ef5952ec39.NewMarkdownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Marketplace_listing the marketplace_listing property
func (m *ApiClient) Marketplace_listing() *i75bb1278ffc7bc959bb48a59206386b6507e2f34032cac0402c0032b059fe38f.Marketplace_listingRequestBuilder {
	return i75bb1278ffc7bc959bb48a59206386b6507e2f34032cac0402c0032b059fe38f.NewMarketplace_listingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Meta the meta property
func (m *ApiClient) Meta() *ic93d3f945177a95a27cd547a45ebe1af89266e16b168bfb4aafbdd44d091592e.MetaRequestBuilder {
	return ic93d3f945177a95a27cd547a45ebe1af89266e16b168bfb4aafbdd44d091592e.NewMetaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Networks the networks property
func (m *ApiClient) Networks() *i9c2b7b9ecd37d6f9b17a83c5375293b4a47201599946743b7a926c1662716e53.NetworksRequestBuilder {
	return i9c2b7b9ecd37d6f9b17a83c5375293b4a47201599946743b7a926c1662716e53.NewNetworksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// NetworksById gets an item from the github.com/octokit/kiota/.networks.item collection
func (m *ApiClient) NetworksById(id string) *i9c2b7b9ecd37d6f9b17a83c5375293b4a47201599946743b7a926c1662716e53.WithOwnerItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["owner"] = id
	}
	return i9c2b7b9ecd37d6f9b17a83c5375293b4a47201599946743b7a926c1662716e53.NewWithOwnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Notifications the notifications property
func (m *ApiClient) Notifications() *iafaf7e38ec55d5234b841bc37ca7ab382b1201d69b893229de6b4c28bbce5997.NotificationsRequestBuilder {
	return iafaf7e38ec55d5234b841bc37ca7ab382b1201d69b893229de6b4c28bbce5997.NewNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Octocat the octocat property
func (m *ApiClient) Octocat() *i693758f02a1fe111bc9d7e9808b2af65875dc48a60c4966d97bd2f337ca209f6.OctocatRequestBuilder {
	return i693758f02a1fe111bc9d7e9808b2af65875dc48a60c4966d97bd2f337ca209f6.NewOctocatRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Organizations the organizations property
func (m *ApiClient) Organizations() *ia06ee45b2593060af9c83898c26eef6da34275abca09940b5b41189cfbcba9f6.OrganizationsRequestBuilder {
	return ia06ee45b2593060af9c83898c26eef6da34275abca09940b5b41189cfbcba9f6.NewOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Orgs the orgs property
func (m *ApiClient) Orgs() *if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.OrgsRequestBuilder {
	return if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.NewOrgsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// OrgsById gets an item from the github.com/octokit/kiota/.orgs.item collection
func (m *ApiClient) OrgsById(id string) *if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.WithOrgItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["org"] = id
	}
	return if221a68ede066e0b6627fbd3d6fffab5e576fb3f5fae809c44e5290f4bc85151.NewWithOrgItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Projects the projects property
func (m *ApiClient) Projects() *ia8baae919ed4e66c225e11c3e720c6da496812cdc045572d9849a1dbb5e59968.ProjectsRequestBuilder {
	return ia8baae919ed4e66c225e11c3e720c6da496812cdc045572d9849a1dbb5e59968.NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ProjectsById gets an item from the github.com/octokit/kiota/.projects.item collection
func (m *ApiClient) ProjectsById(id string) *ia8baae919ed4e66c225e11c3e720c6da496812cdc045572d9849a1dbb5e59968.WithProject_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["project_id"] = id
	}
	return ia8baae919ed4e66c225e11c3e720c6da496812cdc045572d9849a1dbb5e59968.NewWithProject_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Rate_limit the rate_limit property
func (m *ApiClient) Rate_limit() *i59bd8821966f2115a63e1175391dc02fcc43d18d08459578e7eabd1cefe27a71.Rate_limitRequestBuilder {
	return i59bd8821966f2115a63e1175391dc02fcc43d18d08459578e7eabd1cefe27a71.NewRate_limitRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Repos the repos property
func (m *ApiClient) Repos() *i8eafed0d53cfe43f24065757459c8f79db39bb35e89a337dca4b137b91e50ffb.ReposRequestBuilder {
	return i8eafed0d53cfe43f24065757459c8f79db39bb35e89a337dca4b137b91e50ffb.NewReposRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Repositories the repositories property
func (m *ApiClient) Repositories() *ic6eef621790395e913c75097cf29d29c556901f507b531936f0a20d7108a7761.RepositoriesRequestBuilder {
	return ic6eef621790395e913c75097cf29d29c556901f507b531936f0a20d7108a7761.NewRepositoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// RepositoriesById gets an item from the github.com/octokit/kiota/.repositories.item collection
func (m *ApiClient) RepositoriesById(id string) *ic6eef621790395e913c75097cf29d29c556901f507b531936f0a20d7108a7761.WithRepository_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["repository_id"] = id
	}
	return ic6eef621790395e913c75097cf29d29c556901f507b531936f0a20d7108a7761.NewWithRepository_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Search the search property
func (m *ApiClient) Search() *i18c145efc2c69ba291165324ac9e680d45c2d9505b15d2f89944a9e0fb024e4b.SearchRequestBuilder {
	return i18c145efc2c69ba291165324ac9e680d45c2d9505b15d2f89944a9e0fb024e4b.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Teams the teams property
func (m *ApiClient) Teams() *ib5d19adce871fea292bd0c3769fd56df3e2690b21bb4b2fdf75554348d7f08e1.TeamsRequestBuilder {
	return ib5d19adce871fea292bd0c3769fd56df3e2690b21bb4b2fdf75554348d7f08e1.NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TeamsById gets an item from the github.com/octokit/kiota/.teams.item collection
func (m *ApiClient) TeamsById(id string) *ib5d19adce871fea292bd0c3769fd56df3e2690b21bb4b2fdf75554348d7f08e1.WithTeam_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["team_id"] = id
	}
	return ib5d19adce871fea292bd0c3769fd56df3e2690b21bb4b2fdf75554348d7f08e1.NewWithTeam_ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// User the user property
func (m *ApiClient) User() *i3bab0a19bddffbc7922c6730f917bf1fb5a86e8fd5284cdb6242868938480ada.UserRequestBuilder {
	return i3bab0a19bddffbc7922c6730f917bf1fb5a86e8fd5284cdb6242868938480ada.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Users the users property
func (m *ApiClient) Users() *ib030f3063fc29ad31e8f3e9bc376f7bc95fbefeb1b6650bfcda0c013b5100f4c.UsersRequestBuilder {
	return ib030f3063fc29ad31e8f3e9bc376f7bc95fbefeb1b6650bfcda0c013b5100f4c.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// UsersById gets an item from the github.com/octokit/kiota/.users.item collection
func (m *ApiClient) UsersById(id string) *ib030f3063fc29ad31e8f3e9bc376f7bc95fbefeb1b6650bfcda0c013b5100f4c.WithUsernameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["username"] = id
	}
	return ib030f3063fc29ad31e8f3e9bc376f7bc95fbefeb1b6650bfcda0c013b5100f4c.NewWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Versions the versions property
func (m *ApiClient) Versions() *i2cddf8cacd082fe848a95a943df1426a17f1924e5aa0e9abe2b0fc5e665fe614.VersionsRequestBuilder {
	return i2cddf8cacd082fe848a95a943df1426a17f1924e5aa0e9abe2b0fc5e665fe614.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Zen the zen property
func (m *ApiClient) Zen() *ibcbb38b39827be5357edfe958c78d1a47e767f1c31635d5b9abe43243cff0d42.ZenRequestBuilder {
	return ibcbb38b39827be5357edfe958c78d1a47e767f1c31635d5b9abe43243cff0d42.NewZenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
