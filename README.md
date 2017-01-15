# BitBucket Pull Request Manager [![Build Status](https://travis-ci.org/alexhokl/go-bb-pr.svg?branch=master)](https://travis-ci.org/alexhokl/go-bb-pr)

This is a command line tool to help working with BitBucket pull requests.

#### Commands

##### list

To list open pull requests of the current repository. For example,

```console
bb-pr list
```

##### describe

To dump information of a pull request of the current repository. For example,

```console
bb-pr describe 939
```

##### checkout

To checkout the branch of a pull request of the current repository. For example,

```console
bb-pr checkout 939
```

##### approve

To approve a pull request. For example

```console
bb-pr approve 939
```

##### merge

To merge a pull request. For example

```console
bb-pr merge 939
```

##### unapprove

To unapprove an approved pull request. For example

```console
bb-pr unapprove 939
```

##### decline

To decline a pull request. For example

```console
bb-pr decline 939
```
