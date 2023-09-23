![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)
# Overview

The purpose of this project is to demonstrate the development of a simple cli based application to enable a user perform simple statistical based in past draws using *Golang concurrency* and *SQL DB*. Note, the emphasis of this project to demonstrate software engineering principles.

It is worth noting whilst the application derived from this project has functionalities to perform statistical analysis, the intention is **not** to deliver something that can accurately predict lottery winnings. If the application generates or suggests a winning draw, it is purely coincidental. If you are expecting an application that accurately predicts lottery winnings, please refer to other projects.

## Use case

For this project, we will build two applications for two fictional users (also known as personas). 

One persona is named `Ebenezer` and the other is named `Richie`.

`Ebenezer` is familiar with using applications via command line instruction (CLI). He intends to use the application to collect collect draws and store it in a database and to be able to interrograte the database via CLI. He also intends to make it available in the form a a server with RESTFul interfaces.

`Richie` does not have the necessary technical knowlege to work with applications completely via CLIs. He has only work with applications via Graphical User Interface (UI). He intends to use the application from a mac.  

The common functionalities expected from `Ebenezer` and `Richie` are features to: extract lottery draws in CSV from the UK National lottery website, store the downloaded draws in a persistent repository and perform statistical analysis of stored draws.

## Architecture

The project is organised based on the hexagonal architecture principle.

![Architecture Principle](./docs/img/hexagonal.png)

The folders are layout as follows:

* `build` - docker script responsible for building binaries
* `cmd` - folder containing Go main packages
* `examples` - folder contains main packages to illustrative applications and also for benchmarking and other performance analysis
* `internal` - common Go packages that are accessible via main packages under cmd
* `scripts` - a collection of Bash script to trigger operations such as build, benchmarking, etc
* `testdata` - data intended to support test throughout the project

In this project the business models are organised around these packages:

* `euro` -- package of model for Euro Lottery draw
* `sforl` -- package of models for Set for Life draw 

The repositories and services layers are organised around these packages:

* `repo` -- package of database related handlers
* `csvproc` -- package of csv operations

These packages are dependent on the business models

## Disclaimer

This project incorporates mission-critical principles where practicable; however, it is not intended for use in any mission-critical settings. The author(s) of the project are not liable for any consequences resulting from its use outside the scope of this project setting.

Please note that this is an evolving project and is subject to changes without prior notification.

## Copyright notice

Unless otherwise specified, the copyright of this project is assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

