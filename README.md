# BitBucket Pull Request Manager [![Build
Status](https://travis-ci.org/alexhokl/go-bb-pr.svg?branch=master)](https://travis-ci.org/alexhokl/go-bb-pr)

This is a command line tool to help working with BitBucket pull requests.

Usage:
  go-bb-pr [command]

Available Commands:

Command | Description
--- | ---
approve    | Approve the specified pull request
create     | Create the specified pull request
checkout   | Checkout the latest code of the branch of the specified pull request
decline    | Decline the specified pull request
describe   | Describe the specified pull request
list       | List pull requests
merge      | Merge the specified pull request
open       | Open the web page of the specified pull request in a browser
unapprove  | Un-approve the specified pull request

Use "go-bb-pr [command] --help" for more information about a command.

### Installation

If the machine has Go installed, simply run `go get -u
github.com/alexhokl/go-bb-pr`. If not, download the latest binary from
[release](https://github.com/alexhokl/go-bb-pr/releases) page, and drop it to
a directory where it is in of the directories specified in `PATH` environment
variable.

Login to BitBucket is required and it should be set in environment variable
`BB_PR_USERNAME` and `BB_PR_PASSWORD`. Currently, all the API requests are made
in encrypted traffic but credentials stored in environment variables are not
encrypted. This issue should be fixed when this tool has a production release.

### Build

To create a local build, simply execute `go install`.

To create a cross-machine build, for instance, building a Windows executable on
a Mac,

```console
GOOS=windows GOARCH=amd64 go build -o go-bb-pr.win64.exe
```
