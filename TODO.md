# TODO

Findings from a project review (2026-11-14). Checked with `go vet` (clean),
`golangci-lint run` (12 issues) and `golangci-lint config verify` (fails).

Suggested order: the `download` mkdir bug, the `bigfiles` and `rewrites` nil
crashes, then the lint config and the two skipped modules.

## Bugs

- [x] `download/download.go:61` — `os.Mkdir("data")` runs for every URL. It
  fails on the second URL because the folder already exists, and `checkErr`
  panics, so only the first URL ever downloads. Use `os.MkdirAll`, once, in
  `main`.
- [x] `download/download.go:82` — the custom `CheckRedirect` always returns
  `nil`, removing Go's 10-redirect limit, so a redirect loop never ends.
  `resp.StatusCode` is never checked, so a 404 page is saved as the file. A
  URL ending in `/` gives an empty filename.
- [x] `bigfiles/bigfiles.go:13` — `!= "." || != ".."` is always true. The walk
  error is discarded, so an unreadable path arrives with `f == nil` and
  `f.Name()` panics.
- [x] `rewrites/rewrites.go:52-57` — `u.Scheme` is read before `err` is
  checked, so a bad URL panics with a nil pointer. A row with fewer than 3
  fields panics on `each[1:3]`.
- [x] `keywords/keywords.go:20` — `description` reads the `"name"` key. All
  errors are ignored; if `require` is missing, `require.Map()` panics.
- [x] `dial/dial.go:42` — connections are never closed and `net.Dial` has no
  timeout, so a filtered port hangs. `host+":"+port` breaks for IPv6; use
  `net.JoinHostPort`.
- [x] `visitors/visitors.go:62`, `tickets/tickets.go:34` — fixed column
  positions (`entry[9]`, `ticket[50]`) are read without a length check, so a
  short or blank line panics. `tickets` also counts the CSV header row as a
  ticket.
- [x] `lookup/lookup.go:23` — running with no flags exits, so the `hosts.txt`
  default can never be used. Starts one goroutine per host with no limit.
- [x] `pinghosts/pinghosts.go:43` — one host that doesn't resolve stops the
  whole run with `os.Exit(1)`.
- [x] `lookup`, `visitors` — `scanner.Err()` is never checked (`download` fixed), so a
  read error looks the same as end of file.

## Tooling

- [x] `.golangci.yml` is not valid v2 config. `config verify` rejects
  `disable-all`, `linters-settings` and `issues.exclude-dirs`, and `version`
  must be the string `"2"`. Those settings are currently ignored. v2 keys:
  `linters.default: none`, `linters.settings.*`,
  `formatters.settings.goimports.local-prefixes`, `linters.exclusions.paths`.
- [x] `keywords/` and `pinghosts/` are separate modules, so `./...` from the
  root skips them (lint, vet, gosec, `task check`, README `go install`).
  Merge them into the root module, or add a `go.work` and run checks per
  module.
- [ ] Old dependencies: `go-fastping` (used by `pinghosts`) is unmaintained
  since 2016. Run `govulncheck`. (`ioutil` removed and `x/net` updated when the
  modules were merged.)
- [ ] No CI workflow (no `.github/workflows`) and no tests. The hash functions
  and `DiffDays` are easy to test.
- [ ] Makefile and Taskfile overlap: `make check` only lints, `task check`
  also runs goimports and gosec. `make build` uses goreleaser with no config
  file. Keep one.
- [ ] `.gitignore` misses extension-less macOS/Linux binaries and the `data/`
  folder created by `download`.

## Cleanup

- [ ] `md5`, `sha1`, `sha256`, `sha3` are copies apart from the hash function.
  Replace with one `hashsum -algo=sha256` tool. `sha256` defaults to
  `"sha256.go"`; the others default to empty.
- [ ] Add `// #nosec G401 G501` with a reason to the md5/sha1 checksum code.
- [ ] Wrap errors with `%w` instead of `%v`.
- [ ] `email/email.go:10` — sender is a real-looking `@fox.com` address; use an
  `example.com` address.
- [ ] Hard-coded inputs: `countdown` has a fixed 2020 date; `tags` calls
  StackExchange API v2.2. Make these flags or arguments.
- [ ] Remaining revive warnings: missing package comments (unused `via`
  parameter in `download` fixed).
