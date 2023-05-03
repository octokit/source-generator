using Azure.Identity;
using Azure.Core;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Http.HttpClientLibrary;
using Octokit;

var token = Environment.GetEnvironmentVariable("GITHUB_TOKEN");
if (string.IsNullOrEmpty(token))
{
	Console.WriteLine("GITHUB_TOKEN environment variable is not set");
	//   throw new Exception("GITHUB_TOKEN environment variable is not set");
}

var authProvider = new ApiKeyAuthenticationProvider(token, "Authorization", ApiKeyAuthenticationProvider.KeyLocation.Header, new string[] { "https://api.github.com" });
var requestAdapter = new HttpClientRequestAdapter(authProvider, null, null, new HttpClient
{
	BaseAddress = new Uri("https://api.github.com/")
});
var pathParameters = new Dictionary<string, object>() {
		{"baseurl", "https://api.github.com/"},
	};
var octocatPathParams = new Octokit.Octocat.OctocatRequestBuilder(pathParameters, requestAdapter);
Action<Octokit.Octocat.OctocatRequestBuilder.OctocatRequestBuilderGetRequestConfiguration> requestConfig =
	delegate (Octokit.Octocat.OctocatRequestBuilder.OctocatRequestBuilderGetRequestConfiguration config)
	{
		config.QueryParameters.S = "Hey yo!";
		config.Headers.Add("Accept", "application/vnd.github.v3+json");
	};

try
{
	Stream requestPathParam = await octocatPathParams.GetAsync(requestConfig);
	StreamReader rPathParams = new StreamReader(requestPathParam);
	string responsePathParams = rPathParams.ReadToEnd();
	Console.WriteLine(responsePathParams);
}
catch (Exception ex)
{
	Console.WriteLine(ex.Message);
	throw;
}


// Passing path parms

// Basic request

// var client = new Octokit.ApiClient(requestAdapter);
// var request = await client.Octocat.GetAsync();

// StreamReader reader = new StreamReader(request);
// string response = reader.ReadToEnd();
// Console.WriteLine(response);

// Basic request