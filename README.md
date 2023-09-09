
# pprof Server README
## Table of Contents
- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Profiling Endpoints](#profiling-endpoints)
- [Contributing](#contributing)

## Introduction

This repository provides a pprof server for profiling Go applications. pprof is a powerful tool that allows you to analyze the performance and behavior of your Go applications in real-time. This pprof server makes it easy to expose pprof endpoints in your Go applications and view the profiling data in your web browser.

## Getting Started

### Prerequisites

Before you can use this pprof server, make sure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (Go programming language)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) (Version control system)

### Installation

1. Clone this repository to your local machine:

   ```shell
   git clone https://github.com/adehikmatfr/go-pprof.git
   ```

2. Change into the project directory:

   ```shell
   cd go-pprof
   ```

3. Build the pprof server:

   ```shell
   go build
   ```

## Usage

To start the pprof server, run the following command:

```shell
./go-pprof serve-http
```

By default, the pprof server will run on port 3333. You can specify a different port using the `-port` flag:

```shell
./go-pprof serve-http -port 9090
```

Once the server is running, you can access the pprof web UI in your web browser by navigating to `http://localhost:3333` (or the port you specified). You will see a list of available profiling endpoints.

## Profiling Endpoints

The pprof server exposes various profiling endpoints for your Go application. Here are some common endpoints you can access:

- **/debug/pprof/heap**: Heap profile
- **/debug/pprof/goroutine**: Goroutine profile
- **/debug/pprof/cpu**: CPU profile
- **/debug/pprof/threadcreate**: Thread creation profile
- **/debug/pprof/block**: Block profile
- **/debug/pprof/mutex**: Mutex profile

For example, to capture a CPU profile, you can run:

```shell
go tool pprof http://localhost:3333/debug/pprof/cpu
```

Refer to the [pprof documentation](https://pkg.go.dev/net/http/pprof) for more information on how to use these endpoints to profile your Go application.

## Contributing

Contributions to this project are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.
