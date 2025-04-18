// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using GitHub.Octokit.Client.Authentication;
using Microsoft.Kiota.Abstractions.Authentication;
using Moq;
using Xunit;

namespace Tests.Client.Authentication;

public class AppInstallationAuthProviderTests
{
    [Fact]
    public void Constructor_InitializesTokenProvider()
    {
        var tokenProviderMock = new Mock<IAccessTokenProvider>();
        var authProvider = new AppInstallationAuthProvider(tokenProviderMock.Object);

        Assert.NotNull(authProvider.AccessTokenProvider);
        Assert.Same(tokenProviderMock.Object, authProvider.AccessTokenProvider);
    }
}
