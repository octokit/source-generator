// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using System.Security.Cryptography;
using Microsoft.IdentityModel.Tokens;
using Microsoft.Kiota.Abstractions.Authentication;

namespace GitHub.Octokit.Client.Authentication;

/// <summary>
/// Used for GitHub App InstallationToken authentication
/// </summary>
public sealed class AppInstallationTokenProvider : IAccessTokenProvider
{
    private readonly string _sourceId;
    private readonly RSA _privateKey;
    private readonly string _installationId;
    private readonly IGitHubAppTokenProvider _gitHubAppTokenProvider;

    private string _accessToken = string.Empty;
    private SecurityTokenDescriptor? _tokenDescriptor;

    AllowedHostsValidator IAccessTokenProvider.AllowedHostsValidator => new();

    /// <summary>
    /// Constructor for AppInstallationTokenProvider using the clientId
    /// </summary>
    /// <param name="clientId">The client ID of the GitHub App.</param>
    /// <param name="privateKey">The private key of the GitHub App.</param>
    /// <param name="installationId">The installation ID of the GitHub App.</param>
    /// <param name="githubAppTokenProvider">The provider for creating and managing GitHub App tokens.</param>
    public AppInstallationTokenProvider(string clientId, RSA privateKey, string installationId, IGitHubAppTokenProvider githubAppTokenProvider)
    {
        _sourceId = clientId;
        _privateKey = privateKey;
        _installationId = installationId;
        _gitHubAppTokenProvider = githubAppTokenProvider;
    }

    /// <summary>
    /// Constructor for AppInstallationTokenProvider using the appId
    /// </summary>
    /// <param name="appId">The app ID of the GitHub App.</param>
    /// <param name="privateKey">The private key of the GitHub App.</param>
    /// <param name="installationId">The installation ID of the GitHub App.</param>
    /// <param name="githubAppTokenProvider">The provider for creating and managing GitHub App tokens.</param>
    public AppInstallationTokenProvider(int appId, RSA privateKey, string installationId, IGitHubAppTokenProvider githubAppTokenProvider)
    {
        _sourceId = appId.ToString();
        _privateKey = privateKey;
        _installationId = installationId;
        _gitHubAppTokenProvider = githubAppTokenProvider;
    }

    /// <summary>
    /// Get the authorization token
    /// </summary>
    /// <param name="requestUri">The URI of the request.</param>
    /// <param name="additionalAuthenticationContext">Additional context for authentication.</param>
    /// <param name="cancellationToken">A token to cancel the operation.</param>
    /// <returns>A task that represents the asynchronous operation. The task result contains the authorization token.</returns>
    public async Task<string> GetAuthorizationTokenAsync(Uri requestUri, Dictionary<string, object>? additionalAuthenticationContext = default, CancellationToken cancellationToken = default)
    {
        // If the token is empty, about to be expired, or has expired - get a new one
        if (string.IsNullOrEmpty(_accessToken) || (_tokenDescriptor != null && _tokenDescriptor.Expires < DateTime.UtcNow.AddMinutes(-1)))
        {
            var baseUrl = requestUri.GetLeftPart(UriPartial.Authority);

            _tokenDescriptor = _gitHubAppTokenProvider.CreateTokenDescriptor(_privateKey, _sourceId, DateTime.UtcNow);
            var jwt = _gitHubAppTokenProvider.CreateJsonWebToken(_tokenDescriptor);
            _accessToken = await _gitHubAppTokenProvider.GetGitHubAccessTokenAsync(baseUrl, jwt, _installationId);
        }

        return _accessToken;
    }
}
