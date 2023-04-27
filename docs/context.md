# Context

This section describes the use case, the architecture philosophy, and development methodology guiding this project's development lifecycle.

## Use Cases

The following use cases are the target market segment for this project:

* a macOS, Linux and Windows based native application (Codename **ebenezer**) 
* an RESTFul and GraphQL based API enabling third parties to build apps against it (Codename **richie**) 

### Ebenezer

For this use case, this project is delivering in two forms:

* a command line interface (cli) application
* an application with a graphical user interface

<ins>Command line interface application</ins>

The cli version of the app is intended for used by personas that has the technical know how to interact with the application via Linux like (and windows equivalent) shell. The project has identified the following personas as representative of the application's target market segment:

* Alice - A Mac user
* Bob - A Linux user
* Charlotte - A Windows user

All personas have advance knowledge in working with command line applications

<ins>Graphical User Interface application</ins>

Details to be presented later.

### Richie ###

Details to be presented later

## Architectural Decisions

The decisions influencing the architecture and deliverables of this project are guided by these methodologies:

* [Design thinking](https://www.interaction-design.org/literature/topics/design-thinking)
* The [go standard project layout](https://github.com/golang-standards/project-layout) for source code architecture (see also the blog [Design the architecture, name the components, document the details](https://paulwizviz.github.io/go/2022/12/23/go-proverb-architecture.html))

## Quality Considerations

This project is predominantly Go based. The principle as outlined in the official Go coding convention [Effective Go](https://go.dev/doc/effective_go) is used to ensure consistencies in the source code architecture.

The Test Driven Development methodology is also used in this project.

## Project Lifecycle Management

The project's development is based on the [lean startup methodology](https://theleanstartup.com/principles). To compliment lean startup, the project uses the kanban agile methodology. 

[Pivotal tracker](https://www.pivotaltracker.com/n/projects/2639054) is used to track the lifecycle the project.