using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;
using GitHub.Octokit.Client.Middleware;


namespace GitHub.Client;

public static class OctokitRequestAdapter
{
    /// <summary>
    /// Represents an adapter for making HTTP requests using HttpClient.
    /// </summary>
    public static HttpClientRequestAdapter Create(IAuthenticationProvider authenticationProvider)
    {
      var octokitClientFactory = OctokitClientFactory.Create(); 
      var githubRequestAdapter = new HttpClientRequestAdapter(authenticationProvider, null, null, octokitClientFactory, null);
      return githubRequestAdapter;
    }
}
