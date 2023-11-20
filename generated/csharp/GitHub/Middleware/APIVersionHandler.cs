using GitHub.Octokit.Client.Middleware.Options;
using Microsoft.Kiota.Http.HttpClientLibrary.Extensions;

namespace GitHub.Octokit.Client.Middleware;

/// <summary>
/// Represents a handler that adds the API version header to outgoing HTTP requests.
/// </summary>
class APIVersionHandler : DelegatingHandler
{
  private const string ApiVersionHeaderKey = "X-GitHub-Api-Version";
  private const string defaultAPIVersion = "2022-11-28";

  private readonly APIVersionOptions _apiVersionOptions;

  public APIVersionHandler(APIVersionOptions? apiVersionOptions = null)
  {
      _apiVersionOptions = apiVersionOptions ?? new APIVersionOptions();
  }

  protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
  {
    if (request == null){
      throw new ArgumentNullException(nameof(request));
    }

    var apiVersionHandlerOption = request.GetRequestOption<APIVersionOptions>() ?? _apiVersionOptions;

    if (!request.Headers.Contains("X-GitHub-Api-Version") || !request.Headers.GetValues("X-GitHub-Api-Version").Any(x => apiVersionHandlerOption.APIVersion.Equals(x, StringComparison.OrdinalIgnoreCase)))
    {
      request.Headers.Add("X-GitHub-Api-Version", apiVersionHandlerOption.APIVersion);
    }
    return base.SendAsync(request, cancellationToken);
  }
}
