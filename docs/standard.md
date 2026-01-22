# Coding Standard

The backend subsystem is based on Go. The frontend system is web-based, implemented using JavaScript and the ReactJS framework.

Throughout this project, the following conventions apply:

## Go

- Follow [Effective Go](https://go.dev/doc/effective_go) for idiomatic programming.
- Strict Go error handling; use `slog.Error` for logging.
- Error handling: Wrap errors with context using `%w`.
- Include `context.Context` for cancellation.
- Include a `doc.go` file and and Go doc in all packages with header like `//Package <package name> contains ...`
- Include a filename in all packages that has the same name of the package like `some.go` for `package some` to include all interfaces, constants, errors, etc., but create separate files for implementations.

## JavaScript/ReactJS

Follow these for best practices.

- [ReactJS - Coding Guidelines](https://github.com/pillarstudio/standards/blob/master/reactjs-guidelines.md).
- [Material UI bast practices](https://cursorrules.org/article/material-ui-cursor-mdc-file)
