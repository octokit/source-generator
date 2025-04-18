// Copyright (c) GitHub 2023-2025 - Licensed as MIT.

using GitHub.Octokit.Client.Authentication;
using Microsoft.Kiota.Abstractions;
using Xunit;

namespace Tests.Client.Authentication;

public class TokenAuthenticationProviderTests
{
    private const string ValidToken = "validToken";

    private readonly TokenProvider _tokenProvider;
    private readonly TokenAuthProvider _authProvider;

    public TokenAuthenticationProviderTests()
    {
        _tokenProvider = new TokenProvider(ValidToken);
        _authProvider = new TokenAuthProvider(_tokenProvider);
    }

    [Fact]
    public void Constructor_ThrowsException_WhenTokenIsEmpty() => Assert.Throws<ArgumentException>(() => new TokenProvider(""));

    [Fact]
    public async Task AuthenticateRequestAsync_AddsAuthorizationHeader_WhenRequestIsValid()
    {
        var request = new RequestInformation(Method.GET, "https://localhost", new Dictionary<string, object>());
        await _authProvider.AuthenticateRequestAsync(request);
        var headerToken = request.Headers["Authorization"].FirstOrDefault();

        Assert.True(request.Headers.ContainsKey("Authorization"));
        Assert.Equal($"Bearer {ValidToken}", headerToken);
    }
}
