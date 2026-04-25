---
id: ADR-0003
type: adr
title: Base plus type-extension frontmatter schema
created: 2026-04-25
status: accepted
tags: ["schema", "frontmatter", "ia"]
decision: "Frontmatter schemas compose: every doc satisfies a base
  schema plus a per-type extension declared in context.toml."
consequences: "The classifier runs validation in two passes (base,
  then type). Schemas are slightly more verbose to declare. Worth
  it for the indexability and the shared discipline across types."
---

# ADR-0003: Base plus type-extension frontmatter schema

## Context

Three options for the schema shape: one flat schema covering every
doc (simple but rigid — every type would have to carry every
field), per-type schemas with no shared base (flexible but no
cross-type consistency, and the index command would need
type-specific code paths for shared concepts like `tags`), and
base + per-type extension. Base+extension lets the index pivot on
shared fields cheaply (`id`, `type`, `title`, `tags`) while letting
each type declare what's specific to it (`decision` and
`consequences` for ADRs, `scenarios` and `expected_exit_code` for
acceptance specs).

## Decision

`context.toml` declares `[base]` once and `[types.X]` per type.
Validation composes them: required fields union, optional fields
union. Type-specific overrides of base fields are not allowed at
this stage. The classifier validates base first, then the type
extension; failures from each pass are reported separately.

## Consequences

Schemas are slightly more verbose to declare than a single flat
schema, and validation runs two passes per doc (negligible cost).
Worth it for indexability: `harness context index` can pivot on
shared base fields without knowing the type, so adding a new type
doesn't force changes to the index command. This shape also leaves
room for types to declare schema fragments later (a deprecation
mixin, say) without forking the base — but that lives behind a
future ADR if the need is real, not now.
