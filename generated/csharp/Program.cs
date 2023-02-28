using Azure.Identity;
using Azure.Core;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Http.HttpClientLibrary;

var GHPAT = "TODO";

var authProvider = new AnonymousAuthenticationProvider();
var requestAdapter = new HttpClientRequestAdapter(authProvider);


// Passing path parms

var pathParameters =new Dictionary<string, object>() {
  {"s", "Hey yo!"}
};
var octocatPathParms = new Kiota.Octocat.OctocatRequestBuilder(pathParameters, requestAdapter);
Stream request = await octocatPathParms.GetAsync();

StreamReader rPathParms = new StreamReader(request);
string responsePathParms = rPathParms.ReadToEnd();
Console.WriteLine(responsePathParms);

// Passing path parms

// Basic request

var client = new Kiota.ApiClient(requestAdapter);
var request = await client.Octocat.GetAsync();

StreamReader reader = new StreamReader(request);
string response = reader.ReadToEnd();
Console.WriteLine(text);

// Basic request