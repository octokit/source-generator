// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using System.Security.Cryptography;
using Microsoft.IdentityModel.Tokens;

namespace GitHub.Octokit.Client.Authentication;

/// <summary>
/// This is an interface for the provider that creates a token descriptor, uses that
/// descriptor to create a JWT, and uses that JWT to create an access token
/// for authentication as a GitHub App installation.
/// </summary>
public interface IGitHubAppTokenProvider
{
    /// <summary>
    /// Creates a token descriptor we can use to sign a JWT as a GitHub App.
    /// </summary>
    /// <param name="privateKey">The private key belonging to the GitHub App.</param>
    /// <param name="sourceId">The issuer for the token descriptor claims.</param>
    /// <param name="issuedAt">The DateTime for the token descriptor to be issued at.</param>
    /// <returns>A token descriptor that can be used to create a JWT.</returns>
    public SecurityTokenDescriptor CreateTokenDescriptor(RSA privateKey, string sourceId, DateTime issuedAt);

    /// <summary>
    /// Creates a JWT using the given token descriptor. This JWT can be used to authenticate to the endpoints
    /// requiring the App to authenticate as itself. The list of endpoints can be found here:
    /// https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28
    /// </summary>
    /// <param name="tokenDescriptor">The token descriptor to use when creating the JWT.</param>
    /// <returns>A JWT string that can be used for authentication.</returns>
    public string CreateJsonWebToken(SecurityTokenDescriptor tokenDescriptor);

    /// <summary>
    /// Gets or renews a GitHub App installation access token. This token can be used to authenticate
    /// to all GitHub App installation endpoints:
    /// https://docs.github.com/en/rest/authentication/endpoints-available-for-github-app-installation-access-tokens?apiVersion=2022-11-28
    /// </summary>
    /// <param name="baseUrl">The base URL of the GitHub API.</param>
    /// <param name="jwt">The JSON Web Token for authentication.</param>
    /// <param name="installationId">The installation ID of the GitHub App.</param>
    /// <returns>An access token that can be used to authenticate to the GitHub API as an App installation.</returns>
    /// <exception cref="Exception">Thrown when the access token cannot be retrieved.</exception>
    public Task<string> GetGitHubAccessTokenAsync(string baseUrl, string jwt, string installationId);
}
