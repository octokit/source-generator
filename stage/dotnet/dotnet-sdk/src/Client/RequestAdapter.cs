// Copyright (c) GitHub 2023-2024 - Licensed as MIT.
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub.Octokit.Client;

/// <summary>
/// Represents an adapter for making HTTP requests using HttpClient with additional configuration options.
/// </summary>
public class RequestAdapter
{
    private readonly HttpClientRequestAdapter _adapter;

    /// <summary>
    /// Initializes an instance of the <see cref="HttpClientRequestAdapter"/>   .
    /// Allows for support for the Builder pattern.
    /// </summary>
    /// <param name="authenticationProvider">The authentication provider to use for making requests.</param>
    /// <param name="clientFactory">Optional: A custom HttpClient factory. If not provided, a default client will be created.</param>
    public RequestAdapter(IAuthenticationProvider authenticationProvider, HttpClient? clientFactory = null)
    {
        _adapter = buildRequestAdapter(authenticationProvider, clientFactory);
    }

    /// <summary>
    /// Initializes and returns a new instance of the <see cref="HttpClientRequestAdapter"/> class.
    /// </summary>
    /// <param name="authenticationProvider">The authentication provider to use for making requests.</param>
    /// <param name="clientFactory">Optional: A custom HttpClient factory. If not provided, a default client will be created.</param>
    /// <returns>A new instance of the <see cref="HttpClientRequestAdapter"/> class.</returns>

    public static HttpClientRequestAdapter Create(IAuthenticationProvider authenticationProvider, HttpClient? clientFactory = null)
    {
        return buildRequestAdapter(authenticationProvider, clientFactory);
    }

    private static HttpClientRequestAdapter buildRequestAdapter(IAuthenticationProvider authenticationProvider, HttpClient? clientFactory = null)
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

    /// <summary>
    /// Sets the base URL for the HTTP requests.
    /// </summary>
    /// <param name="baseURL">The base URL to set.</param>
    /// <returns>The current instance of <see cref="RequestAdapter"/>, allowing method chaining.</returns>
    public RequestAdapter WithBaseUrl(string baseURL)
    {
        _adapter.BaseUrl = baseURL;
        return this;
    }

    /// <summary>
    /// Builds and returns the configured <see cref="HttpClientRequestAdapter"/>.
    /// </summary>
    /// <returns>The configured <see cref="HttpClientRequestAdapter"/>.</returns>
    public HttpClientRequestAdapter Build()
    {
        return _adapter;
    }
}