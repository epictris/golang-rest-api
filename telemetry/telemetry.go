package telemetry

import (
	"context"
	"errors"
	"log"

	"github.com/epictris/go/environment"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// bootstraps the OpenTelemetry pipeline.
func OtelInit(ctx context.Context) (shutdown func(context.Context) error, err error) {

	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	if environment.OtelExporterOtlpEndpoint == "" || environment.OtelExporterOtlpHeaders == "" ||
		environment.OtelExporterOtlpProtocol == "" {
		log.Println("No OpenTelemetry environment variables found. Skipping OTEL setup.")
		return
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName("go-rest-api"),
			semconv.ServiceVersion("0.0.0"),
			semconv.DeploymentEnvironment(environment.DeployEnv),
		),
	)
	if err != nil {
		handleErr(err)
		return
	}

	shutdownLogger, err := initLoggerProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(
		shutdownFuncs,
		shutdownLogger,
	)

	shutdownTracer, err := initTracerProvider(ctx, res)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(
		shutdownFuncs,
		shutdownTracer,
	)
	return
}

func initLoggerProvider(
	ctx context.Context,
	res *resource.Resource,
) (func(context.Context) error, error) {
	exporter, err := otlploghttp.New(ctx)
	if err != nil {
		return nil, err
	}
	provider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(res),
	)
	global.SetLoggerProvider(provider)
	return provider.Shutdown, nil
}

func initTracerProvider(
	ctx context.Context,
	res *resource.Resource,
) (func(context.Context) error, error) {
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, err
	}
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(provider)
	return provider.Shutdown, nil
}
