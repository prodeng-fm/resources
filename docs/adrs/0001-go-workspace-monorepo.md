---
id: ADR-0001
type: adr
title: Multi-tool Go workspace for agent-tools
created: 2026-04-25
status: accepted
tags: ["workspace", "monorepo", "structure"]
decision: "Organize agent-tools as a Go workspace containing
  independent tools as sibling modules, with shared dist/ at the
  workspace root."
consequences: "Each tool versions independently and has its own
  internal/ scope. A workspace-level Makefile orchestrates builds.
  Subtree-distributing one tool to a consumer repo is clean. Adds
  go.work and the per-tool/workspace Makefile split as new
  structural overhead."
---

# ADR-0001: Multi-tool Go workspace for agent-tools

## Context

We want multiple agent tools (harness, cm, future others) sharing
some conventions but versioning independently. Three options
considered: separate repos, single-module monorepo, Go-workspace
monorepo. Separate repos break the co-evolution story;
single-module forces tools to share versions and import paths.
Go workspaces give us the middle ground.

## Decision

Use Go workspaces (`go.work`) with each tool as its own module
under the workspace root. Workspace-level `docs/adrs/` for
decisions affecting the workspace as a whole; per-tool
`docs/adrs/` for decisions internal to a tool.

## Consequences

Per-tool `internal/` packages are genuinely private to each tool.
Shared code, when it emerges, will live in its own internal
module imported explicitly. The workspace Makefile is the build
entrypoint for CI; per-tool Makefiles remain for local
convenience.
