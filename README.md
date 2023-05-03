# GitHub Organization Recent Contributors

Written by ChatGPT.
A command-line tool in Go to find users who recently committed across all repositories in a GitHub organization within the last week.
The output includes the list of repositories for each user, considering all branches.

## Features

- Lists contributors who have committed in the last week
- Includes all branches in each repository
- Displays the list of repositories for each contributor
- Uses GitHubs API to fetch data

## Requirements

- Go 1.16 or higher
- GitHub Personal Access Token with proper permissions (repo scope)

## Installation

1. Clone this repository
2. Install the required packages:

```bash
go get
```

## Usage

To run the utility, use the following command:

```bash
go run main.go <GitHub API token> <organization>
```

Replace `<GitHub API token>` with your personal GitHub API token and `<organization>` with the organization name you want to analyze. The script will print the users who have committed across all repositories in the organization in the last week along with the list of repositories they have committed to.

## Creating a GitHub Personal Access Token

To create a GitHub Personal Access Token with the proper permissions, follow these steps:

1. Go to your [Personal Access Tokens](https://github.com/settings/tokens) settings page on GitHub.
2. Click on the "Generate new token" button.
3. Give your token a descriptive name, and select the `repo` scope.
4. Click "Generate token" at the bottom of the page.
5. Copy the generated token and use it as `<GitHub API token>` when running the utility.

**Note:** Keep your token secure and do not share it with others. If you suspect your token has been compromised, revoke it immediately and generate a new one.

## Contributing

If you have suggestions for improvements or found a bug, feel free to open an issue or submit a pull request.

## License

This utility is released under the [Apache 2.0 License](LICENSE).
