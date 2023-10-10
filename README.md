![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)
# Overview

The purpose of this project is to demonstrate the development of an application to enable a user perform simple statistical based in past draws using *Golang concurrency* and *SQL DB*. Note, the emphasis of this project to demonstrate software engineering principles.

It is worth noting whilst the application derived from this project has functionalities to perform statistical analysis, the intention is **not** to deliver something that can accurately predict lottery winnings. If the application generates or suggests a winning draw, it is purely coincidental. If you are expecting an application that accurately predicts lottery winnings, please refer to other projects.

## Use case

For this project, we will build two applications for two fictional users (also known as personas). It is worth noting that personas, for real life projects, are stand-in for real users. Personas are a vehicle to help member of a product development team empathize with users's needs and wants. The product team would use information obtained to focus on delivering features that will fulfill the needs and wants of the users.

However, for this project we have not considered beyond the superficial and fictional needs of our personas named `Ebenezer` and `Richie`. The personas are not tied to any real users anyway.

Let's summarise needs of the personas:

* `Ebenezer` is familiar with using applications via command line instruction (CLI). He intends to use the application to collect collect draws and store it in a database and to be able to interrograte the database via CLI.

* `Richie` does not have the necessary technical knowlege to work with applications completely via CLIs. He has only work with applications via Graphical User Interface (UI). He intends to use the application from a mac.  

The common functionalities expected from `Ebenezer` and `Richie` are features to: extract lottery draws in CSV from the UK National lottery website, store the downloaded draws in a persistent repository and perform statistical analysis of stored draws.

Please refer to the section named [user case analysis](./docs/usecase.md) to learn more about how the project derive use cases (via Lean Startup methodology, Design Thinking) and Agile management.

## Architecture

The project is organised based on the hexagonal [architecture principle](./docs/img/hexagonal.png).

The folders are layout as follows:

* `build` - docker script responsible for building binaries
* `cmd` - folder containing Go main packages that are further divided into several sub folders to reflect multiple exectuable builds
* `examples` - folder contains snippets of code to support testing or verification of edge cases that cannot easily be tested using unit tests
* `internal` - common Go packages that are accessible via main packages under cmd
* `scripts` - a collection of Bash script to trigger operations such as build, benchmarking, etc
* `testdata` - data intended to support test throughout the project

Please refer to the section named [architecture](./docs/arch.md) for in-depth discussions of engineering decisions behind the project.

## DevOps

Currently, the DevOps only support the creation of macOS based application (named `ebz`) for Ebenezer. To access `ebz` executable, you will need to build it yourself. The steps are:

1. Git clone this project.
1. At the root of the project, run the command: `./scripts/ebz.sh build`
1. You will find the application in `./build/package/macOS/ebz`

When the project is fully developed, all applications will be built via Docker, which you can inspect in these docker scripts under this folder `./build/`

## Disclaimer

This project incorporates mission-critical principles where practicable; however, it is not intended for use in any mission-critical settings. The author(s) of the project are not liable for any consequences resulting from its use outside the scope of this project setting.

Please note that this is an evolving project and is subject to changes without prior notification.

## Copyright notice

Unless otherwise specified, the copyright of this project is assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

