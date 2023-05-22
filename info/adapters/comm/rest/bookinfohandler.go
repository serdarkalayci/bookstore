package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/bookstore/info/adapters/comm/rest/mappers"
	"github.com/serdarkalayci/bookstore/info/adapters/comm/rest/middleware"
	"github.com/serdarkalayci/bookstore/info/application"
)

type validatedbookInfo struct{}

// swagger:route GET /book book GetBooks
// Return all the bookInfos
// responses:
//	200: OK
//	500: errorResponse

// GetBooks gets the tree of all the bookInfos inside a space.
func (apiContext *APIContext) GetBooks(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	ctx, span := createSpan(ctx, "Rest:BookInfoHandler:GetBooks", r)
	defer span.End()
	BookInfoService := application.NewBookInfoService(apiContext.bookInfoRepo)
	folder, err := BookInfoService.List(ctx)
	if err != nil {
		respondWithError(rw, r, 500, "Cannot get books from database")
	} else {
		respondWithJSON(rw, r, 200, mappers.MapBookInfos2BookInfoDTOs(folder))
	}
}

// swagger:route GET /bookInfo/{id} bookInfo GetBookInfo
// Return the bookInfo with the given id
// responses:
//	200: OK
//  400: Bad Request
//	500: errorResponse

// GetBookInfo gets the bookInfos of the Titanic with the given id
func (apiContext *APIContext) GetBookInfo(rw http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	ctx, span := createSpan(ctx, "docman.GetOne", r)
	defer span.End()

	// parse the bookInfo id from the url
	vars := mux.Vars(r)
	id := vars["id"]
	BookInfoService := application.NewBookInfoService(apiContext.bookInfoRepo)
	bookInfo, err := BookInfoService.Get(ctx, id)
	if err != nil {
		switch err.(type) {
		case *application.ErrorCannotFindBook:
			respondWithError(rw, r, 404, "Cannot get bookInfo from database")
		default:
			respondWithError(rw, r, 500, "Internal server error")
		}
	} else {
		pDTO := mappers.MapBookInfo2BookInfoResponseDTO(bookInfo)
		respondWithJSON(rw, r, 200, pDTO)
	}
}



// MiddlewareValidateNewBookInfo Checks the integrity of new bookInfo in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateNewBookInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		bookInfo, err := middleware.ExtractBookInfoRequestPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the bookInfo
		errs := apiContext.validation.Validate(bookInfo)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the bookInfo")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		apiContext := context.WithValue(r.Context(), validatedbookInfo{}, *bookInfo)
		r = r.WithContext(apiContext)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
