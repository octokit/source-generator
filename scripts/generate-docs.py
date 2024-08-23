# Load the Python env so that you can run in a sandboxed 
# environment to generate documentation for the GitHub SDK.
# Note: Add the path once the env is setup - follow the prompts
# python3 -m venv ~/.local --system-site-packages

# This is strickly a protoype and still needs a lot of work to ensure data and functionally accuracy
# This script generates markdown documentation for classes and methods in the GitHub SDK. 

import os
import re
from typing import List
import html

# Define input and output directories
INPUT_DIR = "./stage/dotnet/dotnet-sdk/src/GitHub"
OUTPUT_DIR = "./documentation"

# Ensure the output directory exists
os.makedirs(OUTPUT_DIR, exist_ok=True)

# Mock data based on parameter types
mock_data = {
    "int": "123",
    "string": "\"mockString\"",
    "bool": "true",
    "double": "123.45",
    "DateTime": "DateTime.Now",
    "CancellationToken": "",  # Exclude CancellationToken in the example call
    # Add other types and their mock data as needed
}

def fetch_files_from_directory(directory: str) -> List[str]:
    """Recursively fetch all .cs files ending with RequestBuilder.cs from the specified directory."""
    files = []
    for root, _, filenames in os.walk(directory):
        for filename in filenames:
            if filename.endswith('RequestBuilder.cs'):
                files.append(os.path.join(root, filename))
    return files

def read_file_content(file_path: str) -> str:
    """Read the content of a file."""
    with open(file_path, 'r', encoding='utf-8') as file:
        return file.read()

def extract_classes_and_methods(file_content: str) -> List[dict]:
    """Extract classes, their methods, and comments from file content."""
    class_pattern = re.compile(r'\bclass\s+(\w+)')
    method_pattern = re.compile(r'\bpublic\s+(?:async\s+)?\S+\s+(GetAsync|PutAsync|DeleteAsync|PostAsync)\s*\(([^)]*)\)')
    constructor_pattern = re.compile(r'\bpublic\s+\w+\s*\(')  # Pattern to match constructors
    comment_pattern = re.compile(r'///\s*(.*)')
    param_pattern = re.compile(r'(\w+)\s+(\w+)')

    classes = []
    current_class = None
    current_comments = []
    method_names_set = set()

    file_lines = file_content.splitlines()
    i = 0
    while i < len(file_lines):
        line = file_lines[i]
        comment_match = comment_pattern.search(line)
        if comment_match:
            current_comments.append(html.unescape(comment_match.group(1).strip()))
        else:
            class_match = class_pattern.search(line)
            if class_match:
                if current_class and current_class['methods']:  # Save the previous class if it has methods
                    classes.append(current_class)
                current_class = {'name': class_match.group(1), 'methods': []}
                method_names_set = set()
                current_comments = []
            else:
                method_match = method_pattern.search(line)
                if method_match and current_class is not None:
                    method_name = method_match.group(1)
                    if method_name not in method_names_set:
                        parameters = method_match.group(2)
                        param_list = []

                        for param_match in param_pattern.findall(parameters):
                            param_type, param_name = param_match
                            param_list.append((param_type, param_name))

                        formatted_comments = format_comments("\n".join(current_comments))
                        current_class['methods'].append({
                            'name': method_name,
                            'parameters': param_list,
                            'comments': formatted_comments
                        })
                        method_names_set.add(method_name)
                    current_comments = []
                elif constructor_pattern.search(line):  # Clear comments if a constructor is found
                    current_comments = []
        i += 1

    if current_class and current_class['methods']:  # Add the last class found if it has methods
        classes.append(current_class)

    return classes

