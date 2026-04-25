---
id: ADR-0001
type: adr
title: Go and Cobra for the harness CLI
created: 2026-04-25
status: accepted
tags: ["language", "cli", "tooling"]
decision: "Build the harness CLI in Go, using Cobra for command
  structure."
consequences: "Slightly more code for some tasks (frontmatter
  round-tripping, future LLM-as-judge sensors). Tradeoff accepted
  given binary distribution and the discipline a typed language
  imposes on a tool that needs to be reliable."
---

# ADR-0001: Go and Cobra for the harness CLI

## Context

We considered Python+Typer (faster to ship, richer LLM ecosystem,
slightly nicer frontmatter library) and Go+Cobra (single static
binary, faster startup, typed discipline, learning value, gh- and
kubectl-style patterns). The harness has to be reliable over
fast-iterating: it operates on the user's repo, and a runtime crash
or a silent miscount of frontmatter fields is worse than a feature
arriving a week later. A static binary that drops into any repo via
`git subtree` without requiring a Python runtime is also a real
distribution benefit.

## Decision

Go 1.23+ with `github.com/spf13/cobra` for command structure.
Binaries land in the workspace `dist/` via the per-tool Makefile.

## Consequences

Some tasks cost more code in Go than in Python — frontmatter
round-tripping is the obvious one, and future LLM-as-judge sensors
will lean on the Python AI ecosystem. We accept that cost. In
return we get a static binary with no runtime dependency, fast
startup, and a typed compiler that catches a class of mistakes a
tool like this can't afford to make at runtime against a user's
repo. Future LLM-heavy components can shell out or call APIs; the
spine stays Go. This ADR is about the harness specifically — Python
remains the right choice for tools where the inferential ecosystem
dominates the decision.
