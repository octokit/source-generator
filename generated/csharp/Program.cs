using Octokit.Client;
using GitHubAuthentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

var token = Environment.GetEnvironmentVariable("GITHUB_TOKEN") ?? "";
var githubRequestAdapter = new HttpClientRequestAdapter(new GitHubTokenAuthenticationProvider("Octokit.Gen",token));

var gitHubClient = new OctokitClient(githubRequestAdapter);
var pullRequests = await gitHubClient.Repos["octokit"]["octokit.net"].Pulls.GetAsync();

if (pullRequests == null)
{
    Console.WriteLine("No pull requests found.");
    return;
}

foreach(var pullRequest in pullRequests)
{
  Console.WriteLine($"#{pullRequest.Number} - {pullRequest.Title}");

  var pullRequestComnments = await gitHubClient.Repos["octokit"]["octokit.net"].Pulls[pullRequest.Number.Value].Comments.GetAsync();
  if (pullRequestComnments == null)
  {
    Console.WriteLine($"#{pullRequest.Number} - {pullRequest.Title} - No reviews found.");
    continue;
  } else {
    foreach(var comments in pullRequestComnments) {
      Console.WriteLine($"#{comments.Body}\n");
    }
  }
}