def format_comments(comments: str) -> str:
    """Format comments for better readability."""
    # Replace specific HTML entities
    comments = comments.replace("&lt;", "<").replace("&gt;", ">")
    formatted_comments = html.unescape(comments)  # Replace remaining HTML entities
    
    # Replace XML tags with readable text
    formatted_comments = re.sub(r'<summary>', '', formatted_comments)
    formatted_comments = re.sub(r'</summary>', '', formatted_comments)
    formatted_comments = re.sub(r'<see href="([^"]+)"\s*/>', r'[see]("\1")', formatted_comments)
    
    # Extract exceptions and parameters separately
    exceptions = re.findall(r'<exception cref="([^"]+)">(.*?)</exception>', formatted_comments)
    parameters = re.findall(r'<param name="([^"]+)">(.*?)</param>', formatted_comments)
    returns = re.findall(r'<returns>(.*?)</returns>', formatted_comments)
    
    # Remove exceptions, parameters, and returns from comments
    formatted_comments = re.sub(r'<exception cref="[^"]+">.*?</exception>', '', formatted_comments)
    formatted_comments = re.sub(r'<param name="[^"]+">.*?</param>', '', formatted_comments)
    formatted_comments = re.sub(r'<returns>.*?</returns>', '', formatted_comments)
    
    # Replace <returns> tags with **Returns:** followed by the captured text
    returns_text = "\n".join([f"**Returns:** {text}" for text in returns])
    # Remove any other remaining tags
    formatted_comments = re.sub(r'<[^>]+>', '', formatted_comments)  # Remove any other tags
    formatted_comments = formatted_comments.strip()
    # Ensure new lines are properly handled
    formatted_comments = re.sub(r'\n', '\n\n', formatted_comments)
    
    # Append exceptions and parameters at the end
    if parameters:
        formatted_comments += "\n\n**Parameters:**\n"
        for param in parameters:
            formatted_comments += f"- `{param[0]}`: {param[1]}\n"
    
    if exceptions:
        formatted_comments += "\n\n**Exceptions:**\n"
        for exc in exceptions:
            formatted_comments += f"- `{exc[0]}`: {exc[1]}\n"
    
    # Append returns at the end
    if returns_text:
        formatted_comments += f"\n\n{returns_text}\n"

    return formatted_comments

def generate_documentation(classes: List[dict], file_path: str, folder_name: str) -> str:
    """Generate markdown documentation for classes and their methods."""
    file_name = os.path.basename(file_path)
    doc_content = f"# Documentation for {file_name}\n\n"

    for cls in classes:
        doc_content += f"## Class: {cls['name']}\n\n"
        doc_content += f"### Methods:\n\n"
        for method in cls['methods']:
            params = method['parameters']
            param_str = ", ".join([mock_data.get(param[0], f"default({param[0]})") for param in params if param[0] != "CancellationToken"])
            comments = method['comments']
            doc_content += f"#### Method: {method['name']}\n\n"
            if comments:
                doc_content += f"**Description:**\n{comments}\n\n"
            doc_content += f"**Signature:** `{method['name']}({', '.join([f'{param[0]} {param[1]}' for param in params])})`\n\n"
            doc_content += f"**Example Usage:**\n"
            doc_content += (
                f"```csharp\n"
                f"using GitHub;\n"
                f"using GitHub.Octokit.Client;\n"
                f"using GitHub.Octokit.Client.Authentication;\n\n"
                f"public class Example\n"
                f"{{\n"
                f"    public static async Task Run()\n"
                f"    {{\n"
                f"        var tokenProvider = new TokenProvider(Environment.GetEnvironmentVariable(\"GITHUB_TOKEN\") ?? \"\");\n"
                f"        var adapter = RequestAdapter.Create(new TokenAuthProvider(tokenProvider));\n"
                f"        var gitHubClient = new GitHubClient(adapter);\n\n"
                f"        try\n"
                f"        {{\n"
                f"            var result = await gitHubClient.{folder_name}.{method['name']}({param_str});\n"
                f"            Console.WriteLine(result);\n"
                f"        }}\n"
                f"        catch (Exception e)\n"
                f"        {{\n"
                f"            Console.WriteLine(e.Message);\n"
                f"        }}\n"
                f"    }}\n"
                f"}}\n"
                f"```\n\n"
            )

    return doc_content

def main():
    files = fetch_files_from_directory(INPUT_DIR)
    
    for file_path in files:
        relative_path = os.path.relpath(file_path, INPUT_DIR)
        output_path = os.path.join(OUTPUT_DIR, relative_path)
        output_folder_path = os.path.dirname(output_path)
        os.makedirs(output_folder_path, exist_ok=True)

        folder_name = os.path.basename(os.path.dirname(file_path))
        file_content = read_file_content(file_path)
        classes = extract_classes_and_methods(file_content)
        if classes:
            doc_content = generate_documentation(classes, file_path, folder_name)
            if doc_content.strip():  # Write only if doc_content is not empty
                doc_file_path = f"{output_path}.md"
                with open(doc_file_path, 'w', encoding='utf-8') as doc_file:
                    doc_file.write(doc_content)

if __name__ == "__main__":
    main()