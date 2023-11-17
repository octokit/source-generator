using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;


namespace GitHub.Client;

public static class OctokitRequestAdapter
{
    public static HttpClientRequestAdapter Create(IAuthenticationProvider authenticationProvider)
    {
      var octokitClientFactory = OctokitClientFactory.Create(); 
      var githubRequestAdapter = new HttpClientRequestAdapter(authenticationProvider, null, null, octokitClientFactory, null);
      return githubRequestAdapter;
    }
}
