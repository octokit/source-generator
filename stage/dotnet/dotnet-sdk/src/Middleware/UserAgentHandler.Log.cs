// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using Microsoft.Extensions.Logging;

internal static partial class Log
{
    [LoggerMessage(
        eventId: 4, LogLevel.Debug, """
            User-Agent header added: {ProductName}/{ProductVersion}
            """
    )]
    public static partial void LogUserAgentHeaderAdded(
        this ILogger logger, string? ProductName, string? ProductVersion);
}
