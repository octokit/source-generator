using Microsoft.Kiota.Abstractions;

namespace GitHub.Octokit.Client.Middleware.Options;

public class UserAgentOptions : IRequestOption
{
    public string ProductName { get; set; } = "octokit-net";
    public string ProductVersion { get; set; } = GetProductVersion();

  private static string GetProductVersion() =>
    //TODO : Get the version from the assembly
    "0.0.0";
}