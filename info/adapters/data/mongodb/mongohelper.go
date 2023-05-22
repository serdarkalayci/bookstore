package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/bookstore/info/application"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoHelper struct {
	coll *mongo.Collection
}

func (mh mongoHelper) Find(ctx context.Context, filter bson.M, projection bson.M) (*mongo.Cursor, error) {
	dbctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	opts := options.Find()
	if len(projection) > 0 {
		opts.SetProjection(projection)
	}
	cur, err := mh.coll.Find(dbctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Msgf("error getting BookInfo list")
		return nil, err
	}
	return cur, nil
}

func (mh mongoHelper) InsertOne(ctx context.Context, document interface{}) (string, error) {
	result, err := mh.coll.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", result.InsertedID), nil
}

func (mh mongoHelper) FindOne(ctx context.Context, filter bson.M, projection bson.M) (*mongo.SingleResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	opts := options.FindOne()
	if len(projection) > 0 {
		opts.SetProjection(projection)
	}
	singleResult := mh.coll.FindOne(dbctx, filter, opts)
	if singleResult.Err() != nil {
		log.Error().Err(singleResult.Err()).Msgf("error getting BookInfo with the filter: %v", filter)
		return nil, &application.ErrorCannotFindBook{}
	}
	return singleResult, nil
}

func (mh mongoHelper) UpdateOne(ctx context.Context, id string, update interface{}) (int, error) {
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	result, err := mh.coll.UpdateOne(ctx, bson.M{"uuid": id}, update, &updateOpts)
	return int(result.ModifiedCount), err
}

func (mh mongoHelper) DeleteOne(ctx context.Context, id string) (int, error) {
	result, err := mh.coll.DeleteOne(ctx, bson.M{"uuid": id})
	return int(result.DeletedCount), err
}
