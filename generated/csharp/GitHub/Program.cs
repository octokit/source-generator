using GitHub.Client;
using GitHub.Octokit;
using GitHub.Authentication;

var token = Environment.GetEnvironmentVariable("GITHUB_TOKEN") ?? "";
var octokitRequest = OctokitRequestAdapter.Create(new GitHubTokenAuthenticationProvider("Octokit.Gen",token));
var gitHubClient = new OctokitClient(octokitRequest);

var pullRequests = await gitHubClient.Repos["octokit"]["octokit.net"].Pulls.GetAsync();

if (pullRequests == null)
{
    Console.WriteLine("No pull requests found.");
    return;
}

foreach(var pullRequest in pullRequests)
{
  Console.WriteLine($"#{pullRequest.Number} - {pullRequest.Title}");

  var pullRequestComnments = await gitHubClient.Repos["octokit"]["octokit.net"]
    .Pulls[pullRequest.Number.Value]
    .Comments.GetAsync();
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
