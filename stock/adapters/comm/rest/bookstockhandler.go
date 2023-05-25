package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/bookstore/stock/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/bookstore/stock/adapters/comm/rest/middleware"
	"github.com/serdarkalayci/bookstore/stock/application"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type validatedbookStock struct{}

// swagger:route GET /bookStock/{id} bookStock GetBookStock
// Return the bookStock with the given id
// responses:
//	200: OK
//  400: Bad Request
//	500: errorResponse

// GetBookStock gets the bookStocks of the Titanic with the given id
func (apiContext *APIContext) GetBookStock(rw http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	duration, _ := otel.Meter("GetBook").Int64Histogram("work_duration")
	counter, _ := otel.Meter("GetBook").Int64Counter("request_counter")

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	ctx, span := createSpan(ctx, "Rest:BookStockHandler:GetBook", r)
	defer span.End()

	// parse the bookStock id from the url
	vars := mux.Vars(r)
	id := vars["id"]
	BookStockService := application.NewBookStockService(apiContext.bookStockRepo)
	bookStock, err := BookStockService.Get(ctx, id)
	opts := metric.WithAttributes(
		attribute.Key("Service").String("BookStock"),
		attribute.Key("Method").String("GetBook"),
	)
	duration.Record(ctx, time.Since(startTime).Milliseconds(), opts)
	counter.Add(ctx, 1, opts)
	if err != nil {
		switch err.(type) {
		case *application.ErrorCannotFindBookStock:
			respondWithError(rw, r, 404, "Cannot get book stock from database")
		default:
			respondWithError(rw, r, 500, "Internal server error")
		}
	} else {
		pDTO := mappers.MapBookStock2BookStockResponseDTO(bookStock)
		respondWithJSON(rw, r, 200, pDTO)
	}
}



// MiddlewareValidateNewBookStock Checks the integrity of new bookStock in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateNewBookStock(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		bookStock, err := middleware.ExtractBookStockRequestPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the bookStock
		errs := apiContext.validation.Validate(bookStock)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the bookStock")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		apiContext := context.WithValue(r.Context(), validatedbookStock{}, *bookStock)
		r = r.WithContext(apiContext)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
