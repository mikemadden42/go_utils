# go_utils
Go utilities

## Tools

| Tool | What it does | Flags |
|---|---|---|
| `bigfiles` | Prints every file larger than 100 MB under a directory | `<dir>` argument |
| `countdown` | Prints the number of days since a start date | `-since` |
| `dial` | Checks whether the TCP ports listed in a JSON file are open | `-file` (`dial.json`) |
| `dirsize` | Prints the size in bytes of each directory under a path | `<dir>` argument |
| `download` | Saves each URL listed in a file into `data/` | `-file` (`urls.txt`) |
| `email` | Sends a test message over SMTP | |
| `inventory` | Renders `template.txt` with sample inventory data | |
| `keywords` | Prints the name, description, keywords and requirements from `composer.json` | `-file` (`composer.json`) |
| `lookup` | Resolves each host in a file and reverse-resolves its addresses | `-file` (`hosts.txt`) |
| `md5`, `sha1`, `sha256`, `sha3` | Print the checksum of a file | `-file` |
| `pinghosts` | Pings each host listed in a file (needs root) | `-file` (`hosts.txt`) |
| `pool` | Demonstrates a worker pool that squares numbers | |
| `rewrites` | Checks the source and destination URLs in a CSV of redirects | `-file` (`rws.csv`) |
| `tags` | Prints the 100 most popular Stack Overflow tags | |
| `tickets` | Groups the tickets in a CSV export by assignee | `-file` (`tickets.csv`) |
| `visitors` | Counts hits by IP, URL or user agent in a caret-separated access log | `-ip`, `-url`, `-agent`, `-log` (`access.log`) |

Defaults are shown in parentheses; run any tool with `-h` for details.

## Installation

Install these with go install:

```bash
go install -ldflags '-s -w' ./...
```

On macOS this installs an `md5` into `$GOBIN` (usually `~/go/bin`). If that
directory comes before `/sbin` on your `PATH`, it hides the system `md5`.

## Development

Common tasks use [Task](https://taskfile.dev):

```bash
task check    # format check, lint, gosec, tests, govulncheck
task test     # tests only
task fmt      # format with goimports
task install  # go install every utility
```
