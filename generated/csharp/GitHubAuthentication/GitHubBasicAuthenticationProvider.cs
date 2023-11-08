using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Authentication;

namespace GitHubAuthentication;

public class GitHubBasicAuthenticationProvider : IAuthenticationProvider
{
  private const string AuthorizationHeaderKey = "Authorization";

  public IAuthenticationProvider BasicAuthProvider {get; private set;}
	public string ClientId { get; set; }
  public string Token { get; set; }

  public GitHubBasicAuthenticationProvider(string clientId, string token)
	{
    if (string.IsNullOrEmpty(clientId))
      throw new ArgumentNullException(nameof(clientId));
    if (string.IsNullOrEmpty(token))
      throw new ArgumentNullException(nameof(token));

    ClientId = clientId;
    Token = token;

    BasicAuthProvider = this;
	}

  public async Task AuthenticateRequestAsync(RequestInformation request, Dictionary<string, object>? additionalAuthenticationContext = null, CancellationToken cancellationToken = default)
  {
    if(request == null) throw new ArgumentNullException(nameof(request));
     
    request.Headers.Add("User-Agent", "Octokit.Gen/1.0.0");
    request.Headers.Add("X-GitHub-Api-Version", "2022-11-28");

    if(!request.Headers.ContainsKey(AuthorizationHeaderKey))
    {
      request.Headers.Add(AuthorizationHeaderKey, Token);
    }
  }
}