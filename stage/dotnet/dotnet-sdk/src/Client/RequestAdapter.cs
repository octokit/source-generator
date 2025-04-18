// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Logging.Abstractions;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub.Octokit.Client;

/// <summary>
/// Represents an adapter for making HTTP requests using HttpClient with additional configuration options.
/// </summary>
public sealed class RequestAdapter
{
    /// <summary>
    /// Initializes and returns a new instance of the <see cref="HttpClientRequestAdapter"/> class.
    /// </summary>
    /// <param name="authenticationProvider">The authentication provider to use for making requests.</param>
    /// <param name="client">Optional: A custom HttpClient. If not provided, a default client will be created.</param>
    /// <param name="loggerFactory">Optional: A logger factory for logging. If not provided, no logging will be done.</param>
    /// <returns>A new instance of the <see cref="HttpClientRequestAdapter"/> class.</returns>
    public static HttpClientRequestAdapter Create(
        IAuthenticationProvider authenticationProvider,
        HttpClient? client = null,
        ILoggerFactory? loggerFactory = null)
    {
        client ??= ClientFactory.Create(loggerFactory: loggerFactory ?? NullLoggerFactory.Instance);

        var gitHubRequestAdapter = new HttpClientRequestAdapter(
            authenticationProvider,
            parseNodeFactory: null, // Node Parser
            serializationWriterFactory: null, // Serializer
            client,
            observabilityOptions: null);

        return gitHubRequestAdapter;
    }
}

