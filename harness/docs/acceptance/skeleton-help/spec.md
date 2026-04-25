---
id: ACC-0001
type: acceptance
title: Skeleton commands print help
created: 2026-04-25
status: current
tags: ["skeleton", "help", "cobra"]
scenarios:
  - "Running `harness` with no args prints root help and exits 0"
  - "Running `harness context` with no args prints context help and exits 0"
expected_exit_code: 0
exercises: ["SPEC-0001"]
---

## Scenario: skeleton commands print help

### Given

- The harness binary is built and on PATH (or invoked at `./dist/harness`).
- No fixture state is required — the binary is self-contained.

### When

- The user runs `harness context` with no further args.

### Then

- The exit code is 0.
- Stdout contains the substring `context` (the command name appears
  in its own help output).
- Stdout contains the substring `Usage:` (Cobra's standard help
  structure).

## Notes

This is the simplest possible acceptance scenario — it exercises
the bare command surface that SPEC-0001 specifies. Subsequent
acceptance specs will exercise `validate`, `classify`, `index`, and
`lookup` as those subcommands are built, and will use the
`fixture/` directory to set up document trees the commands operate
on. The `fixture/` directory is empty here but kept on disk because
the runner contract is "cd into `fixture/`, run `cmd`, compare to
`expected.json`" — `fixture/` is always present even when empty.
