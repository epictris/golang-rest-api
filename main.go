package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/epictris/go/database"
	"github.com/epictris/go/server"
	"github.com/epictris/go/telemetry"
)

func run() (err error) {
	db := database.Init()
	defer db.Close()

	// create context that is cancelled when ctrl+c is pressed
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := telemetry.OtelInit(ctx)
	if err != nil {
		return err
	}
	defer otelShutdown(ctx)

	// start server
	srv := &http.Server{
		Addr:        ":8080",
		BaseContext: func(_ net.Listener) context.Context { return ctx },
		Handler:     server.NewHTTPHandler(db),
	}
	serverError := make(chan error, 1)
	go func() {
		serverError <- srv.ListenAndServe()
	}()

	// wait for server to encounter an error or get cancelled
	select {
	case err = <-serverError:
		return
	case <-ctx.Done():
		stop()
		err = srv.Shutdown(context.Background())
		return
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
