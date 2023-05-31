using Octokit.Client;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

// API requires no authentication, so use the anonymous
var authProvider = new AnonymousAuthenticationProvider();
var adapter = new HttpClientRequestAdapter(authProvider);
var client = new OctokitClient(adapter);

var user = await client.Users["nickfloyd"].GetAsync();
user?.AdditionalData.ToList().ForEach(x => Console.WriteLine(x.Key + ": " + x.Value));