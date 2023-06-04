# Context

This section describes the use case, the architecture philosophy, and development methodology guiding this project's development lifecycle.

## Use Cases

The outcome of this project in intended for two personas, Alice and Bob.

<u>Alice</u>

Alice is a frontend developer who intends to build a front end applications against a RESTful API. For her, this project will deliver an API based solution codename `ebenezer` (shortname `ebz`).

<u>Bob</u>

Bob intends to use an application natively on a (macOS, Linux or Windows) or a cloud based web app. He is not a developer, only a user of the app. For him, this project will deliver a solution codename `richie`.

## Architectural Decisions

The decisions influencing the architecture and deliverables of this project are guided by these methodologies:

* [Design thinking](https://www.interaction-design.org/literature/topics/design-thinking)
* The source code architecture is based on a structure described in the blog [Design the architecture, name the components, document the details](https://paulwizviz.github.io/go/2022/12/23/go-proverb-architecture.html)
* A backend based on the microservices architecture 

## Quality Considerations

This project is predominantly Go based. The principle as outlined in the official Go coding convention [Effective Go](https://go.dev/doc/effective_go) is used to ensure consistencies in the source code architecture.

The Test Driven Development methodology is also used in this project.

## Project Lifecycle Management

The project's development is based on the [lean startup methodology](https://theleanstartup.com/principles). To compliment lean startup, the project uses the kanban agile methodology. [Pivotal tracker](https://www.pivotaltracker.com/n/projects/2639054) is used to track the lifecycle the project.