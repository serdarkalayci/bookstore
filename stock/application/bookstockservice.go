// Package application contains the application logic of the bookStock management system.
package application

import (
	"context"

	"github.com/serdarkalayci/bookstore/stock/domain"
	"go.opentelemetry.io/otel"
)

// BookStockRepository is the interface that we expect to be fulfilled to be used as a backend for BookStock Service
type BookStockRepository interface {
		Get(ctx context.Context, ISBN string) (domain.BookStock, error)
}

// BookStockService represents the struct which contains a BookStockRepository and exports methods to access the data
type BookStockService struct {
	bookStockRepo BookStockRepository
}

// NewBookStockService creates a new BookStockService instance and sets its repository
func NewBookStockService(br BookStockRepository) BookStockService {
	if br == nil {
		panic("missing bookStockRepository")
	}
	return BookStockService{
		bookStockRepo: br,
	}
}

// Get selects the bookStock from the included repository with the given unique identifier, and returns it
// Returns an error if the repository returns one
func (bs BookStockService) Get(ctx context.Context, isbn string) (domain.BookStock, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Application:BookStockService:Get")
	defer childSpan.End()
	bookStock, err := bs.bookStockRepo.Get(ctx, isbn)
	return bookStock, err
}

