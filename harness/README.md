# harness

`harness` is a CLI for modern engineering harnessing — information
architecture, encoding discipline, policy as code, controls, context
lookup, and enterprise knowledge.

The first milestone is repository information architecture: validating
frontmatter, classifying documents, and maintaining indexes.

## Status

Runnable skeleton. Only `harness` and `harness context` exist, and they
both print help.

## Layout

```
cmd/harness/        entrypoint
internal/cmd/       cobra commands
internal/config/    context.toml loading (later)
internal/frontmatter/  frontmatter parsing (later)
dist/               build output (gitignored)
```

## Develop

```sh
make tidy     # fetch deps
make build    # build into dist/harness
make test     # run tests
make run      # go run ./cmd/harness
```

## Module path

`github.com/prodeng-fm/agent-tools/harness`

The repo is currently published as `prodeng-fm/resources` and will be
renamed to `agent-tools`. `go install` will work once the rename is
done.
