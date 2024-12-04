# butts-service :peach:

## Requirements

- [Go](https://go.dev/)
- [NATS](https://nats.io/)

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
