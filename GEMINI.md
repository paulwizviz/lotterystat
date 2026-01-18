# Lotterystat: System Instructions

## Project Objective

An application to help users analyse the results of the UK National Lottery.

## Tech Stack (Locked)

* **Backend:** Go 1.22+ using the Standard Library for core logic.
* **CLI Framework:** `github.com/spf13/cobra`.
* **Database:** SQLite (via `modernc.org/sqlite`).
* **Frontend:** Web UI powered by **HTMX** with Tailwind CSS, standard Go `html/template` and `ChartJS`.
* **Build:** Docker-based multi-stage builds targeting `/build/bin`.

## Project Architecture

* `/assets/`: Source files for image generation.
* `/assets/png` & `/assets/jpg`: Generated asset storage.
* `/build/`: Dockerfiles and build-related automation.
* `/build/bin/`: Destination for compiled binaries from `/cmd`.
* `/cmd/ebz/`: Primary application entry point.
* `/internal/config`: Private package to support configuration operation.
* `/internal/csvops`: Private package of operations to read and process CSV files.
* `/internal/sqlops`: Private package containing SQL operations.
* `/internal/tball`: Private package to support analysis of past Thunderball results.
* `/internal/ebzmux/css/input.css`: HTMX handler tailwind input specification.
* `/internal/ebzmux/static/htmx.min.js`: HTMX library.
* `/internal/ebzmux/templates/*.html`: HTML files.

## Coding Standards

### Go

* Follow [Effective Go](https://go.dev/doc/effective_go) (Idiomatic patterns only).
* Use **Goroutines and Channels** for lottery data processing and analysis.
* Error handling: Wrap errors with context using `%w`.

### Web

* Follow the [htmx standard](https://htmx.org/).
* Use tailwind cli like this `npx @tailwindcss/cli -i ./src/input.css -o ./src/output.css` to generate main.css.

## CLI Prompting Rules

* **Code Generation**: Output code only. No conversational filler or pleasantries. Limit narrative to "Next Steps" or specific file paths.
* **Documentation**: All Markdown (`README.md`, `/docs/*.md`) must use **British English** (e.g., *analyse, colour, optimise*).
* **README Structure**: Must include: Objective, Scope, Documentation References, and Copyright Notice.
