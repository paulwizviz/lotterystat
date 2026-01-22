# Gemini Context Configuration

## Project Context

Refer to these documents for project context.

- @README.md (Project overview and use cases)
- @docs/architecture.md (Project tech stack, structure and system design)
- @docs/specs.md (User Interface, API and cli specification)
- @docs/standard.md (Coding standards)

## Gated Execution Protocols

<PROTOCOL:PLAN>

- **Directory Scan:** Always perform a directory scan using `ls -R` before planning to understand the current file state.
- **Language Verification:** Scan `README.md` and `/docs/*.md` for Americanisms (e.g., "initialize", "color"). List all required corrections in the plan.
- **Structural Mapping:** Identify every Go struct that needs a matching JS interface for the frontend.
- **Checklist Generation:** Provide a step-by-step checklist of every file to be created or modified.
- **STOP:** Do not write any code. **Wait for explicit user approval before moving to IMPLEMENT.**

</PROTOCOL:PLAN>

<PROTOCOL:IMPLEMENT>

- **Approval Confirmation:** Before writing the first line of code, ask: "I have the approved plan. May I begin implementation?"
- **Frontend Placement:** Implement Javascript/ReactJS logic strictly within the `./web` directory.
- **Entry Point:** Ensure the `main` package in `cmd/ebz` remains the application entry point.
- **Backend Architecture:** Place computational logic and APIs as shareable packages under `internal/`.
- **Embedding:** Take the built `index.html` and `bundle.js` and place them in `internal/ebzweb/public` for embedding.
- **Language Revision:** Apply the British English corrections to the documentation as identified in the approved plan (e.g., "initialise", "colour").
- **Verification:** Run `go fmt ./...` and `npm run lint` automatically after writing files to ensure standards are met.

</PROTOCOL:IMPLEMENT>
