package environment

import "os"

var (
	DeployEnv                = os.Getenv("DEPLOY_ENV")
	OtelExporterOtlpProtocol = os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	OtelExporterOtlpEndpoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	OtelExporterOtlpHeaders  = os.Getenv("OTEL_EXPORTER_OTLP_HEADERS")
)
