package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nats.go/micro"
)

const (
	butts   = "Butts"
	version = "0.0.0"
	subject = "svc.butts"
)

func service(ctx context.Context, log *slog.Logger, kv jetstream.KeyValue) micro.Config {
	return config(butts, version, subject, micro.ContextHandler(ctx,
		func(ctx context.Context, req micro.Request) {
			var u, r uint64

			e, err := kv.Get(ctx, butts)
			if err == nil {
				u = binary.LittleEndian.Uint64(e.Value()) + 1
				r = e.Revision()
			}

			kv.Update(ctx, butts, uint64ToBytes(u), r)

			msg := fmt.Sprintf("%d %s!", u, butts)

			log.Info(msg)
			req.Respond([]byte(msg))
		}),
	)
}

func uint64ToBytes(val uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, val)

	return b
}

func config(name, version, subject string, handler micro.Handler) micro.Config {
	return micro.Config{
		Name:    name,
		Version: version,
		Endpoint: &micro.EndpointConfig{
			Subject: subject,
			Handler: handler,
		},
	}
}

func run(ctx context.Context, log *slog.Logger, url string) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return err
	}

	kv, err := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: butts,
	})
	if err != nil {
		return err
	}

	echo, err := micro.AddService(nc, service(ctx, log, kv))
	if err != nil {
		return err
	}
	defer echo.Stop()

	<-ctx.Done()

	return nil
}

type input struct {
	url string
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	url := env("NATS_URL", nats.DefaultURL)

	var in input

	flag.StringVar(&in.url, "url", url, "NATS URL to use")
	flag.Parse()

	if err := run(ctx, log, in.url); err != nil {
		panic(err)
	}
}
