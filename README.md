# BitBucket Pull Request Manager [![Build Status](https://travis-ci.org/alexhokl/bb.svg?branch=master)](https://travis-ci.org/alexhokl/bb)

This is a command line tool to help working with BitBucket pull requests.

Usage:
  bb [command]

Available Commands:

Command | Description
--- | ---
login         | Log onto BitBucket (and retrieve and store access tokens)
approve       | Approve the specified pull request
create        | Create the specified pull request
checkout      | Checkout the latest code of the branch of the specified pull request
decline       | Decline the specified pull request
describe      | Describe the specified pull request
list          | List pull requests
merge         | Merge the specified pull request
open          | Open the web page of the specified pull request in a browser
unapprove     | Un-approve the specified pull request
list-jira-ids | list all JIRA IDs involved in a pull request

Use "bb [command] --help" for more information about a command.

### Installation

If the machine has Go installed, simply run `go get -u
github.com/alexhokl/bb`. If not, download the latest binary from
[release](https://github.com/alexhokl/bb/releases) page, and drop it to
a directory where it is in of the directories specified in `PATH` environment
variable.

Login to BitBucket is required and OAuth client ID and secret should be set in
configuration file. The default path is `$HOME/.bb.yml`. The content should be
similar to this.

```yaml
client_id: your-client-id
client_secret: your-client-secret
```

See section `Create a consumer` in [Use OAuth on Bitbucket
Cloud](https://support.atlassian.com/bitbucket-cloud/docs/use-oauth-on-bitbucket-cloud/)
for the steps to create `client_id` and `client_secret`.

### Build

To create a local build, simply execute `go install`.

To create a cross-machine build, for instance, building a Windows executable on
a Mac,

```console
GOOS=windows GOARCH=amd64 go build -o bb.win64.exe
```
