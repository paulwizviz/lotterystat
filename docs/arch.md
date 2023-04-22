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

## `Ebenezer`

There are two versions of `Ebenezer`:

* `ebzcli` - A command line version of the app
* `ebzui` - An embedded web-based application

By default, the application stores its runtime state in the folder `$HOME/.ebz`. A component of the runtime state folder is a SQLite db file. The location of the state folder is customisable. The application configuration file in macOS and Linux are located in the folder where the app is store.