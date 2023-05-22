package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseCollection interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type dbHelper interface {
	Find(ctx context.Context, filter bson.M, projection bson.M) (*mongo.Cursor, error)
	InsertOne(ctx context.Context, document interface{}) (string, error)
	FindOne(ctx context.Context, filter bson.M, projection bson.M) (*mongo.SingleResult, error)
	UpdateOne(ctx context.Context, id string, update interface{}) (int, error)
	DeleteOne(ctx context.Context, id string) (int, error)
}
