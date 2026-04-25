# CLAUDE.md

Operational rules for working in this repo. Design rationale lives in
`docs/adrs/` (workspace-scope) and `harness/docs/adrs/` (harness-scope).
Read those for *why*; this file is *how*.

## Authorship (load-bearing)

Every commit MUST be authored AND committed as
`prodengfm <info@productengineering.fm>`. Verified by:

- `.githooks/pre-commit` — local hook (activate with
  `git config core.hooksPath .githooks`)
- `.github/workflows/author-check.yml` — server-side gate on every
  push and PR

A clone's local git config may legitimately be set to a different
identity for day-to-day work elsewhere; do NOT modify it from inside
this repo. Override per-commit instead:

```sh
git -c user.name='prodengfm' -c user.email='info@productengineering.fm' commit ...
```

## Document discipline

- **Append-only ADRs and specs.** Each new decision/increment gets a
  new file (ADR-NNNN, SPEC-NNNN, ACC-NNNN). Never edit a past ADR or
  spec; supersede or refine it via a new one with cross-reference.
  ADR-0005 is the most recent example (refines ADR-0004 without
  superseding it).
- **Spec follows code.** When a spec contradicts the code, fix the
  spec. Specs are harvested descriptions of reality, not wishlists.
- **Acceptance specs are literate markdown** (per ADR-0005): one
  `.md` per scenario, fenced sh blocks with `# exit:` /
  `# stdout contains:` / `# stderr contains:` annotations. No more
  directory shape. Fixtures, when needed, live in sibling
  `<name>.fixture/` directories.

## Build

- `make build` from workspace root → builds every tool in `TOOLS`
  into `dist/`
- `make test` from workspace root → runs every tool's test suite
- `make build` from `harness/` → same artifact at `dist/harness`

## Layout

- `docs/adrs/` — workspace-scope decisions (cross-cutting only)
- `<tool>/docs/{adrs,specs,acceptance}/` — per-tool decisions, impl
  specs, executable acceptance specs
- `<tool>/internal/` — private to that tool (Go `internal` package)
- `dist/` — shared build output, gitignored except `.gitkeep`

## Schema

`harness/context.toml` declares the IA schema (base + per-type
extensions). The future `harness context classify` walks the tree and
validates every doc against it. `[classify].ignore` lists paths
skipped during the walk.

## Push and branches

Direct push to `main` is currently blocked by the Claude Code
sandbox. Once branch protection is enabled on GitHub, the workflow
becomes: feature branch → PR → CI's `author-check` validates →
merge. Until then, the user pushes manually.
