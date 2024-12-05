# Go Architecture

This project aims to provide examples of some of the popular software architectures you can utilise when building web services. These examples build out services with exactly the same functionality, only the architecture and associated design patterns change.

## The Architectures

I'll be demonstrating how to build a simple web service using the following architectures -

- [Simple](https://github.com/rmhubbert/go-architecture/tree/main/simple)
- [Layered](https://github.com/rmhubbert/go-architecture/tree/main/layered)
- [Modular](https://github.com/rmhubbert/go-architecture/tree/main/modular)
- [Model View Controller](https://github.com/rmhubbert/go-architecture/tree/main/mvc)
- [Hexagonal / Ports & Adapters](https://github.com/rmhubbert/go-architecture/tree/main/hexagonal)

## The Project

The project I'll be using to demonstrate these architectures is a simple HTTP server that provides an API for managing users and user roles for these projects. The implementation is kept as simple as possible (this code is definitely not production ready, and isn't intended to be!), and third party libraries are kept to a minimum. The idea is that there is just enough code to enable the functionality, so as to not distract from the actual point, which is to demonstrate the different architectures.

## Tests

Unit tests have also been omitted, but you will find some simple end to end tests that can be run to ensure the server is working as intended. These tests require one of the projects to be currently running on port 8080.

Once you have the server running, you can run the e2e tests by running the following command from this directory.

```go
go test ./...
```
