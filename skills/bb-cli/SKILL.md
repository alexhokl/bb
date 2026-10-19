---
name: bb-cli
description: Use when the user asks to fetch, list, create, update, manage, or merge BitBucket pull requests, repositories, and projects directly from the terminal.
---

# bb-cli Skill

This skill provides `opencode` with the ability to interact with BitBucket via the local `bb` command-line application.

## Plan / Ask Mode (Read-Only)

Before executing any command, check whether a `<system-reminder>` block containing `Plan mode ACTIVE` is present in your current context. If it is, you are in **read-only mode** and MUST NOT execute any command that creates, updates, deletes, or merges data.

**Forbidden in read-only mode:**
- `bb pr` (create pull request)
- `bb update pr` / `bb approve` / `bb decline` / `bb merge`
- `bb create branch-permission` / `bb update branch-permission` / `bb delete branch-permission`
- Any command with side effects on the remote repository

**Allowed in read-only mode:**
- `bb list pr` / `bb get pr` / `bb list repos` / `bb list workspaces`
- `bb list branch-permissions` / `bb get branch-permission`
- Any read-only fetch or inspection command

If the user requests a mutating operation while in read-only mode, respond:
> "I am currently in plan/read-only mode. I can describe what this operation would do, but I will not execute it until plan mode ends."

## Prerequisites
1. The `bb` binary must be compiled and available in the system's PATH. If not compiled, it can be built from source in the repository using `go install` or downloaded from its releases page.
2. A configuration file must exist at `~/.bb.yml` containing:
   ```yaml
   client_id: your-client-id
   client_secret: your-client-secret
   port: 8080 # or any available local port for the OAuth flow
   ```
3. The user must be logged in. The command `bb login` will open a browser to authenticate and store access tokens.
4. Many `bb` commands that operate on repositories (like creating/listing PRs) infer the repository context from the local git repository. Ensure the working directory is inside a cloned git repository linked to BitBucket when running those.

## How to use this skill
When the user asks to manage or fetch BitBucket pull requests, repositories, or other entities, follow these steps:
1. Verify `bb` is accessible by running `bb --help` in the bash tool.
2. Formulate the appropriate command based on the user's intent using the provided examples below.
3. Execute the command using the `bash` tool.

## Key Commands & Workflows

### 1. Pull Requests
To list pull requests for the current repository:
```bash
bb list pr [--source <branch>] [--destination <branch>] [--detail]
```

To create a new pull request:
```bash
bb pr -t "<title>" -m "<description>" [-s <source-branch>] [-d <destination-branch>]
```

To view details of a specific pull request:
```bash
bb describe -i <pr-id>
```

To approve, decline, or unapprove a pull request:
```bash
bb approve -i <pr-id>
bb decline -i <pr-id>
bb unapprove -i <pr-id>
```

To merge a pull request:
```bash
bb merge -i <pr-id> [--strategy <merge_commit|no_commit|squash>] [--close-source-branch=true|false]
```

To add a comment to a pull request:
```bash
bb comment -i <pr-id> -m "<comment body>"
```

To checkout the branch of a pull request locally:
```bash
bb checkout -i <pr-id>
```

### 2. Exploring Repositories & Projects
To list repositories in a workspace:
```bash
bb list repos -w <workspace-slug> [-p <project-key>]
```

To list projects in a workspace:
```bash
bb list projects -w <workspace-slug>
```

### 3. Users & Permissions
To show the currently authenticated user:
```bash
bb whoami
```

To list users in a workspace:
```bash
bb list users -w <workspace-slug>
```

To check repository permissions:
```bash
bb list repository-permissions -w <workspace-slug> -s <repo-slug>
```

### 4. Branch Permissions (Branch Restrictions)
Manage branch restriction rules (push/force/delete restrictions, merge checks, required approvals, etc.) on a repository. `-w/--workspace` and `-s/--slug` are optional on all of these — if omitted, they are inferred from the git remote of the current directory (like `bb list environment`).

List branch permission rules (rendered as a table):
```bash
bb list branch-permissions [-w <workspace-slug>] [-s <repo-slug>] [--kind <kind>] [--pattern <glob>]
```

Get details of a specific rule:
```bash
bb get branch-permission [-w <workspace-slug>] [-s <repo-slug>] --id <rule-id>
```

Create a new rule (either `--pattern` for a glob match, or `--branch-type` for a branching-model match; only one may be given):
```bash
bb create branch-permission [-w <workspace-slug>] [-s <repo-slug>] --kind <kind> (--pattern <glob> | --branch-type <production|development|bugfix|release|feature|hotfix>) [--value <n>] [--users <uuid1,uuid2>] [--groups <slug1,slug2>]
```

Update an existing rule (only supplied fields are changed):
```bash
bb update branch-permission [-w <workspace-slug>] [-s <repo-slug>] --id <rule-id> [--kind <kind>] [--pattern <glob> | --branch-type <type>] [--value <n>] [--users <uuid1,uuid2>] [--groups <slug1,slug2>]
```

Delete a rule:
```bash
bb delete branch-permission [-w <workspace-slug>] [-s <repo-slug>] --id <rule-id>
```

Valid `--kind` values include: `push`, `force`, `delete`, `restrict_merges`, `require_tasks_to_be_completed`, `require_approvals_to_merge`, `require_default_reviewer_approvals_to_merge`, `require_no_changes_requested`, `require_passing_builds_to_merge`, `require_commits_behind`, `reset_pullrequest_approvals_on_change`, `smart_reset_pullrequest_approvals`, `reset_pullrequest_changes_requested_on_change`, `require_all_dependencies_merged`, `enforce_merge_checks`, `allow_auto_merge_when_builds_pass`.

**Read-only mode note:** `bb list branch-permissions` and `bb get branch-permission` are safe in plan/read-only mode; `create`/`update`/`delete branch-permission` are mutating and forbidden in that mode.

### 5. Environments
The CLI also supports interacting with environments and variables.
- `bb list environment`
- `bb list environment-variable`
- `bb create environment`
- `bb create environment-variable`

*(Note: Discover more flags and nested subcommands using `bb <command> --help` if needed)*