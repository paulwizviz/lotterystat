# Lotterystat: System Instructions

## Project Objective

An application to help users analyse the results of the UK National Lottery.

## Tech Stack (Locked)

* **Backend:** Go 1.22+ using the Standard Library for core logic.
* **CLI Framework:** `github.com/spf13/cobra`.
* **Database:** SQLite (via `github.com/mattn/go-sqlite3`).
* **Frontend:** Web UI powered by **HTMX** and standard Go `html/template`.
* **Build:** Docker-based multi-stage builds targeting `/build/bin`.

## Project Architecture

* `/assets/`: Source files for image generation.
* `/assets/png` & `/assets/jpg`: Generated asset storage.
* `/build/`: Dockerfiles and build-related automation.
* `/build/bin/`: Destination for compiled binaries from `/cmd`.
* `/cmd/ebz/`: Primary application entry point.
* `/internal/`: Private library code.
* `/internal/ebzmux/`: HTMX handler logic and routing.

## Coding Standards

### Go

* Follow [Effective Go](https://go.dev/doc/effective_go) (Idiomatic patterns only).
* Use **Goroutines and Channels** for lottery data processing and analysis.
* Error handling: Wrap errors with context using `%w`.

## CLI Prompting Rules

* **Code Generation**: Output code only. No conversational filler or pleasantries. Limit narrative to "Next Steps" or specific file paths.
* **Documentation**: All Markdown (`README.md`, `/docs/*.md`) must use **British English** (e.g., *analyse, colour, optimise*).
* **README Structure**: Must include: Objective, Scope, Documentation References, and Copyright Notice.
