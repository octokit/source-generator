using Microsoft.Kiota.Abstractions;

namespace GitHub.Octokit.Client.Middleware;

class VersionHandler : DelegatingHandler
{
    private const string ApiVersionHeaderKey = "X-GitHub-Api-Version";

    private readonly string _apiVersion;

    public VersionHandler(string apiVersion = "2022-11-28")
    {
        _apiVersion = apiVersion;
    }

    public void Handle(RequestInformation request)
    {
        request.Headers.Add(ApiVersionHeaderKey, _apiVersion);
    }
}