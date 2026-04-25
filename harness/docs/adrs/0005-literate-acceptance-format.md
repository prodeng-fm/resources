---
id: ADR-0005
type: adr
title: Literate-markdown format for acceptance specs
created: 2026-04-25
status: accepted
tags: ["specs", "testing", "format"]
decision: "Acceptance specs are single .md files whose body
  interleaves prose with fenced sh blocks; each block declares its
  expected exit code and stdout/stderr substrings via comment-line
  annotations. Refines ADR-0004's format without superseding its
  type-separation decision."
consequences: "Prose and assertions cannot drift, since the
  assertions ARE the executable artifact. The runner becomes a
  shell-block extractor rather than a multi-file orchestrator.
  Fixtures, when needed, move from nested fixture/ to sibling
  <name>.fixture/ dirs. ACC-0001 is migrated as part of the same
  change; future acceptance specs use only the new shape."
---

# ADR-0005: Literate-markdown format for acceptance specs

## Context

The directory shape from ADR-0004 (`spec.md` + `cmd` +
`expected.json` + `fixture/`) cleanly separated prose from
assertions but allowed the two to drift — `spec.md` could claim
one thing while `expected.json` checked another. Three options
weighed: keep the directory shape (status quo, drift risk),
Gherkin/BDD (heavyweight, verbose, separate parser), and literate
testing (single source of truth; lineage in Knuth's literate
programming, Python `doctest`, Rust `cargo test --doc`). Literate
testing wins on the no-drift property.

## Decision

One `.md` per scenario at `harness/docs/acceptance/<name>.md`. The
body interleaves prose and fenced sh blocks. Each block is one
assertion: the first non-comment line is the command; comment lines
starting with `# exit:`, `# stdout contains:`, or
`# stderr contains:` declare expectations. Default exit code is 0.
Multiple `contains` lines are AND-conjoined. Fixtures, when
present, are sibling `<name>.fixture/` directories the runner
copies to a tempdir before execution.

## Consequences

Prose and assertions are now coupled in the same block; editing one
without the other is impossible. The future runner is simpler —
extract sh blocks, parse annotations, execute, compare — rather
than orchestrating four files per scenario. Refines ADR-0004
without superseding it: the decision that acceptance is a separate
type from spec still holds; only the on-disk shape of an acceptance
spec changes. ACC-0001 is migrated in the same change.
