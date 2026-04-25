---
id: ACC-0001
type: acceptance
title: Skeleton commands print help
created: 2026-04-25
status: current
tags: ["skeleton", "help", "cobra"]
exercises: ["SPEC-0001"]
---

The harness skeleton ships two Cobra commands — `harness` (root)
and `harness context` (parent for future subcommands) — both of
which print help when invoked with no further args and exit
cleanly. This spec exercises both.

## `harness context`

The context parent command prints its own help. `RunE` in
`internal/cmd/context.go` calls `cmd.Help()`, so invoking with no
args exits 0 and emits Cobra's standard help structure.

```sh
# exit: 0
# stdout contains: context
# stdout contains: Usage:
harness context
```

## `harness` (no args)

The root command has no `Run`/`RunE`, so Cobra's default behavior
prints the root help listing the available subcommands.

```sh
# exit: 0
# stdout contains: Available Commands
# stdout contains: context
harness
```
