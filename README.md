# butts-service :peach:

## Requirements

- [Go](https://go.dev/)
- [NATS](https://nats.io/)

## Usage

If you have a NATS Server (with JetStream :rocket: enabled) on `localhost:4222` then:

```
go run github.com/peterhellberg/butts-service@latest
```

Now you should be able to run:

```
go run github.com/nats-io/natscli/nats@latest req svc.butts --count 100 ""
```
