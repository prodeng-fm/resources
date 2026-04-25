---
id: ADR-0004
type: adr
title: Specs and acceptance specs are separate types
created: 2026-04-25
status: accepted
tags: ["specs", "testing", "ia"]
decision: "Implementation specs (prose contracts) and acceptance
  specs (executable scenarios) live under docs/specs/ and
  docs/acceptance/ respectively, as distinct types in context.toml."
consequences: "Two directories to maintain. Acceptance specs are
  directories (not single files) containing spec.md, fixture/, cmd,
  and expected.json — the classifier ignores fixture/ contents.
  Future `harness acceptance run` subcommand executes them."
---

# ADR-0004: Specs and acceptance specs are separate types

## Context

Two alternatives were on the table: one `docs/specs/` directory
with a `kind` field distinguishing prose specs from executable
specs, or a single `docs/validate/` for everything executable. The
first blurs lifecycles — implementation specs evolve with design,
acceptance specs evolve with bugs, and conflating them in one
directory hides which is which when scanning the tree. The second
collides with the future `harness validate` subcommand name, which
will already do enough work in the user's head without overloading
the directory name on top.

## Decision

Two top-level types under `harness/docs/`. `docs/specs/` holds
prose specs as `.md` files. `docs/acceptance/<name>/` holds
acceptance specs as directories — `spec.md` carries the
frontmatter, `fixture/` is a self-contained tree the command runs
against, `cmd` is a one-line invocation, `expected.json` carries
the expected exit code and substring matches.

## Consequences

Two directories to maintain instead of one, and the classifier
needs to know `fixture/` is ignored — already declared in
`context.toml`'s `[classify].ignore` as `docs/acceptance/*/fixture`.
The directory-shaped acceptance type generalizes: future commands
(`harness check`, `harness release`) get the same acceptance
pattern. One acceptance framework drives every subcommand, so
executable-spec discipline becomes a property of the tool, not a
property of any one command.
