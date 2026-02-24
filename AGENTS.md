# AGENTS.md - AI Coding Agent Guidelines for bb

This document provides guidelines for AI coding agents working in the `bb` (BitBucket CLI) repository.

## Project Overview

- **Language**: Go 1.23.4 (toolchain 1.24.1)
- **Type**: CLI application for managing BitBucket pull requests
- **CLI Framework**: Cobra + Viper for commands and configuration
- **API**: Swagger-generated BitBucket API client in `/swagger`

## Project Structure

```
bb/
├── main.go              # Entry point - calls command.Execute()
├── command/             # CLI commands (Cobra commands)
├── models/              # Data models and output helpers
├── swagger/             # Auto-generated BitBucket API client (DO NOT EDIT)
├── go.mod               # Go module definition
└── Makefile             # Build commands
```

## Build/Lint/Test Commands

```bash
go build                                    # Build the binary
go install                                  # Install locally
go test ./...                               # Run all tests
go test ./command                           # Run tests in a specific package
go test -run TestFunctionName ./command     # Run a single test by name
go test -v ./...                            # Run tests with verbose output
golangci-lint run                           # Linting (if installed)
go mod tidy                                 # Tidy dependencies
```

## Code Style Guidelines

### Import Organization
Group imports in order, separated by blank lines: standard library, external deps, internal.

```go
import (
    "context"
    "fmt"

    "github.com/spf13/cobra"

    "github.com/alexhokl/bb/models"
)
```

### Naming Conventions
- **Files**: `snake_case.go` (e.g., `create_pullrequest.go`)
- **Packages**: lowercase, single word (e.g., `command`, `models`)
- **Exported**: `PascalCase` (e.g., `PrintShortDescription`)
- **Unexported**: `camelCase` (e.g., `runApprove`, `validateToken`)
- **Options structs**: `<command>Options` (e.g., `approveOptions`)
- **Command variables**: `<command>Cmd` (e.g., `approveCmd`)

### Command File Structure

```go
package command

type exampleOptions struct {
    idOptions              // Embed common options if needed
    someFlag string
}

var exampleOpts exampleOptions

func (opts *exampleOptions) validate() error {
    return nil
}

var exampleCmd = &cobra.Command{
    Use:   "example",
    Short: "Short description of command",
    RunE:  runExample,
}

func init() {
    rootCmd.AddCommand(exampleCmd)
    flags := exampleCmd.Flags()
    flags.StringVarP(&exampleOpts.someFlag, "flag", "f", "", "Flag description")
}

func runExample(_ *cobra.Command, _ []string) error {
    if err := exampleOpts.validate(); err != nil {
        return err
    }
    // Command implementation
    return nil
}
```

### Error Handling
- Return errors; don't panic or log.Fatal in library code
- Use `fmt.Errorf()` for creating/wrapping errors
- Validate inputs early

```go
if opts.id <= 0 {
    return fmt.Errorf("invalid pull request ID")
}
if err != nil {
    return fmt.Errorf("failed to get repository info: %w", err)
}
```

### Struct Tags and Comments
```go
// User represents a BitBucket user
type User struct {
    Username    string `json:"username"`
    DisplayName string `json:"display_name"`
}
```

### API Client Usage Pattern
```go
savedToken, err := authhelper.LoadTokenFromViper()
if err != nil {
    return err
}
auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
config := swagger.NewConfiguration()
client := swagger.NewAPIClient(config)
```

## Important Notes

### Do NOT Edit
- **`/swagger` directory**: Auto-generated from BitBucket API OpenAPI spec. Never manually edit.

### Configuration
- Config file: `$HOME/.bb.yml`
- Required keys: `client_id`, `client_secret`, `port`

### Key Dependencies
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration
- `github.com/alexhokl/helper` - Shared utilities (includes git helpers)
- `github.com/fatih/color` - Terminal colors
- `golang.org/x/oauth2` - OAuth2 authentication

### Git Operations
Use `github.com/alexhokl/helper/git` for git operations rather than executing git commands directly.

### Output Formatting
- Use `fmt.Printf` for standard output
- Use `github.com/fatih/color` for colored output
- Output helpers are in `models/print_helper.go`
