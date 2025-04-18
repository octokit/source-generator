// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

namespace GitHub.Octokit.Client.Middleware;

/// <summary>
/// Defines the options for the rate limit handler.
/// </summary>
public interface IRateLimitHandlerOptions
{
    /// <summary>
    /// Gets the function that determines if the request is rate limited.
    /// </summary>
    Func<HttpRequestMessage, HttpResponseMessage, RateLimitType> IsRateLimited { get; }
}
