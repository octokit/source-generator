using System.Net;
using GitHub.Octokit.Client.Middleware;

namespace GitHub.Client;


public static class OctokitClientFactory
{
    public static HttpClient Create(HttpMessageHandler? finalHandler = null)
    {
      var defaultHandlers = CreateDefaultHandlers();
      var handler = ChainHandlersCollectionAndGetFirstLink(finalHandler ?? GetDefaultHttpMessageHandler(), defaultHandlers.ToArray());
      return handler != null ? new HttpClient(handler) : new HttpClient();
    }

    public static IList<DelegatingHandler> CreateDefaultHandlers()
    {
        return new List<DelegatingHandler>
        {
            new VersionHandler(),
            new UserAgentHandler(),
        };
    }

    public static DelegatingHandler? ChainHandlersCollectionAndGetFirstLink(HttpMessageHandler? finalHandler, params DelegatingHandler[] handlers)
    {
      if(handlers == null || !handlers.Any()) return default;
      var handlersCount = handlers.Length;
      for(var i = 0; i < handlersCount; i++)
      {
          var handler = handlers[i];
          var previousItemIndex = i - 1;
          if(previousItemIndex >= 0)
          {
              var previousHandler = handlers[previousItemIndex];
              previousHandler.InnerHandler = handler;
          }
      }
      if(finalHandler != null)
          handlers[handlers.Length-1].InnerHandler = finalHandler;
      return handlers.First();
    }

    public static DelegatingHandler? ChainHandlersCollectionAndGetFirstLink(params DelegatingHandler[] handlers)
    {
      return ChainHandlersCollectionAndGetFirstLink(null,handlers);
    }

    public static HttpMessageHandler GetDefaultHttpMessageHandler(IWebProxy? proxy = null)
    {
      return new HttpClientHandler { Proxy = proxy, AllowAutoRedirect = false };
    }
}
