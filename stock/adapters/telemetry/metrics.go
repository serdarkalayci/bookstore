package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
)

func initMetricProvider(ctx context.Context, conn *grpc.ClientConn, res *resource.Resource) (func(context.Context) error, error) {
	// Set up a metric exporter
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	metricProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(5*time.Second),
			),
		),
		sdkmetric.WithResource(res),
	)

	otel.SetMeterProvider(metricProvider)
	return func(ctx context.Context) error {
		if err := metricProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown MeterProvider: %w", err)
		}
		return nil
	}, nil
}