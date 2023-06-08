package middleware

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

// Option allows configuration of the httptrace Extract()
// and Inject() functions.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

type config struct {
	propagators propagation.TextMapPropagator
}

func newConfig(opts []Option) *config {
	c := &config{propagators: otel.GetTextMapPropagator()}
	for _, o := range opts {
		o.apply(c)
	}
	return c
}

// WithPropagators sets the propagators to use for Extraction and Injection.
func WithPropagators(props propagation.TextMapPropagator) Option {
	return optionFunc(func(c *config) {
		if props != nil {
			c.propagators = props
		}
	})
}

// Extract returns the Attributes, Context Entries, and SpanContext that were encoded by Inject.
func Extract(ctx context.Context, req *http.Request, opts ...Option) ([]attribute.KeyValue, baggage.Baggage, trace.SpanContext) {
	c := newConfig(opts)
	ctx = c.propagators.Extract(ctx, propagation.HeaderCarrier(req.Header))

	attrs := append(
		semconv.HTTPServerAttributesFromHTTPRequest("", "", req),
		semconv.NetAttributesFromHTTPRequest("tcp", req)...,
	)

	return attrs, baggage.FromContext(ctx), trace.SpanContextFromContext(ctx)
}

// Inject sets attributes, context entries, and span context from ctx into
// the request.
func Inject(ctx context.Context, req *http.Request, opts ...Option) {
	c := newConfig(opts)
	c.propagators.Inject(ctx, propagation.HeaderCarrier(req.Header))
}