// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using Microsoft.Extensions.Logging;

internal static partial class Log
{
    [LoggerMessage(
    eventId: 5, LogLevel.Trace, """
            Appending GitHub API version HTTP header to request. {Key}: {Value}
            """
    )]
    public static partial void LogAppendingApiVersionHeader(
        this ILogger logger, string Key, string Value);
}
