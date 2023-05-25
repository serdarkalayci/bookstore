// Package telemetry provides a common interface for telemetry adapters.
package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Initializes an OTLP exporter, and configures the corresponding trace and
// metric providers.
func InitOTEL(ctx context.Context, endpoint string) (func(context.Context) error, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display in observer backends
			semconv.ServiceNameKey.String("Bookstore"),
		),
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// OpenTelemtry Collector is running on docker-compose in same network as this go app,
	// and exposes the OTLP receiver on port 4317.
	conn, err := grpc.DialContext(ctx, endpoint,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	traceShutdownFunc, err := initTracerProvider(ctx, conn, res)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize trace provider: %w", err)
	}
	metricShutdownFunc, err := initMetricProvider(ctx, conn, res)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize metric provider: %w", err)
	}
	return func(ctx context.Context) error {
		if err := traceShutdownFunc(ctx); err != nil {
			return fmt.Errorf("failed to shutdown trace provider: %w", err)
		}
		if err := metricShutdownFunc(ctx); err != nil {
			return fmt.Errorf("failed to shutdown metric provider: %w", err)
		}
		return nil
	}, nil
}