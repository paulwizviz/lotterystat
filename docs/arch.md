# Architecture

This section discusses thinking behind the layout of the project and other implementation decisions

## Project Layout

This artefacts in this projects are layout this way:

* `cmd` - a folder containing main packages
* `docs` - a collection of markdown documents
* `examples` - a collection of code snippets to verify certain ideals and thoughts of implementation strategy
* `internal` - shared packages to service the main packages in `cmd`
* `scripts` - bash scripts to support DevOps such as build and local deployment
* `test` - a container for stuff to support testing
* `web` - a container for one or more Web UI codes, such as ReactJS

## Database operations

There is a requirement for common interfaces to persists data into either SQL and NoSQL as seamlessly as possible whilst providing the flexibility to facility nuanced implementation specific to each DB types.

The design pattern chosen to based our DB related implementation is the handler or chain-of-reponsibility.