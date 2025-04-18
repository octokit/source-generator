// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using System.Net.Http.Headers;
using GitHub.Octokit.Client.Middleware.Options;
using Microsoft.Extensions.Logging;
using Microsoft.Kiota.Http.HttpClientLibrary.Extensions;

namespace GitHub.Octokit.Client.Middleware;

/// <summary>
/// A handler to add or update the User-Agent header in HTTP requests.
/// </summary>
/// <param name="logger">The logger to use for logging.</param>
/// <param name="userAgentHandlerOption">The user agent options to use for the handler.</param>
public sealed class UserAgentHandler(
    ILogger<UserAgentHandler> logger,
    UserAgentOptions? userAgentHandlerOption = null) : DelegatingHandler
{
    /// <summary>
    /// Sends an HTTP request to the inner handler to send to the server as an asynchronous operation.
    /// </summary>
    /// <param name="request">The HTTP request message to send to the server.</param>
    /// <param name="cancellationToken">A cancellation token to cancel the operation.</param>
    /// <returns>The task object representing the asynchronous operation.</returns>
    /// <exception cref="ArgumentNullException">Thrown when the request is null.</exception>
    protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        ArgumentNullException.ThrowIfNull(request);

        var options = request.GetRequestOption<UserAgentOptions>() ?? userAgentHandlerOption;

        if (!request.Headers.UserAgent.Any(x => options?.ProductName?.Equals(x.Product?.Name, StringComparison.OrdinalIgnoreCase) ?? false))
        {
            request.Headers.UserAgent.Add(new ProductInfoHeaderValue(options?.ProductName ?? string.Empty, options?.ProductVersion));

            logger.LogUserAgentHeaderAdded(options?.ProductName, options?.ProductVersion);
        }

        return base.SendAsync(request, cancellationToken);
    }
}
