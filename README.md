# KafkaLib

KafkaLib is a library for publishing and consuming Kafka messages using Protobuf.

# Usage

For usage example, please check the `example` directory.

The basic concept is that you define your message in a proto file, and then use the generated code to publish and consume messages.

The convention of this library is one topic is used for one message type, and the topic name can be defined in the proto file using the `kafkalib.opts.topic_name` option. If the model generated is missing this option, and then being used to publish or consume messages, the library will return an error.

# Development setup

## Prerequisites

- Go 1.21
- `buf` cli (https://buf.build/docs/cli/installation/)

## Setup

1. `buf dep update`

## Generating code

1. After doing change to the proto files, run `buf generate`.
2. Sometime `gopls` would have errors after generate, just reload the IDE window.