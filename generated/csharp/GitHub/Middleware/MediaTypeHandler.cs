using GitHub.Octokit.Client.Middleware.Options;
using Microsoft.Kiota.Http.HttpClientLibrary.Extensions;

namespace GitHub.Octokit.Client.Middleware;

/// <summary>
/// Represents a handler that adds the API version header to outgoing HTTP requests.
/// </summary>
class MediaTypeHandler : DelegatingHandler
{
  private const string MediaTypeHeaderKey = "Accept";

  private readonly MediaTypeOptions _mediaTypeOptions;

  public MediaTypeHandler(MediaTypeOptions? mediaTypeOptions = null)
  {
      _mediaTypeOptions = mediaTypeOptions ?? new MediaTypeOptions();
  }

  protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
  {
    if (request == null){
      throw new ArgumentNullException(nameof(request));
    }

    var mediaTypeHandlerOption = request.GetRequestOption<MediaTypeOptions>() ?? _mediaTypeOptions;
    
    if (!request.Headers.Contains(MediaTypeHeaderKey) 
    || !request.Headers.GetValues(MediaTypeHeaderKey)
      .Any(x => mediaTypeHandlerOption.MediaTypeValue.Equals(x, StringComparison.OrdinalIgnoreCase)))
    {
      request.Headers.Add(MediaTypeHeaderKey, mediaTypeHandlerOption.MediaTypeValue);
    }
    return base.SendAsync(request, cancellationToken);
  }
}
