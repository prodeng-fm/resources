---
id: ADR-0002
type: adr
title: Noun-verb command structure
created: 2026-04-25
status: accepted
tags: ["cli", "ergonomics"]
decision: "Adopt gh-style noun-verb command structure (harness
  <noun> <verb>) with flat commands for operations that are not
  nouns."
consequences: "Slightly more typing for common commands. The shape
  matches gh, kubectl, and cargo, which is familiar to the user
  population. Adding a new noun is a clean Cobra subcommand
  registration."
---

# ADR-0002: Noun-verb command structure

## Context

Three command shapes were on the table: verb-noun
(`harness validate context`), flat-only
(`harness-validate-context` or one giant root with every operation
hung directly off it), and noun-verb (`harness context validate`).
The CLI surface is going to grow — `context`, `adr`, `skill`,
`sensor`, others — and noun-verb scales: each noun becomes a Cobra
subcommand group with its own verbs underneath. `gh`, `kubectl`,
and `cargo` all use this shape, so the muscle memory is already
there for the user population.

## Decision

Noun-verb is the default. Flat commands like `harness check` or
`harness release` are allowed when the command IS the operation
rather than an action on a noun. Heuristic for "is this a noun":
at least two distinct verbs hang off it. Single-verb concepts stay
flat.

## Consequences

Common commands cost a few extra characters compared to flat.
Adding a new noun is a clean Cobra registration: one file under
`internal/cmd/`, an `init()` that hangs the noun off the root and
the verbs off the noun. The grouping helps `--help` discoverability
— `harness context --help` shows only context verbs instead of a
wall of every operation. The flat-allowed escape hatch keeps the
surface honest: forcing a fake noun (`harness check run`) when
there is no real noun would be worse than no grouping at all.
