# Events and messages

There are several message brokers for go and we'll sample a few here in our
testcases.

## Requirements

- go 1.23
- testcontainers-go 0.34
- sarama 1.43

## How to build

## How to run/test

Since this is built as a library, there is no main module. so all you have is
the test command:

```bash
go test -v ./...
```

## Noteworthy

- testify library is quite good, especially the bootstrap/setup step
- testcontainers for kafka handles well our needs
