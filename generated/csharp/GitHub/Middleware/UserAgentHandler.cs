using System.Net.Http.Headers;
using Microsoft.Kiota.Http.HttpClientLibrary.Extensions;
using Microsoft.Kiota.Http.HttpClientLibrary.Middleware.Options;


namespace GitHub.Octokit.Client.Middleware;

public class UserAgentHandler : DelegatingHandler
{
    private readonly UserAgentHandlerOption _userAgentOption;

    public UserAgentHandler(UserAgentHandlerOption? userAgentHandlerOption = null)
    {
        _userAgentOption = userAgentHandlerOption ?? new UserAgentHandlerOption();
    }

    protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
    {
        if(request == null)
            throw new ArgumentNullException(nameof(request));

        var userAgentHandlerOption = request.GetRequestOption<UserAgentHandlerOption>() ?? _userAgentOption;

        if(userAgentHandlerOption.Enabled &&
            !request.Headers.UserAgent.Any(x => userAgentHandlerOption.ProductName.Equals(x.Product?.Name, StringComparison.OrdinalIgnoreCase)))
        {
            request.Headers.UserAgent.Add(new ProductInfoHeaderValue(userAgentHandlerOption.ProductName, userAgentHandlerOption.ProductVersion));
        }
        return base.SendAsync(request, cancellationToken);
    }
}
