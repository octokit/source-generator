// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using Microsoft.Extensions.Logging;

namespace GitHub.Octokit.SDK.Middleware;

internal static partial class Log
{
    [LoggerMessage(
        eventId: 0, LogLevel.Debug, """
            Primary rate limit (reset: {HeaderValue}) exceeded.        
                    Sleeping for {RetryAfterDurationSeconds} seconds before retrying.
            """
    )]
    public static partial void LogPrimaryRateLimitExceeded(
        this ILogger logger, string? headerValue, double retryAfterDurationSeconds);

    [LoggerMessage(
        eventId: 1, LogLevel.Information, """
            Rate limit information:
                    {XRateLimitLimitKey}: {XRateLimitLimitValue}, 
                    {XRateLimitUsedKey}: {XRateLimitUsedValue}, 
                    {XRateLimitResourceKey}: {XRateLimitResourceValue}
            """
    )]
    public static partial void LogRateLimitInformation(
        this ILogger logger, string XRateLimitLimitKey, string? XRateLimitLimitValue,
        string XRateLimitUsedKey, string? XRateLimitUsedValue,
        string XRateLimitResourceKey, string? XRateLimitResourceValue);

    [LoggerMessage(
        eventId: 2, LogLevel.Warning, """
            Abuse detection mechanism (secondary rate limit) triggered. 
                    Sleeping for {RetryAfterDurationSeconds} seconds before retrying.
            """
    )]
    public static partial void LogSecondaryRateLimitExceeded(
        this ILogger logger, double retryAfterDurationSeconds);

    [LoggerMessage(
    eventId: 3, LogLevel.Warning, """
            Could not parse a valid time to wait for rate limit reset (parsed {RetryAfterDurationSeconds} seconds).
                    Retrying request immediately.
            """
    )]
    public static partial void LogInvalidRateLimitResetTime(
        this ILogger logger, double retryAfterDurationSeconds);
}
