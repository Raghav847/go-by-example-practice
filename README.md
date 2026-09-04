# Go By Example Practice

This repository is a personal sandbox for testing, experimenting, and learning Go. It contains a mix of standalone runnable examples, experiments, and small projects organized by directory.

## Structure

- `lang-specific/` — Go practice examples (e.g. `01_basics`), each a standalone `main` package
- `experiments/` — ad-hoc experiments
- `goBlog/` — a Go blog project
- `gopl/` — exercises from *The Go Programming Language* book
- `htmxGo/` — Go + htmx experiments
- `kvStore/` — a key-value store project

## Running examples

Run any standalone example with:

```bash
go run ./lang-specific/01_basics/hello_world/
```

Each example directory is its own `main` package. Some directories are currently scaffolded placeholders so the project matches the requested layout and can be filled in incrementally.