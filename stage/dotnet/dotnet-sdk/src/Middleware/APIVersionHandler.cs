// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using GitHub.Octokit.Client.Middleware.Options;
using Microsoft.Extensions.Logging;
using Microsoft.Kiota.Http.HttpClientLibrary.Extensions;

namespace GitHub.Octokit.Client.Middleware;

/// <summary>
/// Represents a handler that adds the API version header to outgoing HTTP requests.
/// </summary>
/// <param name="logger">The logger to use for logging.</param>
/// <param name="apiVersionOptions">The API version options to use for the handler.</param>
public sealed class APIVersionHandler(
    ILogger<APIVersionHandler> logger,
    APIVersionOptions? apiVersionOptions = null) : DelegatingHandler
{
    private const string ApiVersionHeaderKey = "X-GitHub-Api-Version";

    private readonly APIVersionOptions _apiVersionOptions = apiVersionOptions ?? new APIVersionOptions();

    /// <summary>
    /// Sends an HTTP request to the inner handler to send to the server as an asynchronous operation.
    /// </summary>
    /// <param name="request">The HTTP request message to send.</param>
    /// <param name="cancellationToken">A cancellation token to cancel the operation.</param>
    /// <returns>A task that represents the asynchronous operation. The task result contains the HTTP response message.</returns>
    protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        ArgumentNullException.ThrowIfNull(request);

        var apiVersionHandlerOption = request.GetRequestOption<APIVersionOptions>() ?? _apiVersionOptions;

        if (!request.Headers.Contains(ApiVersionHeaderKey) || !request.Headers.GetValues(ApiVersionHeaderKey).Any(x => APIVersionOptions.APIVersion.Equals(x, StringComparison.OrdinalIgnoreCase)))
        {
            request.Headers.Add(ApiVersionHeaderKey, APIVersionOptions.APIVersion);

            logger.LogAppendingApiVersionHeader(ApiVersionHeaderKey, APIVersionOptions.APIVersion);
        }

        return base.SendAsync(request, cancellationToken);
    }
}
