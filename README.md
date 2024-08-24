![build](https://github.com/paulwizviz/go-web/workflows/build/badge.svg)
# Overview

The objective of this project is to demonstrate software engineering principles by developing an example application in Go that applies *concurrent programming* and *database integration* techniques.

The example application, named `ebz`, extracts UK national lottery results in CSV format, persists data into SQLite and PostgreSQL, and performs simple statistical analyses, such as frequency analysis (see Figure 1).

![ebz functionality](./assets/img/ebz.png)
<figcaption>Figure 1 - ebz Functionalities</figcaption>

## Project Layout

The layout of this project is based on the principles articulated in [this article](https://paulwizviz.github.io/go/2022/12/23/go-proverb-architecture.html).

## Architecture Patterns

This project has incorporated the following architecture patterns:

* [Fan-in and Fan-out](./docs/fan-in-out.md)
* [Worker Pool](./docs/worker-pool.md)

## Disclaimer

This project uses UK National Lottery draws for illustrative purposes only. It is intended to demonstrate software engineering principles, **not** to predict lottery results.

Please note that this is an evolving project and is subject to changes without prior notice.

## Copyright notice

Unless otherwise specified, the copyright of this project is assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

