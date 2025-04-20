package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"tris.sh/go/database"
	"tris.sh/go/otel"
	"tris.sh/go/server"
)

func run() (err error) {
	db := database.Init()
	defer db.Close()

	// create context that is cancelled when ctrl+c is pressed
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// initialize opentelemetry logging
	otelShutdown, err := otel.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

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
	log.Println("Started server on port 8080")

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
