---
id: SPEC-0001
type: spec
title: Harness CLI skeleton
created: 2026-04-25
status: current
tags: ["cli", "skeleton", "cobra"]
scope: "harness/cmd/harness, harness/internal/cmd"
contract: "Two Cobra commands exist: `harness` (root, prints help)
  and `harness context` (parent for future subcommands, prints
  help). A Makefile builds, tests, and cleans the binary at the
  workspace-shared dist/."
---

# SPEC-0001: Harness CLI skeleton

## Scope

This spec covers the bootstrap shape of the harness CLI: the root
command (`harness`), the `context` parent command, and the
`harness/Makefile` that builds, tests, and cleans the binary. It
does NOT cover frontmatter validation, classification, indexing,
lookup, or any of the verbs that will eventually hang off
`context` and other nouns — each gets its own SPEC-NNNN as it is
built.

## Commands

### `harness`

- **Behavior:** the root command has no `Run`/`RunE`, so Cobra
  prints the root help and exits 0 when invoked with no args or
  with `--help`.
- **Available subcommands:** `context` (the only manually
  registered command), plus `help` and `completion` that Cobra
  auto-adds.
- **Exit codes:** 0 on help display. On an unknown flag or
  subcommand Cobra writes an error to stderr and the `Execute()`
  wrapper in `internal/cmd/root.go` exits 1.

### `harness context`

- **Behavior:** the context command has a `RunE` that calls
  `cmd.Help()`, so invoking with no args prints the context help
  and exits 0.
- **Subcommands:** none registered yet. `init`, `validate`,
  `classify`, `index`, and `lookup` arrive in future specs.
- **Exit codes:** same as root.

## Build and test

`harness/Makefile` exposes the following targets:

- `make build` → `go build -o ../dist/harness ./cmd/harness`. The
  output lands in the workspace-shared `dist/`.
- `make test` → `go test ./...` within `harness/`.
- `make tidy` → `go mod tidy`.
- `make run` → `go run ./cmd/harness`. Compiles and runs in one
  step via Go's temporary build cache; no artifact in `dist/`.
- `make clean` → removes `../dist/harness`.
- `make` (default `all`) → equivalent to `make build`.

The workspace-root `Makefile` delegates to `harness/Makefile` via
its `TOOLS` loop, so `make build` from the workspace root and
`make build` from `harness/` produce the same artifact at
`dist/harness`.

## Test coverage

`harness/internal/cmd/root_test.go` has one test,
`TestRootAndContextPrintHelp`, with two subtests. It invokes
`rootCmd` with no args and asserts the captured output contains
the expected root long-help substring, then does the same with
`["context"]` against the context long help. This validates that
both commands print help when run bare. Nothing else is tested —
no exit-code assertions, no explicit `--help` flag exercise, no
error paths.

## Out of scope

This spec deliberately does NOT specify: frontmatter parsing or
validation, classification of documents against `context.toml`,
index emission, lookup commands, skill integration, sensor
wiring, or any LLM-as-judge mechanism. Each gets its own
SPEC-NNNN when built.
