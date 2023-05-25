package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/bookstore/stock/adapters/data/redis/mappers"
	"github.com/serdarkalayci/bookstore/stock/application"
	"github.com/serdarkalayci/bookstore/stock/domain"
	"go.opentelemetry.io/otel"
)

// BookStockRepository holds the client for the database
type BookStockRepository struct {
	helper  dbHelper
}

func newBookStockRepository(client *redis.Client) BookStockRepository {
	return BookStockRepository{
		helper: redisHelper{client: client},
	}
}

// Get selects a single bookStock from the database with the given unique identifier
// Returns an error if database fails to provide service
func (br BookStockRepository) Get(ctx context.Context, ISBN string) (domain.BookStock, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Data:BookStockRepository:Get")
	defer childSpan.End()
	stock, err := br.helper.GetIntValue(ctx, ISBN)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting BookStock")
		return domain.BookStock{}, &application.ErrorCannotFindBookStock{ISBN: ISBN}
	}

	return mappers.MapBookStockDAO2BookStock(ISBN, stock), nil
}

// Add adds a new bookStock or a new folder to the underlying database.
// It returns the bookStock inserted on success or error
func (dr BookStockRepository) Add(bookStock domain.BookStock, parentID string, spaceID string) (string, error) {
	return "", errors.New("not implemented")
}

// Update updates fields of a single bookStock from the database with the given unique identifier
// Returns an error if database fails to provide service
func (dr BookStockRepository) Update(id string, p domain.BookStock) error {
	return errors.New("not implemented")
}

// Delete selects a single bookStock from the database with the given unique identifier
// Returns an error if database fails to provide service
func (dr BookStockRepository) Delete(id string) error {
	return errors.New("not implemented")
}
