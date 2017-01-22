# BitBucket Pull Request Manager [![Build Status](https://travis-ci.org/alexhokl/go-bb-pr.svg?branch=master)](https://travis-ci.org/alexhokl/go-bb-pr)

This is a command line tool to help working with BitBucket pull requests.

Usage:
  go-bb-pr [command]

Available Commands:

Command | Description
--- | ---
approve    | Approve a pull request
checkout   | Checkout the latest code of the branch of a pull request
decline    | Decline a pull request
describe   | Describe a pull request
list       | List pull requests
merge      | Merge a pull request
open       | Open the web page of the specified pull request in a browser
unapprove  | Unapprove a pull request

Use "go-bb-pr [command] --help" for more information about a command.

### Installation

If the machine has Go installed, simply run `go get -u github.com/alexhokl/go-bb-pr`. If not, download the latest binary from [release](https://github.com/alexhokl/go-bb-pr/releases) page, and drop it to a directory where it is in of the directories specified in `PATH` environment variable.

Login to BitBucket is required and it should be set in environment variable `bbuser` and `bbpassword`. Currently, all the API requests are made in encrypted traffic but credentials stored in environment variables are not encrypted. This issue should be fixed when this tool has a production release.
