# BitBucket Pull Request Manager [![Build Status](https://travis-ci.org/alexhokl/go-bb-pr.svg?branch=master)](https://travis-ci.org/alexhokl/go-bb-pr)

This is a command line tool to help working with BitBucket pull requests.

#### Commands

Command | Description | Example
--- | --- | ---
list | Lists all open pull requests | `pr list`
describe | Dumps information of the specified pull request | `pr describe 939`
checkout | Checkouts the branch of the specified pull request | `pr checkout 939`
approve | Approves the specified pull request | `pr approve 939`
merge | Merges the specified pull request | `pr merge 939`
unapprove | Un-approves the specified pull request | `pr unapprove 939`
decline | Declines the specified pull request | `pr decline 939`
merge | Merges the specified pull request | `pr merge 939`
open | Opens the specified pull request in a browser | `pr open 939`
