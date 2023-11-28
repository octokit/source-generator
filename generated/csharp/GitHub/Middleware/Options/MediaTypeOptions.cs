using Microsoft.Kiota.Abstractions;

namespace GitHub.Octokit.Client.Middleware.Options;

public class MediaTypeOptions : IRequestOption
{

  public string MediaTypeValue { get; set; } = GetMediaType();

  /// <summary>
  /// This needs to be extended to consider all the media types that are supported by the SDK
  /// https://docs.github.com/en/rest/overview/media-types?apiVersion=2022-11-28
  /// </summary>
  /// <returns></returns>
  private static string GetMediaType() => "application/vnd.github+json";
}