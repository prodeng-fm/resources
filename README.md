# agent-tools

Go-workspace monorepo of agent tools. Each tool is its own module
under the workspace root, versioned independently and built into a
shared `dist/`.

## Tools

- [`harness/`](./harness) — CLI for repository information
  architecture (validate frontmatter, classify documents, maintain
  indexes), growing into a broader harness-engineering toolkit.
- `cm/` — capability modelling (skeleton coming).

## Build

```sh
make build      # builds every declared tool into dist/
make test       # runs every tool's test suite
make clean      # empties dist/
```

## Decisions

Workspace-level decisions live in [`docs/adrs/`](./docs/adrs).
Per-tool decisions live in each tool's own `docs/adrs/`.

## Contributing

All commits in this repo are authored as
`prodengfm <info@productengineering.fm>`. Configure your clone:

```sh
git config user.name  prodengfm
git config user.email info@productengineering.fm
git config core.hooksPath .githooks
```

Author identity is verified by a pre-commit hook
([`.githooks/pre-commit`](./.githooks/pre-commit)) and by CI
([`.github/workflows/author-check.yml`](./.github/workflows/author-check.yml)).
Non-conforming commits cannot merge to `main`.
