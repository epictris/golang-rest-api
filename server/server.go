package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/protobuf/proto"

	"github.com/bufbuild/protovalidate-go"
	"github.com/epictris/go/api/errors"
	"github.com/epictris/go/gen/proto/v1/apiv1connect"
	"github.com/epictris/go/server/handlers"
)

func errorInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				span := trace.SpanFromContext(ctx)
				response, err := next(ctx, req)
				if err != nil {
					span.RecordError(err, trace.WithStackTrace(true))
					log.Printf("%v", err)

					if error, ok := err.(errors.APIError); ok {
						return nil, error
					} else {
						return nil, errors.APIErrorInternalServerError()
					}
				}
				return response, nil
			},
		)
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func validationInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				validator, err := protovalidate.New()
				if err != nil {
					panic(err)
				}
				msg := req.Any()

				protoMsg, ok := msg.(proto.Message)
				if !ok {
					return nil, connect.NewError(
						connect.CodeInternal,
						fmt.Errorf("message doesn't implement proto.Message interface"),
					)
				}

				if err := validator.Validate(protoMsg); err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, err)
				}
				return next(ctx, req)
			},
		)
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func NewHTTPHandler(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	server := &handlers.Server{
		DB: db,
	}

	opts := []connect.HandlerOption{
		connect.WithInterceptors(validationInterceptor(), errorInterceptor()),
	}

	path, handler := apiv1connect.NewServiceHandler(server, opts...)
	mux.Handle(path, handler)

	// add otel HTTP logging instrumentation for all routes
	return h2c.NewHandler(otelhttp.NewHandler(mux, "/"), &http2.Server{})
}
