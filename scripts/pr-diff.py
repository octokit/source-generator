import requests
import difflib
import re
import sys


def fetch_pr_diff(repo, pr_number):
    """Fetch the diff for a given PR from a repository."""
    url = f"https://api.github.com/repos/{repo}/pulls/{pr_number}"
    headers = {"Accept": "application/vnd.github.v3.diff"}
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        return response.text
    else:
        return None


def summarize_diff(diff, verbose=False):
    # TODO: This needs to be cleaned up and improved
    """Attempt to summarize the diff, added and removed lines."""
    addDictionary = {}
    removeDictionary = {}
    added_lines = 0
    removed_lines = 0
    add = ""
    remove = ""
    for line in diff.split("\n"):
        if "kiotaVersion" in line and not verbose:
            print(line)

        if line.startswith("+") and not line.startswith("+++"):
            added_lines += 1
            add = line
            item = parse_lines(line)
            if item:
                addDictionary[item] = True
        elif line.startswith("-") and not line.startswith("---"):
            removed_lines += 1
            remove = line
            item = parse_lines(line)
            if item:
                removeDictionary[item] = True
        elif add and remove and verbose:
            print(add + "\n" + remove + "\n")
            add, remove = "", ""

    commit_msg = "FEAT:"
    if len(addDictionary) > 0:
        print("Items added/updated:")
        print(", ".join(addDictionary.keys()) + "\n")
        commit_msg += "| Added/Updated [" + ", ".join(addDictionary.keys()) + "]"
    if len(removeDictionary) > 0:
        print("Items removed:")
        print(", ".join(removeDictionary.keys()) + "\n")
        commit_msg += "| Removed [" + ", ".join(removeDictionary.keys()) + "]"

    print(
        "NOTE: Mismatches between items added and removed indicate that there may be breaking changes introduced in this changeset. \n"
    )

    print("Commit message: " + commit_msg)

    summary = f"Added lines: {added_lines}\nRemoved lines: {removed_lines}"
    return summary


def parse_lines(line):
    # This needs to be updated to handle other languages
    # C# specific
    match = re.search(r"public (\w+)", line)
    if match:
        return match.group(1)
    # Go specific
    match = re.search(r"func (\w+)", line)
    if match:
        return match.group(1)
    else:
        return ""


def compare_lines(line1, line2):
    diff = difflib.ndiff(line1, line2)
    return "\n".join(list(diff))


if len(sys.argv) == 3:
    repo = sys.argv[1]
    pr_number = int(sys.argv[2])
    diff = fetch_pr_diff(repo, pr_number)
    if diff:
        summary = summarize_diff(diff, True)
        print(summary)
    else:
        print("Failed to fetch PR diff.")
else:
    print("Usage: python3 ./scripts/pr-diff.py <repo> <pr_number>")
    print("Example: python3 ./scripts/pr-diff.py octokit/dotnet-sdk 123")
    sys.exit(1)
