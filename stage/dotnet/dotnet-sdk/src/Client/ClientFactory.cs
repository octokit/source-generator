// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using System.Net;
using GitHub.Octokit.Client.Middleware;
using GitHub.Octokit.Client.Middleware.Options;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Logging.Abstractions;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub.Octokit.Client;

/// <summary>
/// Represents a client factory for creating <see cref="HttpClient"/>.
/// </summary>
public sealed class ClientFactory
{
    private readonly ILoggerFactory _loggerFactory;
    private readonly HttpMessageHandler? _finalHandler;
    private readonly Lazy<List<DelegatingHandler>>? _handlers;

    private TimeSpan? _requestTimeout;
    private string? _baseUrl;
    private IAuthenticationProvider? _authenticationProvider;

    /// <summary>
    /// Initializes a new instance of the <see cref="ClientFactory"/> class.
    /// </summary>
    /// <param name="loggerFactory">The logger factory used to create <see cref="ILogger{TCategoryName}"/> instances for the handlers.</param>
    /// <param name="finalHandler">The final <see cref="HttpMessageHandler"/> in the chain.</param>
    public ClientFactory(ILoggerFactory loggerFactory, HttpMessageHandler? finalHandler = null)
    {
        _loggerFactory = loggerFactory;
        _finalHandler = finalHandler;
        _handlers = new(() =>
        [
            new APIVersionHandler(_loggerFactory.CreateLogger<APIVersionHandler>()),
            new UserAgentHandler(_loggerFactory.CreateLogger<UserAgentHandler>()),
            new RateLimitHandler(_loggerFactory.CreateLogger<RateLimitHandler>()),
        ]);
    }

    /// <summary>
    /// Creates an <see cref="HttpClient"/> instance with the specified <see cref="HttpMessageHandler"/>.
    /// If no <see cref="HttpMessageHandler"/> is provided, a default one will be used.
    /// </summary>
    /// <param name="finalHandler">The final <see cref="HttpMessageHandler"/> in the chain.</param>
    /// <param name="loggerFactory">The logger factory used to create <see cref="ILogger{TCategoryName}"/> instances for the handlers.</param>
    /// <returns>An instance of <see cref="HttpClient"/>.</returns>
    public static HttpClient Create(HttpMessageHandler? finalHandler = null, ILoggerFactory? loggerFactory = null)
    {
        var clientFactory = new ClientFactory(
            loggerFactory: loggerFactory ?? NullLoggerFactory.Instance,
            finalHandler: finalHandler);

        var defaultHandlers = clientFactory?._handlers?.Value ?? [];

        var handler = ChainHandlersCollectionAndGetFirstLink(
            finalHandler: finalHandler ?? GetDefaultHttpMessageHandler(),
            handlers: [.. defaultHandlers]);

        return handler is not null ? new HttpClient(handler) : new HttpClient();
    }

    /// <summary>
    /// Sets the user agent for the <see cref="HttpClient"/>.
    /// </summary>
    /// <param name="productName">The product name.</param>
    /// <param name="productVersion">The product version.</param>
    /// <returns>The current instance of <see cref="ClientFactory"/>.</returns>
    public ClientFactory WithUserAgent(string productName, string productVersion)
    {
        AddOrCreateHandler(new UserAgentHandler(_loggerFactory.CreateLogger<UserAgentHandler>(), new UserAgentOptions
        {
            ProductName = productName,
            ProductVersion = productVersion
        }));

        return this;
    }

    /// <summary>
    /// Sets the request timeout for the <see cref="HttpClient"/>.
    /// </summary>
    /// <param name="timeSpan">The request timeout.</param>
    /// <returns>The current instance of <see cref="ClientFactory"/>.</returns>
    public ClientFactory WithRequestTimeout(TimeSpan timeSpan)
    {
        _requestTimeout = timeSpan;

        return this;
    }

    /// <summary>
    /// Sets the base URL for the <see cref="HttpClient"/>.
    /// </summary>
    /// <param name="baseUrl">The base URL.</param>
    /// <returns>The current instance of <see cref="ClientFactory"/>.</returns>
    public ClientFactory WithBaseUrl(string baseUrl)
    {
        _baseUrl = baseUrl;

        return this;
    }

    /// <summary>
    /// Sets the authentication provider for the <see cref="HttpClient"/>.
    /// </summary>
    /// <param name="authenticationProvider">The authentication provider.</param>
    /// <returns>The current instance of <see cref="ClientFactory"/>.</returns>
    public ClientFactory WithAuthenticationProvider(IAuthenticationProvider authenticationProvider)
    {
        _authenticationProvider = authenticationProvider;

        return this;
    }

