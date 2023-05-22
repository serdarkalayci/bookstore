package rest

import (
	"context"
	"net/http"
	"time"

	openapimw "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	middleware "github.com/serdarkalayci/bookstore/info/adapters/comm/rest/middleware"
	"github.com/serdarkalayci/bookstore/info/application"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// APIContext handler for getting and updating Ratings
type APIContext struct {
	validation *middleware.Validation
	//dbContext  DBContext
	healthRepo    application.HealthRepository
	bookInfoRepo  application.BookInfoRepository
	configuration map[string]string
	tp 		  trace.TracerProvider
}

// NewAPIContext returns a new APIContext handler with the given logger
func NewAPIContext(bindAddress *string, hr application.HealthRepository, br application.BookInfoRepository) (*http.Server) {
	apiContext := &APIContext{
		healthRepo:   hr,
		bookInfoRepo: br,
	}
	return apiContext.prepareContext(bindAddress)

}

func (apiContext *APIContext) prepareContext(bindAddress *string) (*http.Server) {
	otel.SetTextMapPropagator(propagation.TraceContext{})
	apiContext.validation = middleware.NewValidation()

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	sm.Use(middleware.MetricsMiddleware)
	// CORS handler
	optR := sm.Methods(http.MethodOptions).Subrouter()
	optR.HandleFunc("/{path:.*}", CorsHandler)
	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	// Generic handlers
	getR.HandleFunc("/", apiContext.Index)
	getR.HandleFunc("/version", apiContext.Version)
	getR.HandleFunc("/health/live", apiContext.Live)
	getR.HandleFunc("/health/ready", apiContext.Ready)
	// bookInfo handlers
	getR.HandleFunc("/book/{id}", apiContext.GetBookInfo)
	getR.HandleFunc("/book", apiContext.GetBooks)
	// Documentation handler
	opts := openapimw.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openapimw.Redoc(opts, nil)
	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create a new server
	s := &http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	sm.PathPrefix("/metrics").Handler(promhttp.Handler())
	prometheus.MustRegister(middleware.RequestCounterVec)
	prometheus.MustRegister(middleware.RequestDurationGauge)

	return s
}

// createSpan extracts the span from the request if exists or creates a new one using openTelemetry. Span with the given name and returns it
func createSpan(ctx context.Context, opName string, r *http.Request) (context.Context, trace.Span) {
	spanContext := otel.GetTextMapPropagator().Extract(
		ctx,
		propagation.HeaderCarrier(r.Header))

	ctx, span := otel.Tracer("BookStore").Start(
		spanContext,
		opName,
	)
	return ctx, span
}

// injectSpanToResponse injects the span context into the response header
func injectSpanContextToResponse(ctx context.Context, w http.ResponseWriter) {
	// Inject the span context into the response header
	otel.GetTextMapPropagator().Inject(
		ctx,
		propagation.HeaderCarrier(w.Header()))
}


