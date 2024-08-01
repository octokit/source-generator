// Copyright (c) GitHub 2023-2024 - Licensed as MIT.

using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub.Octokit.Client;

public static class RequestAdapter
{
    /// <summary>
    /// Represents an adapter for making HTTP requests using HttpClient.
    /// </summary>
    /// TODO: Implement the missing props and methods
    public static HttpClientRequestAdapter Create(IAuthenticationProvider authenticationProvider, HttpClient? clientFactory = null)
    {
        clientFactory ??= ClientFactory.Create();

        var gitHubRequestAdapter =
          new HttpClientRequestAdapter(
            authenticationProvider,
            null, // Node Parser
            null, // Serializer
            clientFactory,
            null);

        return gitHubRequestAdapter;
    }
    public static HttpClientRequestAdapter Create(IAuthenticationProvider authenticationProvider, string baseUrl, HttpClient? clientFactory = null)
    {
        if (string.IsNullOrEmpty(baseUrl))
        {
            throw new ArgumentException("Base URL cannot be null or empty.", nameof(baseUrl));
        }

        clientFactory ??= ClientFactory.Create();

        var gitHubRequestAdapter = new HttpClientRequestAdapter(
            authenticationProvider,
            null, // Node Parser
            null, // Serializer
            clientFactory,
            null)
        {
            BaseUrl = baseUrl
        };

        return gitHubRequestAdapter;
    }
}