    /// <summary>
    /// Builds the <see cref="HttpClientRequestAdapter"/> with the configured settings.
    /// </summary>
    /// <returns>An instance of <see cref="HttpClientRequestAdapter"/>.</returns>
    /// <exception cref="ArgumentNullException">Thrown when the authentication provider is null.</exception>
    public HttpClientRequestAdapter Build()
    {
        ArgumentNullException.ThrowIfNull(_authenticationProvider);

        var httpClient = new HttpClient();

        var handler = ChainHandlersCollectionAndGetFirstLink(
            finalHandler: _finalHandler ?? GetDefaultHttpMessageHandler(),
            handlers: [.. CreateDefaultHandlers(_loggerFactory)]);

        if (handler is not null)
        {
            httpClient = new HttpClient(handler);
        }

        if (_requestTimeout.HasValue)
        {
            httpClient.Timeout = _requestTimeout.Value;
        }

        if (!string.IsNullOrWhiteSpace(_baseUrl))
        {
            httpClient.BaseAddress = new Uri(_baseUrl);
        }

        return RequestAdapter.Create(_authenticationProvider, httpClient);
    }

    /// <summary>
    /// Creates a list of default delegating handlers for the Octokit client.
    /// </summary>
    /// <param name="loggerFactory">The logger factory used to create <see cref="ILogger{TCategoryName}"/> instances for the handlers.</param>
    /// <returns>A list of default delegating handlers.</returns>
    public static IList<DelegatingHandler> CreateDefaultHandlers(ILoggerFactory loggerFactory)
    {
        var clientFactory = new ClientFactory(
            loggerFactory: loggerFactory ?? NullLoggerFactory.Instance);

        var defaultHandlers = clientFactory?._handlers?.Value ?? [];

        var kiotaDefaultHandlers = KiotaClientFactory.CreateDefaultHandlers();

        return [.. kiotaDefaultHandlers, .. defaultHandlers];
    }

    /// <summary>
    /// Creates a list of default delegating handlers for the Octokit client.
    /// </summary>
    /// <returns>A list of default delegating handlers.</returns>
    public static IList<DelegatingHandler> CreateDefaultHandlers() => CreateDefaultHandlers(NullLoggerFactory.Instance);

    /// <summary>
    /// Chains a collection of <see cref="DelegatingHandler"/> instances and returns the first link in the chain.
    /// </summary>
    /// <param name="finalHandler">The final <see cref="HttpMessageHandler"/> in the chain.</param>
    /// <param name="handlers">The collection of <see cref="DelegatingHandler"/> instances to be chained.</param>
    /// <returns>The first link in the chain of <see cref="DelegatingHandler"/> instances.</returns>
    public static DelegatingHandler? ChainHandlersCollectionAndGetFirstLink(HttpMessageHandler? finalHandler, params DelegatingHandler[] handlers)
    {
        if (handlers is null or { Length: 0 })
        {
            return default;
        }

        for (var i = 0; i < handlers.Length; i++)
        {
            var handler = handlers[i];
            var previousItemIndex = i - 1;

            if (previousItemIndex >= 0)
            {
                var previousHandler = handlers[previousItemIndex];
                previousHandler.InnerHandler = handler;
            }
        }

        if (finalHandler is not null)
        {
            handlers[^1].InnerHandler = finalHandler;
        }

        return handlers.First();
    }

    /// <summary>
    /// Chains a collection of <see cref="DelegatingHandler"/> instances and returns the first link in the chain.
    /// </summary>
    /// <param name="handlers">The collection of <see cref="DelegatingHandler"/> instances to chain.</param>
    /// <returns>The first link in the chain of <see cref="DelegatingHandler"/> instances.</returns>
    public static DelegatingHandler? ChainHandlersCollectionAndGetFirstLink(
        params DelegatingHandler[] handlers) =>
        ChainHandlersCollectionAndGetFirstLink(null, handlers);

    /// <summary>
    /// Gets the default <see cref="HttpMessageHandler"/> for the Octokit client.
    /// </summary>
    /// <param name="proxy">The optional web proxy to use.</param>
    /// <returns>The default HTTP message handler.</returns>
    public static HttpMessageHandler GetDefaultHttpMessageHandler(IWebProxy? proxy = null) =>
        new HttpClientHandler
        {
            Proxy = proxy,
            AllowAutoRedirect = false
        };

    /// <summary>
    /// In support of the constructor approach to building a client factory, this method allows for adding or updating
    /// a handler in the list of handlers.
    /// The final result of the list of handlers will be processed in the Build() method.
    /// </summary>
    /// <typeparam name="THandler">The type of the handler.</typeparam>
    /// <param name="handler">The handler to add or update.</param>
    private void AddOrCreateHandler<THandler>(THandler handler) where THandler : DelegatingHandler
    {
        // Find the index of the handler that matches the specified type
        var handlers = _handlers is { IsValueCreated: true }
            ? _handlers!.Value
            : null;

        if (handlers is null)
        {
            return;
        }

        var index = handlers.FindIndex(h => h is THandler);

        // If the handler is found, replace it with the new handler otherwise add the new handler to the list
        if (index >= 0)
        {
            handlers[index] = handler;
        }
        else
        {
            handlers.Add(handler);
        }
    }
}
