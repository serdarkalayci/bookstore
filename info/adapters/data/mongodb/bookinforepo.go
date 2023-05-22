package mongodb

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/bookstore/info/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/bookstore/info/adapters/data/mongodb/mappers"
	"github.com/serdarkalayci/bookstore/info/application"
	"github.com/serdarkalayci/bookstore/info/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
)

// BookInfoRepository holds the arangodb client and database name for methods to use
type BookInfoRepository struct {
	helper  dbHelper
}

func newBookInfoRepository(client *mongo.Client, databaseName string) BookInfoRepository {
	return BookInfoRepository{
		helper: mongoHelper{coll: client.Database(databaseName).Collection(viper.GetString("BookInfoCollection"))},
	}
}

// List loads all the bookInfo records from tha database and returns it
// Returns an error if database fails to provide service
func (br BookInfoRepository) List(ctx context.Context) ([]domain.BookInfo, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Data:BookInfoRepository:List")
	defer childSpan.End()
	projection := bson.M{"_id": 0, "isbn": 1, "title": 1, "author": 1}
	cursor, err := br.helper.Find(ctx, bson.M{}, projection)
	if err != nil {
		log.Error().Err(err).Msgf("error getting books")
		return nil, &application.ErrorCannotFindBooks{}
	}
	defer cursor.Close(ctx)
	var books []dao.BookInfoDAO
	if err = cursor.All(ctx, &books); err != nil {
		log.Error().Err(err).Msgf("error getting all books")
		return nil, &application.ErrorCannotFindBooks{}
	}
	if books == nil {
		return nil, &application.ErrorCannotFindBooks{}
	}
	return mappers.MapBookInfoDAOs2BookInfos(books), nil
}

// Get selects a single bookInfo from the database with the given unique identifier
// Returns an error if database fails to provide service
func (br BookInfoRepository) Get(ctx context.Context, ISBN string) (domain.BookInfo, error) {
	ctx, childSpan := otel.Tracer("BookStore").Start(ctx, "Data:BookInfoRepository:Get")
	defer childSpan.End()
	filter := bson.M{"isbn": ISBN}
	projection := bson.M{"_id": 0, "isbn": 1, "title": 1, "author": 1, "price": 1, "publishDate": 1}
	result, err := br.helper.FindOne(ctx, filter, projection)
	if err != nil {
		log.Error().Err(err).Msgf("error getting books")
		return domain.BookInfo{}, &application.ErrorCannotFindBooks{}
	}
	var book dao.BookInfoDAO
	err = result.Decode(&book)
	if err != nil {
		log.Error().Err(err).Msgf("error decoding book")
		return domain.BookInfo{}, &application.ErrorCannotFindBook{}
	}
	return mappers.MapBookInfoDAO2BookInfo(book), nil
}

// AddItem adds a new bookInfo or a new folder to the underlying database.
// It returns the bookInfo inserted on success or error
func (dr BookInfoRepository) Add(bookInfo domain.BookInfo, parentID string, spaceID string) (string, error) {
	return "", errors.New("not implemented")
}

// Update updates fields of a single bookInfo from the database with the given unique identifier
// Returns an error if database fails to provide service
func (dr BookInfoRepository) Update(id string, p domain.BookInfo) error {
	return errors.New("not implemented")
}

// Delete selects a single bookInfo from the database with the given unique identifier
// Returns an error if database fails to provide service
func (dr BookInfoRepository) Delete(id string) error {
	return errors.New("not implemented")
}
