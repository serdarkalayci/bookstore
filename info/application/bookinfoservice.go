// Package application contains the application logic of the bookInfo management system.
package application

import (
	"context"

	"github.com/serdarkalayci/bookstore/info/domain"
	"go.opentelemetry.io/otel"
)

// BookInfoRepository is the interface that we expect to be fulfilled to be used as a backend for BookInfo Service
type BookInfoRepository interface {
	List(ctx context.Context) ([]domain.BookInfo, error)
	Get(ctx context.Context, ISBN string) (domain.BookInfo, error)
}

// BookInfoService represents the struct which contains a BookInfoRepository and exports methods to access the data
type BookInfoService struct {
	bookInfoRepo BookInfoRepository
}

// NewBookInfoService creates a new BookInfoService instance and sets its repository
func NewBookInfoService(dr BookInfoRepository) BookInfoService {
	if dr == nil {
		panic("missing bookInfoRepository")
	}
	return BookInfoService{
		bookInfoRepo: dr,
	}
}

// List loads all the data from the included repository from the given space and returns them
// Returns an error if the repository returns one
func (ps BookInfoService) List(ctx context.Context) ([]domain.BookInfo, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Application:BookInfoService:List")
	defer childSpan.End()
	bookInfos, err := ps.bookInfoRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return bookInfos, nil
}

// Get selects the bookInfo from the included repository with the given unique identifier, and returns it
// Returns an error if the repository returns one
func (ps BookInfoService) Get(ctx context.Context, isbn string) (domain.BookInfo, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Application:BookInfoService:Get")
	defer childSpan.End()
	bookInfo, err := ps.bookInfoRepo.Get(ctx, isbn)
	return bookInfo, err
}

