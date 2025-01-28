# Protobox
![GitHub release](https://img.shields.io/github/release/notjustmoney/protobox.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/notjustmoney/protobox)
![GoDoc](https://godoc.org/github.com/notjustmoney/protobox?status.svg)
[![Codecov](https://codecov.io/gh/notjustmoney/protobox/branch/main/graph/badge.svg)](https://codecov.io/gh/notjustmoney/protobox)


> For reliable delivery of In-process or Interprocess messages.

Protobox is a reliable message delivery framework for Go, providing both in-process and network message delivery capabilities with type safety and transaction support.



# Overview
<p>
<img src="./logo/protobox-gopher.png" style="width: 25%; float: right; margin: 0 0 10px 20px;" alt="protobox-gopher">
Protobox leverages Protocol Buffers (protobuf) to define message schemas and dispatchers through custom message options. Using Protobox's protoc-gen-go-protobox plugin alongside protobuf's protoc-gen-go plugin, you can generate type-safe message dispatchers and message structures with minimal boilerplate.
</p>

Key capabilities:

- Type-safe message handler registration
- In-memory message bus functionality
- Transactional message delivery using outbox pattern
- Idempotent message processing with inbox pattern
- Built-in support for both in-process and network message delivery


# Background
Modern distributed systems often require reliable message delivery mechanisms across various components, both within the same process and across network boundaries. While patterns like outbox and retry mechanisms are well-established solutions, implementing them repeatedly brings several challenges:

Defining and maintaining message schemas
Ensuring type safety across message handlers
Managing message persistence and delivery guarantees
Handling both in-process and network message delivery consistently
Maintaining idempotency in message processing

While robust messaging frameworks exist in other language ecosystems, the Go ecosystem lacked a comprehensive solution that leverages Go's strengths. Protobox addresses this gap by:

Utilizing Protocol Buffers for schema definition and code generation
Providing type-safe message handling through code generation
Offering built-in support for reliable messaging patterns
Embracing Go's idioms while providing high-performance implementations

# Features
## Core Features

- Schema-first message definition using Protocol Buffers
- Type-safe message dispatcher generation
- In-process message bus with handler registration
- Transactional outbox pattern support
- Idempotent message processing through inbox pattern

## Coming Soon

- Advanced outbox/inbox toolkit
  - Scalable message scheduling
  - Flexible message routing
  - Customizable storage interfaces
- Resilience patterns
  - Retryer
  - Circuit breaker
  - Bulkhead pattern implementation

# Prerequisites
- [Go](https://golang.org/dl/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/downloads)
- [Buf(Optional)](https://docs.buf.build/installation)

# Installation
```bash
go get -u github.com/notjustmoney/protobox
```

To install the protoc plugin:
```bash
go install github.com/notjustmoney/protoc-gen-go-protobox
```

# Getting Started
Recommend to use [Buf](https://docs.buf.build/installation) for generating and managing Protobuf files.

For more examples, see [examples](./examples).

# Documentation
Documentation is available soon!

# Contributing
Issues, PRs for design suggestion, contributing are welcome.

# License
This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=notjustmoney/protobox&type=Date)](https://star-history.com/#notjustmoney/protobox&Date)