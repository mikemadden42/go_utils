# go_utils
Go utilities

## Installation

Install these with go install:

```bash
go install -ldflags '-s -w' ./...
```

## Development

Common tasks use [Task](https://taskfile.dev):

```bash
task check    # format check, lint, gosec, tests, govulncheck
task test     # tests only
task fmt      # format with goimports
task install  # go install every utility
```
