# butts-service :peach:

An example of using [NATS](https://nats.io/) [JetStream](https://docs.nats.io/nats-concepts/jetstream)
[Key/Value Store](https://docs.nats.io/nats-concepts/jetstream/key-value-store)
and [Micro](https://github.com/nats-io/nats.go/tree/main/micro#nats-micro)

## Requirements

- [Go](https://go.dev/)

## Usage

> [!TIP]
> Running NATS with JetStream enabled under docker can be done like so:
> `docker pull nats:latest`
> and then:
> `docker run --rm -it --name nats-server -d -p 4222:4222 nats -js`

If you have a NATS Server (with JetStream :rocket: enabled) on `localhost:4222` then:

```
go run github.com/peterhellberg/butts-service@latest
```

Now you should be able to run:

```
go run github.com/nats-io/natscli/nats@latest req svc.butts --count 100 ""
```
