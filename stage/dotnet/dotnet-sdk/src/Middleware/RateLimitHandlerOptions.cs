// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using System.Net;

namespace GitHub.Octokit.Client.Middleware;

/// <inheritdoc cref="IRateLimitHandlerOptions" />
public sealed class RateLimitHandlerOptions : IRateLimitHandlerOptions
{
    /// <summary>
    /// Gets the function that determines if the request is rate limited.
    /// The function should return the type of rate limit that is applied to the request.
    /// If the request is not rate limited, the function should return <see cref="RateLimitType.None"/>.
    /// </summary>
    public Func<HttpRequestMessage, HttpResponseMessage, RateLimitType> IsRateLimited => (request, response) =>
    {
        if (response.StatusCode is not HttpStatusCode.TooManyRequests and not HttpStatusCode.Forbidden)
        {
            return RateLimitType.None;
        }

        var retryAfter = response.Headers.RetryAfter;

        var rateLimitRemaining = response.Headers.TryGetValues("X-RateLimit-Remaining", out var rateLimitRemainingValues)
            ? rateLimitRemainingValues.FirstOrDefault()
            : null;

        return retryAfter is not null && rateLimitRemaining is not "0"
            ? RateLimitType.Secondary
            : rateLimitRemaining switch
            {
                "0" => RateLimitType.Primary,
                _ => RateLimitType.None
            };
    };
}
